package grpc

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"google.golang.org/grpc"
)

func WithLogger(log logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx = logger.WrapLogger(ctx, log)

		return handler(ctx, req)
	}
}
