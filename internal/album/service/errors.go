package service

import (
	"errors"
	"fmt"
	"spotify/internal/album/repository/postgres"
)

var (
	ErrNotFound = errors.New("album not found")
)

func mapError(err error, location string) error {
	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return ErrNotFound
	default:
		return fmt.Errorf("%s: %w", location, err)
	}
}
