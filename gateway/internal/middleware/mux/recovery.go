package mux

import (
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"net/http"
)

func WithRecovery(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.ExtractLogger(r.Context()).Error("panic recovery", logger.WithArg("error", err))
			}
		}()

		fn.ServeHTTP(w, r)
	})
}
