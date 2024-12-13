// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	types_validators "github.com/PrusovaIU/otus_golang_homework/hw09_struct_validator/types_validators"
	mock "github.com/stretchr/testify/mock"
)

// IntValidatorInterface is an autogenerated mock type for the IntValidatorInterface type
type IntValidatorInterface struct {
	mock.Mock
}

type IntValidatorInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *IntValidatorInterface) EXPECT() *IntValidatorInterface_Expecter {
	return &IntValidatorInterface_Expecter{mock: &_m.Mock}
}

// Validate provides a mock function with given fields: _a0, _a1, _a2
func (_m *IntValidatorInterface) Validate(_a0 types_validators.IntInterface, _a1 string, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(types_validators.IntInterface, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IntValidatorInterface_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type IntValidatorInterface_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - _a0 types_validators.IntInterface
//   - _a1 string
//   - _a2 string
func (_e *IntValidatorInterface_Expecter) Validate(_a0 interface{}, _a1 interface{}, _a2 interface{}) *IntValidatorInterface_Validate_Call {
	return &IntValidatorInterface_Validate_Call{Call: _e.mock.On("Validate", _a0, _a1, _a2)}
}

func (_c *IntValidatorInterface_Validate_Call) Run(run func(_a0 types_validators.IntInterface, _a1 string, _a2 string)) *IntValidatorInterface_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types_validators.IntInterface), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *IntValidatorInterface_Validate_Call) Return(_a0 error) *IntValidatorInterface_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IntValidatorInterface_Validate_Call) RunAndReturn(run func(types_validators.IntInterface, string, string) error) *IntValidatorInterface_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewIntValidatorInterface creates a new instance of IntValidatorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIntValidatorInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *IntValidatorInterface {
	mock := &IntValidatorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}