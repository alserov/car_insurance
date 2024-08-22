package app

import (
	"github.com/alserov/car_insurance/gateway/internal/clients/grpc"
	"github.com/alserov/car_insurance/gateway/internal/config"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/server"
	"github.com/alserov/car_insurance/gateway/internal/service"
)

func MustStart(cfg *config.Config) {
	// logger
	log := logger.NewLogger(logger.Zap, cfg.Env)

	log.Info("starting server")

	// *grpc* clients
	insuranceClient := grpc.NewInsuranceClient()

	cls := service.Clients{
		InsuranceClient: insuranceClient,
	}

	// service
	srvc := service.NewService(cls)

	// server
	srvr := server.NewServer(server.Fiber, srvc)

	// starting server
	srvr.Serve(cfg.Port)
}
