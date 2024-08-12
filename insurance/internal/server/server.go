package server

import (
	"context"
	"os/signal"
	"syscall"
)

type Server interface {
	Serve(host string) error
}

func MustServe(srv Server, port string) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	cancel()

	go func() {
		if err := srv.Serve(port); err != nil {
			panic("failed to start server: " + err.Error())
		}
	}()

	<-ctx.Done()
}
