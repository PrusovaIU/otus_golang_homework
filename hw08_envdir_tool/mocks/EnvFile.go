// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EnvFile is an autogenerated mock type for the EnvFile type
type EnvFile struct {
	mock.Mock
}

type EnvFile_Expecter struct {
	mock *mock.Mock
}

func (_m *EnvFile) EXPECT() *EnvFile_Expecter {
	return &EnvFile_Expecter{mock: &_m.Mock}
}

// ReadLine provides a mock function with given fields:
func (_m *EnvFile) ReadLine() ([]byte, bool, error) {
	ret := _m.Called()

	var r0 []byte
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func() ([]byte, bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnvFile_ReadLine_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadLine'
type EnvFile_ReadLine_Call struct {
	*mock.Call
}

// ReadLine is a helper method to define mock.On call
func (_e *EnvFile_Expecter) ReadLine() *EnvFile_ReadLine_Call {
	return &EnvFile_ReadLine_Call{Call: _e.mock.On("ReadLine")}
}

func (_c *EnvFile_ReadLine_Call) Run(run func()) *EnvFile_ReadLine_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EnvFile_ReadLine_Call) Return(line []byte, isPrefix bool, err error) *EnvFile_ReadLine_Call {
	_c.Call.Return(line, isPrefix, err)
	return _c
}

func (_c *EnvFile_ReadLine_Call) RunAndReturn(run func() ([]byte, bool, error)) *EnvFile_ReadLine_Call {
	_c.Call.Return(run)
	return _c
}

// NewEnvFile creates a new instance of EnvFile. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEnvFile(t interface {
	mock.TestingT
	Cleanup(func())
}) *EnvFile {
	mock := &EnvFile{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
