// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IntInterface is an autogenerated mock type for the IntInterface type
type IntInterface struct {
	mock.Mock
}

type IntInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *IntInterface) EXPECT() *IntInterface_Expecter {
	return &IntInterface_Expecter{mock: &_m.Mock}
}

// Int provides a mock function with given fields:
func (_m *IntInterface) Int() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// IntInterface_Int_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Int'
type IntInterface_Int_Call struct {
	*mock.Call
}

// Int is a helper method to define mock.On call
func (_e *IntInterface_Expecter) Int() *IntInterface_Int_Call {
	return &IntInterface_Int_Call{Call: _e.mock.On("Int")}
}

func (_c *IntInterface_Int_Call) Run(run func()) *IntInterface_Int_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IntInterface_Int_Call) Return(_a0 int64) *IntInterface_Int_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IntInterface_Int_Call) RunAndReturn(run func() int64) *IntInterface_Int_Call {
	_c.Call.Return(run)
	return _c
}

// NewIntInterface creates a new instance of IntInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIntInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *IntInterface {
	mock := &IntInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
