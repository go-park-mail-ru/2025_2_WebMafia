package postgres

import "errors"

var (
	ErrNotFound = errors.New("not_found")
	ErrConflict = errors.New("conflict")
	ErrInternal = errors.New("internal")
)
