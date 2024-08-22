package fiber

import (
	"fmt"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/middleware"
	fiber_mw "github.com/alserov/car_insurance/gateway/internal/middleware/fiber"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewServer(srvc *service.Service, log logger.Logger) *server {
	return &server{
		app:  fiber.New(),
		srvc: srvc,
		log:  log,
	}
}

type server struct {
	app  *fiber.App
	srvc *service.Service
	log  logger.Logger
}

func (s server) Serve(port string) error {
	setupRoutes(s.app, newHandler(s.srvc))

	s.app.Use(
		fiber_mw.WithErrorHandler,
		fiber_mw.WithRecovery,
		fiber_mw.WithWrappers(
			middleware.WithLogger(s.log),
		),
	)

	if err := s.app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		return err
	}

	return nil
}

func (s server) Shutdown() error {
	return s.app.Shutdown()
}
