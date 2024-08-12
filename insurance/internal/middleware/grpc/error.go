package grpc

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"github.com/alserov/car_insurance/insurance/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const internalError = "internal error"

func WithErrorHandler() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		res, err := handler(ctx, req)
		if err != nil {
			msg, st := utils.FromError(err)

			switch st {
			case utils.BadRequest:
				return nil, status.Error(codes.InvalidArgument, msg)
			case utils.Internal:
				logger.ExtractLogger(ctx).Error(err.Error())
				return nil, status.Error(codes.Internal, internalError)
			}

			logger.ExtractLogger(ctx).Error("unknown error type", logger.WithArg("error", err.Error()))

			return nil, status.Error(codes.Internal, internalError)
		}

		return res, nil
	}
}
