package app

import (
	"context"
	"github.com/alserov/car_insurance/contract/internal/async"
	"github.com/alserov/car_insurance/contract/internal/config"
	"github.com/alserov/car_insurance/contract/internal/logger"
	server "github.com/alserov/car_insurance/contract/internal/server/async"
	"github.com/alserov/car_insurance/contract/internal/service/models"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	log := logger.NewLogger(logger.Zap, cfg.Env)

	log.Info("starting server")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	payoffCons := async.NewConsumer[models.Payoff](async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.Payoff)
	newInsuranceCons := async.NewConsumer[models.NewInsurance](async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.NewInsurance)

	log.Info("server is running")

	server.StartServer(ctx, server.Consumers{PayoffCons: payoffCons, NewInsuranceCons: newInsuranceCons})

	log.Info("server stopped")
}
