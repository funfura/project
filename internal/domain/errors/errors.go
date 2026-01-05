package errors

import "errors"

var (
	ErrNotFound        = errors.New("entity not found")
	ErrAlreadyComleted = errors.New("already completed")
	errForbidden       = errors.New("forbidden")
)
