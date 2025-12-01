package service

import (
	"errors"
	"fmt"
	"spotify/microservices/playlist/repository/postgres"
)

var ErrNotFound = errors.New("not_found")

func mapRepositoryError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return fmt.Errorf("not found: %w", ErrNotFound)
	default:
		return fmt.Errorf("internal error: %w", err)
	}
}
