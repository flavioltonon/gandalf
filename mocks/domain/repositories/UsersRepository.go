// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/flavioltonon/gandalf/domain/entity"
	mock "github.com/stretchr/testify/mock"

	valueobject "github.com/flavioltonon/gandalf/domain/valueobject"
)

// UsersRepository is an autogenerated mock type for the UsersRepository type
type UsersRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UsersRepository) CreateUser(ctx context.Context, user *entity.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByUsernameAndPassword provides a mock function with given fields: ctx, username, password
func (_m *UsersRepository) GetUserByUsernameAndPassword(ctx context.Context, username valueobject.Username, password valueobject.Password) (*entity.User, error) {
	ret := _m.Called(ctx, username, password)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, valueobject.Username, valueobject.Password) *entity.User); ok {
		r0 = rf(ctx, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, valueobject.Username, valueobject.Password) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUsersRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsersRepository creates a new instance of UsersRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersRepository(t mockConstructorTestingTNewUsersRepository) *UsersRepository {
	mock := &UsersRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
