package utils

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Error struct {
	msg  string
	code int
}

func (e Error) Error() string {
	return e.msg
}

const (
	Internal = iota
	BadRequest
)

func NewError(msg string, code int) error {
	return &Error{
		msg:  msg,
		code: code,
	}
}

func FromError(in error) (string, int) {
	var e *Error
	errors.As(in, &e)

	switch e.code {
	case Internal:
		return "internal error", http.StatusInternalServerError
	case BadRequest:
		return e.msg, http.StatusBadRequest
	default:
		panic("unknown error: " + in.Error())
	}
}

func FromGRPCError(in error) error {
	res, ok := status.FromError(in)
	if !ok {
		return NewError("failed to get error from grpc error", Internal)
	}

	switch res.Code() {
	case codes.Internal:
		return NewError(res.Message(), Internal)
	case codes.InvalidArgument:
		return NewError(res.Message(), BadRequest)
	default:
		return NewError(fmt.Sprintf("unknown status: %s", res.Message()), Internal)
	}

}
