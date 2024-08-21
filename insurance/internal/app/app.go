package app

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/async"
	async_cl "github.com/alserov/car_insurance/insurance/internal/clients/async"
	http_cl "github.com/alserov/car_insurance/insurance/internal/clients/http"
	"github.com/alserov/car_insurance/insurance/internal/config"
	"github.com/alserov/car_insurance/insurance/internal/db/mongo"
	"github.com/alserov/car_insurance/insurance/internal/db/postgres"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"github.com/alserov/car_insurance/insurance/internal/server"
	"github.com/alserov/car_insurance/insurance/internal/server/grpc"
	"github.com/alserov/car_insurance/insurance/internal/service"
	"github.com/alserov/car_insurance/insurance/internal/utils"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// logger
	log := logger.NewLogger(logger.Zap, cfg.Env)

	log.Info("starting server", logger.WithArg("port", cfg.Port))

	// db (connections to all DBs)
	mongoConn := mongo.MustConnect(cfg.Databases.Mongo.Addr)
	defer mongoConn.Disconnect(context.Background())

	postgresConn := postgres.MustConnect(cfg.Databases.Postgres.Addr)
	defer postgresConn.Close()

	// repositories (initializing repos from db connections)
	repo := postgres.NewRepository(postgresConn)
	outboxRepo := mongo.NewOutbox(mongoConn)

	// *http* clients (initializing clients which use http)
	recognitionClient := http_cl.NewRecognitionClient(cfg.Services.RecognitionAddr)

	// *async* clients (initializing clients which use message queues/brokers)
	contractProducer := async.NewProducer(async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.NewInsurance)
	defer contractProducer.Close()

	contractClient := async_cl.NewContractClient(contractProducer)

	// outbox worker (initializing outbox worker, that will interact with other services)
	outboxWorker := utils.NewOutboxWorker(outboxRepo, contractClient, log)

	// service (initializing service)
	srvc := service.NewService(service.Clients{
		Recognition: recognitionClient,
		Contract:    contractClient,
	}, outboxRepo, repo)

	// server (initializing server)
	srvr := grpc.NewServer(srvc)

	// starting server and workers
	log.Info("server is running")

	outboxWorker.Start(ctx)
	server.MustServe(ctx, srvr, cfg.Port)

	log.Info("server stopped")
}
