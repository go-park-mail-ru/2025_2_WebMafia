package service

import (
	"errors"
	"spotify/internal/ticket/repository/postgres"
)

var (
	ErrForbidden             = errors.New("forbidden")
	ErrNotFound              = errors.New("not found")
	ErrInvalidStateForAction = errors.New("invalid state for this action")
	ErrInvalidRating         = errors.New("rating must be between 1 and 5")
)

func mapError(err error) error {
	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return ErrNotFound
	default:
		return err
	}
}
