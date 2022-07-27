package domain

import "errors"

var (
	// ErrAlreadyExists is an error that should be used when an entity was expected not to exist, but it does
	ErrAlreadyExists = errors.New("already exists")

	// ErrNotFound is an error that should be used when an entity was expected to exist, but it does not
	ErrNotFound = errors.New("not found")
)
