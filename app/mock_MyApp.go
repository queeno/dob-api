// Code generated by mockery v1.0.0. DO NOT EDIT.

package app

import mock "github.com/stretchr/testify/mock"

// MockMyApp is an autogenerated mock type for the MyApp type
type MockMyApp struct {
	mock.Mock
}

// GetDateOfBirth provides a mock function with given fields: _a0
func (_m *MockMyApp) GetDateOfBirth(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUsername provides a mock function with given fields: _a0, _a1
func (_m *MockMyApp) UpdateUsername(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}