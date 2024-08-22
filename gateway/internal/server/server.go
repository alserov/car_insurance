package server

import (
	"github.com/alserov/car_insurance/gateway/internal/server/fiber"
	"github.com/alserov/car_insurance/gateway/internal/service"
)

type Server interface {
	Serve(port string)
	Stop() error
}

const (
	Fiber = iota
)

func NewServer(routerType uint, srvc *service.Service) Server {
	switch routerType {
	case Fiber:
		return fiber.NewServer(srvc)
	default:
		panic("invalid router type")
	}
}
