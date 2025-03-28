// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	errors "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/errors"

	mock "github.com/stretchr/testify/mock"

	reflect "reflect"
)

// SliceValidatorInterface is an autogenerated mock type for the SliceValidatorInterface type
type SliceValidatorInterface struct {
	mock.Mock
}

type SliceValidatorInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *SliceValidatorInterface) EXPECT() *SliceValidatorInterface_Expecter {
	return &SliceValidatorInterface_Expecter{mock: &_m.Mock}
}

// Validate provides a mock function with given fields: _a0, _a1, _a2
func (_m *SliceValidatorInterface) Validate(_a0 reflect.Value, _a1 reflect.StructField, _a2 string) (errors.ValidationErrors, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 errors.ValidationErrors
	var r1 error
	if rf, ok := ret.Get(0).(func(reflect.Value, reflect.StructField, string) (errors.ValidationErrors, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(reflect.Value, reflect.StructField, string) errors.ValidationErrors); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.ValidationErrors)
		}
	}

	if rf, ok := ret.Get(1).(func(reflect.Value, reflect.StructField, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SliceValidatorInterface_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type SliceValidatorInterface_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - _a0 reflect.Value
//   - _a1 reflect.StructField
//   - _a2 string
func (_e *SliceValidatorInterface_Expecter) Validate(_a0 interface{}, _a1 interface{}, _a2 interface{}) *SliceValidatorInterface_Validate_Call {
	return &SliceValidatorInterface_Validate_Call{Call: _e.mock.On("Validate", _a0, _a1, _a2)}
}

func (_c *SliceValidatorInterface_Validate_Call) Run(run func(_a0 reflect.Value, _a1 reflect.StructField, _a2 string)) *SliceValidatorInterface_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(reflect.Value), args[1].(reflect.StructField), args[2].(string))
	})
	return _c
}

func (_c *SliceValidatorInterface_Validate_Call) Return(_a0 errors.ValidationErrors, _a1 error) *SliceValidatorInterface_Validate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SliceValidatorInterface_Validate_Call) RunAndReturn(run func(reflect.Value, reflect.StructField, string) (errors.ValidationErrors, error)) *SliceValidatorInterface_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewSliceValidatorInterface creates a new instance of SliceValidatorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSliceValidatorInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SliceValidatorInterface {
	mock := &SliceValidatorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
