package handler

import (
	"errors"
	"fmt"
	"runtime/debug"
)

type CustomError struct {
	Error error `json:"error"`
	Message string `json:"message"`
}

var (
	ErrInvalidFileFormat = errors.New("invalid file format")
)

func (h *Handler) Error(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	h.errLogger.Output(2, trace)
}