package application

import (
	"errors"
)

var (
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrUsernameAlreadyTaken = errors.New("username already taken")
)

type ErrorType string

const (
	ValidationError_ErrorType ErrorType = "validation_error"
	InternalError_ErrorType   ErrorType = "internal_error"
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

// IsValidationError returns true if an input error is an Error with a ValidationError_ErrorType
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

// IsValidationError returns true if an input error is an Error with a InternalError_ErrorType
func IsInternalError(err error) bool {
	var e Error
	return errors.As(err, &e) && e.Type == InternalError_ErrorType
}
