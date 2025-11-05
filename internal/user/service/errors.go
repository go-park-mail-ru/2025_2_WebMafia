package service

import (
	"errors"
	"fmt"
	"spotify/internal/user/repository/postgres"
)

var (
	ErrValidation = errors.New("validation_error")
	ErrNotFound   = errors.New("not_found")
	ErrConflict   = errors.New("conflict_error")
)

func mapRepositoryError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return fmt.Errorf("not found: %w", err)
	case errors.Is(err, postgres.ErrConflict):
		return fmt.Errorf("user already exists: %w", err)
	default:
		return fmt.Errorf("internal error: %w", err)
	}
}
