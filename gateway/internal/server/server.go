package server

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/cache"
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

type Args struct {
	RouterType int
	Service    *service.Service
	Cache      cache.Cache
	Tracer     trace.Tracer
	Log        logger.Logger
}

func NewServer(arg Args) Server {
	switch arg.RouterType {
	case Mux:
		return mux.NewServer(arg.Service, arg.Cache, arg.Tracer, arg.Log)
	default:
		panic("invalid router type")
	}
}
