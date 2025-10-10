package service

import (
	"errors"
)

type ErrorType string

var (
	ErrValidation = errors.New("validation_error")
	ErrConflict   = errors.New("conflict_error")
	ErrInternal   = errors.New("internal_error")
)

type IsServiceError struct {
	Message string
	Err     error
}

func (e *IsServiceError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func NewServiceError(err error, msg string) *IsServiceError {
	return &IsServiceError{Message: msg, Err: err}
}

func (e *IsServiceError) Unwrap() error {
	return e.Err
}
