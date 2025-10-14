package postgres

import "errors"

var (
	ErrNotFound = errors.New("track not found in postgres repository")
)
