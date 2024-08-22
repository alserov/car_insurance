package logger

import (
	"go.uber.org/zap"
)

type log struct {
	l *zap.Logger
}

func (l log) Info(msg string, arg ...Arg) {
	if len(arg) > 0 {
		l.l.Info(msg, zap.Any(arg[0].Key, arg[0].Val))
	} else {
		l.l.Info(msg)
	}
}

func (l log) Warn(msg string, arg ...Arg) {
	if len(arg) > 0 {
		l.l.Warn(msg, zap.Any(arg[0].Key, arg[0].Val))
	} else {
		l.l.Warn(msg)
	}
}

func (l log) Error(msg string, arg ...Arg) {
	if len(arg) > 0 {
		l.l.Error(msg, zap.Any(arg[0].Key, arg[0].Val))
	} else {
		l.l.Error(msg)
	}
}

func (l log) Debug(msg string, arg ...Arg) {
	if len(arg) > 0 {
		l.l.Debug(msg, zap.Any(arg[0].Key, arg[0].Val))
	} else {
		l.l.Debug(msg)
	}
}

func newZap(env string) Logger {
	var (
		l   *zap.Logger
		err error
	)

	switch env {
	case Local:
		l, err = zap.NewDevelopment()
	case Prod:
		l, err = zap.NewProduction()
	}

	if err != nil {
		panic("failed to init logger: " + err.Error())
	}

	return &log{l: l}
}
