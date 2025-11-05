package postgres

import (
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("track not found in postgres repository")
)

func mapErrors(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return err
	}
}
