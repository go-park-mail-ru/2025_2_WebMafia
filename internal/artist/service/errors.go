package service

import (
	"errors"
	"spotify/internal/artist/repository/postgres"
)

var (
	ErrNotFound = errors.New("artist not found")
)

func mapError(err error) error {
	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return ErrNotFound
	default:
		return err
	}
}
