package postgres

import (
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

func mapErrors(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return err
	}
}
