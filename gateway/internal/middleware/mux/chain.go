package mux

import "net/http"

func NewChain(mws ...http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, mw := range mws {
			mw.ServeHTTP(w, r)
		}
	})
}
