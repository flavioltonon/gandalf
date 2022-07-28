// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Presenter is an autogenerated mock type for the Presenter type
type Presenter struct {
	mock.Mock
}

// Present provides a mock function with given fields: rw, statusCode, data
func (_m *Presenter) Present(rw http.ResponseWriter, statusCode int, data interface{}) error {
	ret := _m.Called(rw, statusCode, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, int, interface{}) error); ok {
		r0 = rf(rw, statusCode, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPresenter interface {
	mock.TestingT
	Cleanup(func())
}

// NewPresenter creates a new instance of Presenter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPresenter(t mockConstructorTestingTNewPresenter) *Presenter {
	mock := &Presenter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}