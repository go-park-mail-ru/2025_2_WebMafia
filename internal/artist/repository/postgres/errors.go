package postgres

import "errors"

var (
	ErrNotFound = errors.New("artist not found in postgres repository")
)
