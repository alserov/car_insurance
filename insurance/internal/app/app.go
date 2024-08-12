package app

import (
	"github.com/alserov/car_insurance/insurance/internal/config"
	"github.com/alserov/car_insurance/insurance/internal/server"
	"github.com/alserov/car_insurance/insurance/internal/server/grpc"
	"github.com/alserov/car_insurance/insurance/internal/service"
)

func MustStart(cfg *config.Config) {
	srvc := service.NewService(service.Clients{
		Recognition: nil,
		Contract:    nil,
	})
	srvr := grpc.NewServer(srvc)

	server.MustServe(srvr, cfg.Port)
}
