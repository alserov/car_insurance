package app

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/cache"
	"github.com/alserov/car_insurance/gateway/internal/clients/grpc"
	"github.com/alserov/car_insurance/gateway/internal/config"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/server"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"github.com/alserov/car_insurance/gateway/internal/tracing"
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

	// tracer
	exp := tracing.NewOtlExporter(ctx, cfg.Tracing.Endpoint)
	defer func() {
		_ = exp.Shutdown(ctx)
	}()

	tracer, tp := tracing.NewTracer(exp, cfg.Tracing.Name)
	defer func() {
		_ = tp.Shutdown(ctx)
	}()

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

	// cache
	c := cache.NewCache(cache.Redis)

	// server
	srvr := server.NewServer(server.Args{
		RouterType: server.Mux,
		Service:    srvc,
		Cache:      c,
		Tracer:     tracer,
		Log:        log,
	})

	// starting server
	log.Info("server is running")

	server.MustServe(ctx, srvr, cfg.Port)

	log.Info("server stopped")
}
