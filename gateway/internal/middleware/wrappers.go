package middleware

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/tracing"
	"go.opentelemetry.io/otel/trace"
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

func WithTracer(trace trace.Tracer) Wrapper {
	return func(ctx context.Context) context.Context {
		return tracing.WrapTracer(ctx, trace)
	}
}
