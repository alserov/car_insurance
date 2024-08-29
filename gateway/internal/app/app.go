package app

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/clients/grpc"
	"github.com/alserov/car_insurance/gateway/internal/config"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/server"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"os/signal"
	"syscall"

	_ "github.com/alserov/car_insurance/gateway/docs"
	_ "github.com/joho/godotenv/autoload"
)

// @title Car insurance API
// @version 1.0
// @BasePath /v1

func MustStart(cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// logger
	log := logger.NewLogger(logger.Zap, cfg.Env)

	log.Info("starting server")

	// *grpc* clients
	insuranceConn := grpc.Dial(cfg.Services.Insurance.Addr)
	defer func() {
		_ = insuranceConn.Close()
	}()

	insuranceClient := grpc.NewInsuranceClient(insuranceConn)

	cls := service.Clients{
		InsuranceClient: insuranceClient,
	}

	// service
	srvc := service.NewService(cls)

	// server
	srvr := server.NewServer(server.Fiber, srvc, log)

	// starting server
	log.Info("server is running")

	server.MustServe(ctx, srvr, cfg.Port)

	log.Info("server stopped")
}
