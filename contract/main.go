package main

import (
	"github.com/alserov/car_insurance/contract/internal/app"
	"github.com/alserov/car_insurance/contract/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
