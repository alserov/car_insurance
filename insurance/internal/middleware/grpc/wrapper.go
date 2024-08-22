package grpc

import (
	"context"
	"google.golang.org/grpc"
)

func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.ServerOption {
	n := len(interceptors)

	if n == 0 {
		return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			return handler(ctx, req)
		})
	}

	if n == 1 {
		return grpc.UnaryInterceptor(interceptors[0])
	}

	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		currHandler := handler
		for i := n - 1; i > 0; i-- {
			innerHandler, i := currHandler, i
			currHandler = func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return interceptors[i](currentCtx, currentReq, info, innerHandler)
			}
		}

		return interceptors[0](ctx, req, info, currHandler)
	})
}
