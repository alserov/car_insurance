package mux

import (
	"fmt"
	"github.com/alserov/car_insurance/gateway/internal/middleware/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

var (
	v1Prefix = "/v1"

	insurancePrefix = "/insurance"

	swaggerRoute = fmt.Sprintf("%s/swagger/*", v1Prefix)

	insuranceInfoRoute   = fmt.Sprintf("GET %s/%s/info", v1Prefix, insurancePrefix)
	insuranceNewRoute    = fmt.Sprintf("POST %s/%s/new", v1Prefix, insurancePrefix)
	insurancePayoffRoute = fmt.Sprintf("POST %s/%s/payoff", v1Prefix, insurancePrefix)
)

func setupRoutes(r *http.ServeMux, h *handler) {
	r.HandleFunc(swaggerRoute, httpSwagger.WrapHandler)

	r.HandleFunc(insuranceInfoRoute, mux.WithErrorHandler(h.insurance.GetInsuranceData))
	r.HandleFunc(insuranceNewRoute, mux.WithErrorHandler(h.insurance.CreateInsurance))
	r.HandleFunc(insurancePayoffRoute, mux.WithErrorHandler(h.insurance.Payoff))
}
