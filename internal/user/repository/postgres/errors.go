package postgres

import "errors"

var (
	ErrNotFound      = errors.New("not_found")
	ErrConflict      = errors.New("conflict")
	ErrQueryFailed   = errors.New("query_failed")
	ErrScanFailed    = errors.New("scan_failed")
	ErrRowsIteration = errors.New("rows_iteration_failed")
	ErrCreateFailed  = errors.New("create_failed")
	ErrInternal      = errors.New("internal")
)
