package service

import (
	"errors"
	"fmt"
	"spotify/internal/user/repository/postgres"
)

func mapRepositoryError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return fmt.Errorf("not found: %w", ErrValidation)
	case errors.Is(err, postgres.ErrConflict):
		return fmt.Errorf("user already exists: %w", postgres.ErrConflict)
	default:
		return fmt.Errorf("internal error: %w", postgres.ErrInternal)
	}
}
