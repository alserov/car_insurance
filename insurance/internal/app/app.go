package app

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/config"
	server "github.com/alserov/car_insurance/insurance/internal/server/grpc"
	"github.com/alserov/car_insurance/insurance/internal/service"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	srvc := service.NewService(service.Clients{
		Recognition: nil,
		Contract:    nil,
	})
	srvr := server.NewServer(srvc)

	run(srvr, cfg.Port)
}

type Server interface {
	Serve(host string) error
}

func run(server Server, port string) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	cancel()

	go func() {
		if err := server.Serve(port); err != nil {
			panic("failed to start server: " + err.Error())
		}
	}()

	<-ctx.Done()
}
