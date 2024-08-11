package utils

import (
	"errors"
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
