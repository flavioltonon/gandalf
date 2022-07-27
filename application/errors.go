package application

import (
	"errors"
)

var (
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrUsernameAlreadyTaken = errors.New("username already taken")
)

type ErrorType int

const (
	_ ErrorType = iota
	ValidationError_ErrorType
	InternalError_ErrorType
)

type Error struct {
	Type    ErrorType
	Message string
}

func (e Error) Error() string { return e.Message }

// ValidationError is a wrapper for errors created from data validation
func ValidationError(err error) error {
	return Error{
		Type:    ValidationError_ErrorType,
		Message: err.Error(),
	}
}

func IsValidationError(err error) bool {
	var e Error
	return errors.As(err, &e) && e.Type == ValidationError_ErrorType
}

// InternalError is an unexpected, internal error
func InternalError(err error) error {
	return Error{
		Type:    InternalError_ErrorType,
		Message: err.Error(),
	}
}

func IsInternalError(err error) bool {
	var e Error
	return errors.As(err, &e) && e.Type == InternalError_ErrorType
}
