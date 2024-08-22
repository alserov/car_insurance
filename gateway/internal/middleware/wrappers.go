package middleware

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/logger"
)

type Wrapper func(ctx context.Context) context.Context

const (
	CtxKey = "ctx"
)

func WithLogger(log logger.Logger) Wrapper {
	return func(ctx context.Context) context.Context {
		return logger.WrapLogger(ctx, log)
	}
}
