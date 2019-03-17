// Code generated by mockery v1.0.0. DO NOT EDIT.

package db

import mock "github.com/stretchr/testify/mock"

// MockDatabase is an autogenerated mock type for the Database type
type MockDatabase struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockDatabase) Close() {
	_m.Called()
}

// Get provides a mock function with given fields: _a0
func (_m *MockDatabase) Get(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Open provides a mock function with given fields:
func (_m *MockDatabase) Open() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Put provides a mock function with given fields: _a0, _a1
func (_m *MockDatabase) Put(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
