package fiber

import (
	"fmt"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewServer(srvc *service.Service) *server {
	return &server{
		srvc: srvc,
	}
}

type server struct {
	app  *fiber.App
	srvc *service.Service
}

func (s server) Serve(port string) {
	setupRoutes(s.app, newHandler(s.srvc))

	if err := s.app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		panic("failed to serve: " + err.Error())
	}
}

func (s server) Stop() error {
	return s.app.Shutdown()
}
