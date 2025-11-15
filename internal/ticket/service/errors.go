package service

import (
	"errors"
	"spotify/internal/ticket/repository/postgres"
)

var (
	ErrForbidden           = errors.New("forbidden")
	ErrNotFound            = errors.New("not found")
	ErrCannotUpdateTicket  = errors.New("ticket can no longer be updated")
	ErrCannotRateTicket    = errors.New("only closed tickets can be rated")
	ErrInvalidRating       = errors.New("rating must be between 1 and 5")
	ErrInvalidStatusChange = errors.New("user can only change status to 'Closed'")
	ErrCannotCloseTicket   = errors.New("only open tickets can be closed by user")
)

func mapError(err error) error {
	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return ErrNotFound
	default:
		return err
	}
}
