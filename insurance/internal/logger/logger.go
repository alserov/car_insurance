package logger

import (
	"context"
)

type Logger interface {
	Info(msg string, arg ...Arg)
	Warn(msg string, arg ...Arg)
	Error(msg string, arg ...Arg)
	Debug(msg string, arg ...Arg)
}

type Arg struct {
	Key string
	Val any
}

func WithArg(key string, val any) Arg {
	return Arg{
		Key: key,
		Val: val,
	}
}

func WrapLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, CtxContextKey, l)
}

func ExtractLogger(ctx context.Context) Logger {
	l, ok := ctx.Value(CtxContextKey).(Logger)
	if !ok {
		panic("failed to extract logger")
	}

	return l
}

type ContextKey string

const (
	CtxContextKey ContextKey = "logger"
)

const (
	Zap = iota
)

const (
	Local = "local"
	Prod  = "prod"
)

func NewLogger(logType uint, env string) Logger {
	switch logType {
	case Zap:
		return newZap(env)
	default:
		panic("invalid logger type")
	}
}
