// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BufferByteReader is an autogenerated mock type for the BufferByteReader type
type BufferByteReader struct {
	mock.Mock
}

type BufferByteReader_Expecter struct {
	mock *mock.Mock
}

func (_m *BufferByteReader) EXPECT() *BufferByteReader_Expecter {
	return &BufferByteReader_Expecter{mock: &_m.Mock}
}

// ReadByte provides a mock function with given fields:
func (_m *BufferByteReader) ReadByte() (byte, error) {
	ret := _m.Called()

	var r0 byte
	var r1 error
	if rf, ok := ret.Get(0).(func() (byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() byte); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(byte)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BufferByteReader_ReadByte_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadByte'
type BufferByteReader_ReadByte_Call struct {
	*mock.Call
}

// ReadByte is a helper method to define mock.On call
func (_e *BufferByteReader_Expecter) ReadByte() *BufferByteReader_ReadByte_Call {
	return &BufferByteReader_ReadByte_Call{Call: _e.mock.On("ReadByte")}
}

func (_c *BufferByteReader_ReadByte_Call) Run(run func()) *BufferByteReader_ReadByte_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BufferByteReader_ReadByte_Call) Return(_a0 byte, _a1 error) *BufferByteReader_ReadByte_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BufferByteReader_ReadByte_Call) RunAndReturn(run func() (byte, error)) *BufferByteReader_ReadByte_Call {
	_c.Call.Return(run)
	return _c
}

// NewBufferByteReader creates a new instance of BufferByteReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBufferByteReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *BufferByteReader {
	mock := &BufferByteReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
