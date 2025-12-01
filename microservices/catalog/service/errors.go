package service

import (
	"errors"

	"spotify/microservices/catalog/repository/postgres"
)

var ErrNotFound = errors.New("entity not found")

func mapError(err error) error {
	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return ErrNotFound
	default:
		return err
	}
}
