package server

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/server/fiber"
	"github.com/alserov/car_insurance/gateway/internal/service"
)

type Server interface {
	Serve(port string) error
	Shutdown() error
}

func MustServe(ctx context.Context, srv Server, port string) {
	go func() {
		if err := srv.Serve(port); err != nil {
			panic("failed to start server: " + err.Error())
		}
	}()

	<-ctx.Done()

	_ = srv.Shutdown()
}

const (
	Fiber = iota
)

func NewServer(routerType uint, srvc *service.Service, log logger.Logger) Server {
	switch routerType {
	case Fiber:
		return fiber.NewServer(srvc, log)
	default:
		panic("invalid router type")
	}
}
