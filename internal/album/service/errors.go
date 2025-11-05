package service

import (
	"errors"
	"spotify/internal/album/repository/postgres"
)

var (
	ErrNotFound = errors.New("album not found")
)

func mapError(err error) error {
	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return ErrNotFound
	default:
		return err
	}
}
