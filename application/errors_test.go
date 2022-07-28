package application

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationError(t *testing.T) {
	err := errors.New("some error")

	type args struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Given an error, a new ValidationError should be created",
			args: args{
				err: err,
			},
			want: Error{
				Type:    ValidationError_ErrorType,
				Message: err.Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ValidationError(tt.args.err))
		})
	}
}

func TestIsValidationError(t *testing.T) {
	type args struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "IsValidationError should return false if the input error is not an Error",
			args: args{
				err: nil,
			},
			want: false,
		},
		{
			name: "IsValidationError should return false if the input error is an Error, but doesn't have a ValidationError_ErrorType",
			args: args{
				err: Error{
					Type:    InternalError_ErrorType,
					Message: "some error",
				},
			},
			want: false,
		},
		{
			name: "IsValidationError should return true if the input error is an Error and it has a ValidationError_ErrorType",
			args: args{
				err: Error{
					Type:    ValidationError_ErrorType,
					Message: "some error",
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsValidationError(tt.args.err))
		})
	}
}

func TestInternalError(t *testing.T) {
	err := errors.New("some error")

	type args struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Given an error, a new InternalError should be created",
			args: args{
				err: err,
			},
			want: Error{
				Type:    InternalError_ErrorType,
				Message: err.Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, InternalError(tt.args.err))
		})
	}
}

func TestIsInternalError(t *testing.T) {
	type args struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "IsInternalError should return false if the input error is not an Error",
			args: args{
				err: nil,
			},
			want: false,
		},
		{
			name: "IsInternalError should return false if the input error is an Error, but doesn't have a InternalError_ErrorType",
			args: args{
				err: Error{
					Type:    ValidationError_ErrorType,
					Message: "some error",
				},
			},
			want: false,
		},
		{
			name: "IsInternalError should return true if the input error is an Error and it has a InternalError_ErrorType",
			args: args{
				err: Error{
					Type:    InternalError_ErrorType,
					Message: "some error",
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsInternalError(tt.args.err))
		})
	}
}
