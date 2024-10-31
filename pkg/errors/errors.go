package errors

import "errors"

var (
	ErrAlreadyExists  = errors.New("already exists")
	ErrAlreadyStarted = errors.New("already started")
	ErrNotFound       = errors.New("not found")
	ErrNotStarted     = errors.New("not started")
)
