package storage

import "errors"

var (
	ErrUploadFailed = errors.New("upload_failed")
	ErrDeleteFailed = errors.New("delete_failed")
	ErrURLFailed    = errors.New("url_generation_failed")
)
