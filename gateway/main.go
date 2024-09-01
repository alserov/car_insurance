package main

import (
	"github.com/alserov/car_insurance/gateway/internal/app"
	"github.com/alserov/car_insurance/gateway/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
