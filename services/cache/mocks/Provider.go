// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	cache "github.com/zgordan-vv/zacmm-server/services/cache"
	mock "github.com/stretchr/testify/mock"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Provider) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connect provides a mock function with given fields:
func (_m *Provider) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCache provides a mock function with given fields: opts
func (_m *Provider) NewCache(opts *cache.CacheOptions) cache.Cache {
	ret := _m.Called(opts)

	var r0 cache.Cache
	if rf, ok := ret.Get(0).(func(*cache.CacheOptions) cache.Cache); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Cache)
		}
	}

	return r0
}
