package service

import (
	"errors"
	"spotify/internal/track/repository/postgres"
)

var (
	ErrNotFound = errors.New("track not found")
)

func mapError(err error) error {
	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return ErrNotFound
	default:
		return err
	}
}
