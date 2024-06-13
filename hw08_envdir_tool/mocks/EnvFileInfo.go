// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EnvFileInfo is an autogenerated mock type for the EnvFileInfo type
type EnvFileInfo struct {
	mock.Mock
}

type EnvFileInfo_Expecter struct {
	mock *mock.Mock
}

func (_m *EnvFileInfo) EXPECT() *EnvFileInfo_Expecter {
	return &EnvFileInfo_Expecter{mock: &_m.Mock}
}

// Name provides a mock function with given fields:
func (_m *EnvFileInfo) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// EnvFileInfo_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type EnvFileInfo_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *EnvFileInfo_Expecter) Name() *EnvFileInfo_Name_Call {
	return &EnvFileInfo_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *EnvFileInfo_Name_Call) Run(run func()) *EnvFileInfo_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EnvFileInfo_Name_Call) Return(_a0 string) *EnvFileInfo_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EnvFileInfo_Name_Call) RunAndReturn(run func() string) *EnvFileInfo_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewEnvFileInfo creates a new instance of EnvFileInfo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEnvFileInfo(t interface {
	mock.TestingT
	Cleanup(func())
}) *EnvFileInfo {
	mock := &EnvFileInfo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
