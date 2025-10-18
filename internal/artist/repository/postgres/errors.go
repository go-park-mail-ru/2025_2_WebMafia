package postgres

import (
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("artist not found in postgres repository")
)

func mapErrors(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return err
	}
}
