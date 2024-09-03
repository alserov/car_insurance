package server

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/server/mux"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"go.opentelemetry.io/otel/trace"
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
	Mux
)

func NewServer(routerType uint, srvc *service.Service, tracer trace.Tracer, log logger.Logger) Server {
	switch routerType {
	case Mux:
		return mux.NewServer(srvc, tracer, log)
	default:
		panic("invalid router type")
	}
}
