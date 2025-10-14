package postgres

import "errors"

var (
	ErrNotFound = errors.New("album not found in postgres repository")
)
