package mux

import (
	"github.com/alserov/car_insurance/gateway/internal/limiter"
	"net/http"
)

func WithLimiter(fn http.Handler, lim limiter.Limiter) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if lim.Limit(request.Context()) {
			return
		}

		fn.ServeHTTP(writer, request)
	})
}
