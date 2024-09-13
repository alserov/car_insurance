package mux

import (
	"fmt"
	"github.com/alserov/car_insurance/gateway/internal/cache"
	"github.com/alserov/car_insurance/gateway/internal/limiter"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/middleware"
	"github.com/alserov/car_insurance/gateway/internal/middleware/mux"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"go.opentelemetry.io/otel/trace"
	"net"
	"net/http"
)

func NewServer(srvc *service.Service, cache cache.Cache, tracer trace.Tracer, log logger.Logger) *server {
	return &server{
		app:   http.NewServeMux(),
		srvc:  srvc,
		cache: cache,
		trace: tracer,
		log:   log,
	}
}

type server struct {
	app *http.ServeMux
	lis net.Listener

	cache cache.Cache
	srvc  *service.Service
	trace trace.Tracer
	log   logger.Logger
}

func (s server) Serve(port string) error {
	setupRoutes(s.app, newHandler(s.srvc, s.cache))

	s.app.Handle("/",
		mux.NewChain(
			mux.WithRecovery(s.app),
			mux.WithWrappers(
				middleware.WithTracer(s.trace),
				middleware.WithLogger(s.log),
			)(s.app),
			mux.WithLimiter(s.app, limiter.NewLimiter(limiter.LeakyBucket, limiter.DefaultLimit)),
		),
	)

	srvr := &http.Server{
		Handler: s.app,
	}

	var err error
	s.lis, err = net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	if err = srvr.Serve(s.lis); err != nil {
		return err
	}

	return nil
}

func (s server) Shutdown() error {
	return s.lis.Close()
}
