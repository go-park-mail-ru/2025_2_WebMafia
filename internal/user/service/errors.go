package service

import "fmt"

type ErrorType string

const (
	ErrValidation ErrorType = "validation_error"
	ErrConflict   ErrorType = "conflict_error"
	ErrInternal   ErrorType = "internal_error"
)

type IsServiceError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *IsServiceError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func WrapError(err error, typ ErrorType, msg string) *IsServiceError {
	return &IsServiceError{Type: typ, Message: msg, Err: err}
}
