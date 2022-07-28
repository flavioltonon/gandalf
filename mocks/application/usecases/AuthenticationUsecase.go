// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/flavioltonon/gandalf/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// AuthenticationUsecase is an autogenerated mock type for the AuthenticationUsecase type
type AuthenticationUsecase struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, username, password
func (_m *AuthenticationUsecase) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	ret := _m.Called(ctx, username, password)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *entity.User); ok {
		r0 = rf(ctx, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, username, password
func (_m *AuthenticationUsecase) RegisterUser(ctx context.Context, username string, password string) (*entity.User, error) {
	ret := _m.Called(ctx, username, password)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *entity.User); ok {
		r0 = rf(ctx, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthenticationUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthenticationUsecase creates a new instance of AuthenticationUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthenticationUsecase(t mockConstructorTestingTNewAuthenticationUsecase) *AuthenticationUsecase {
	mock := &AuthenticationUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}