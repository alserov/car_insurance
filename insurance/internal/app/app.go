package app

import (
	"github.com/alserov/car_insurance/insurance/internal/clients/async"
	"github.com/alserov/car_insurance/insurance/internal/clients/http"
	"github.com/alserov/car_insurance/insurance/internal/config"
	"github.com/alserov/car_insurance/insurance/internal/server"
	"github.com/alserov/car_insurance/insurance/internal/server/grpc"
	"github.com/alserov/car_insurance/insurance/internal/service"
)

func MustStart(cfg *config.Config) {

	// clients
	recongitionClient := http.NewRecognitionClient(cfg.Services.RecognitionAddr)
	contractClient := async.NewContractClient(cfg.Broker.Addr, cfg.Broker.Topics.NewInsurance, cfg.Broker.Topics.Payoff)

	// service
	srvc := service.NewService(service.Clients{
		Recognition: recongitionClient,
		Contract:    contractClient,
	})

	// server
	srvr := grpc.NewServer(srvc)

	server.MustServe(srvr, cfg.Port)
}
