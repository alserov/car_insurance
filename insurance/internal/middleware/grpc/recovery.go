package grpc

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"google.golang.org/grpc"
)

func WithRecovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if e := recover(); err != nil {
				logger.ExtractLogger(ctx).Error(
					"panic recovery",
					logger.WithArg("error", e),
				)
			}
		}()

		res, err := handler(ctx, req)

		return res, err
	}
}
