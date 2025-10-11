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
