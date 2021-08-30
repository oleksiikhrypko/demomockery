// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DataProvider is an autogenerated mock type for the DataProvider type
type DataProvider struct {
	mock.Mock
}

// GetData provides a mock function with given fields: idx
func (_m *DataProvider) GetData(idx int) (int, error) {
	ret := _m.Called(idx)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(idx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}