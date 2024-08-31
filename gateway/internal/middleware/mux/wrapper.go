package mux

import (
	"github.com/alserov/car_insurance/gateway/internal/middleware"
	"net/http"
)

func WithWrappers(wrs ...middleware.Wrapper) func(fn http.Handler) http.Handler {
	return func(fn http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			for _, wr := range wrs {
				ctx = wr(ctx)
			}

			fn.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
