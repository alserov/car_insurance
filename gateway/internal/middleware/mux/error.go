package mux

import (
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/utils"
	"net/http"
)

func WithErrorHandler(fn func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			msg, st := utils.FromError(err)
			if st == utils.Internal {
				w.WriteHeader(http.StatusInternalServerError)
				logger.ExtractLogger(r.Context()).Error(msg)
			}
		}
	}
}
