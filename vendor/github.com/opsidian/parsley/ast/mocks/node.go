// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import reader "github.com/opsidian/parsley/reader"

// Node is an autogenerated mock type for the Node type
type Node struct {
	mock.Mock
}

// Pos provides a mock function with given fields:
func (_m *Node) Pos() reader.Position {
	ret := _m.Called()

	var r0 reader.Position
	if rf, ok := ret.Get(0).(func() reader.Position); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reader.Position)
		}
	}

	return r0
}

// Token provides a mock function with given fields:
func (_m *Node) Token() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Value provides a mock function with given fields: ctx
func (_m *Node) Value(ctx interface{}) (interface{}, reader.Error) {
	ret := _m.Called(ctx)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 reader.Error
	if rf, ok := ret.Get(1).(func(interface{}) reader.Error); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(reader.Error)
		}
	}

	return r0, r1
}