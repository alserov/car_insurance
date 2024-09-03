package main

import (
	"github.com/alserov/car_insurance/insurance/internal/app"
	"github.com/alserov/car_insurance/insurance/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
