package server

import (
	"context"
)

type Server interface {
	Serve(host string) error
	Shutdown() error
}

func MustServe(ctx context.Context, srv Server, port string) {
	go func() {
		if err := srv.Serve(port); err != nil {
			panic("failed to start server: " + err.Error())
		}
	}()

	<-ctx.Done()

	_ = srv.Shutdown()
}
