package app

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/async"
	async_cl "github.com/alserov/car_insurance/insurance/internal/clients/async"
	http_cl "github.com/alserov/car_insurance/insurance/internal/clients/http"
	"github.com/alserov/car_insurance/insurance/internal/config"
	"github.com/alserov/car_insurance/insurance/internal/db/redis"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"github.com/alserov/car_insurance/insurance/internal/server"
	"github.com/alserov/car_insurance/insurance/internal/server/grpc"
	"github.com/alserov/car_insurance/insurance/internal/service"
	"github.com/alserov/car_insurance/insurance/internal/utils"
	"os/signal"
	"syscall"
	"time"
)

const (
	outboxTickPeriod = time.Second * 60 * 30
)

func MustStart(cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// utils
	log := logger.NewLogger(logger.Zap, cfg.Env)
	pool := utils.NewPool(ctx, 2)

	log.Info("starting server", logger.WithArg("port", cfg.Port))

	// *http* clients
	recognitionClient := http_cl.NewRecognitionClient(cfg.Services.RecognitionAddr)

	// *async* clients
	contractProducer := async.NewProducer(async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.NewInsurance)
	defer contractProducer.Close()

	contractClient := async_cl.NewContractClient(contractProducer)

	// db
	outbox := redis.NewOutbox(redis.MustConnect(cfg.Databases.Redis.Addr), outboxTickPeriod)
	_ = pool.Add(func() error {
		outbox.ProcessNotCommittedInsurances(ctx)
		return nil
	})
	_ = pool.Add(func() error {
		outbox.ProcessNotCommittedPayoffs(ctx)
		return nil
	})

	// service
	srvc := service.NewService(service.Clients{
		Recognition: recognitionClient,
		Contract:    contractClient,
	})

	// server
	srvr := grpc.NewServer(srvc)

	// starting server
	log.Info("server is running")

	server.MustServe(ctx, srvr, cfg.Port)

	log.Info("server stopped")
}
