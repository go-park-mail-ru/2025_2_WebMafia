package service

import "errors"

var (
	ErrEmailExists        = errors.New("user with this email already exists")
	ErrLoginExists        = errors.New("user with this login already exists")
	ErrInvalidCredentials = errors.New("invalid login or password")
)
