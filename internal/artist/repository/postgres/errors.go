package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("artist not found in postgres repository")
)

func mapErrors(err error, location string) error {

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return fmt.Errorf("%s: %w", location, err)
	}
}
