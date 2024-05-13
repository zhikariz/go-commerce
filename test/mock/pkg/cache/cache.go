// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/cache/cache.go
//
// Generated by this command:
//
//	mockgen -source=./pkg/cache/cache.go -destination=test/mock/./pkg/cache/cache.go
//

// Package mock_cache is a generated GoMock package.
package mock_cache

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockCacheable is a mock of Cacheable interface.
type MockCacheable struct {
	ctrl     *gomock.Controller
	recorder *MockCacheableMockRecorder
}

// MockCacheableMockRecorder is the mock recorder for MockCacheable.
type MockCacheableMockRecorder struct {
	mock *MockCacheable
}

// NewMockCacheable creates a new mock instance.
func NewMockCacheable(ctrl *gomock.Controller) *MockCacheable {
	mock := &MockCacheable{ctrl: ctrl}
	mock.recorder = &MockCacheableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheable) EXPECT() *MockCacheableMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCacheable) Get(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCacheableMockRecorder) Get(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheable)(nil).Get), key)
}

// Set mocks base method.
func (m *MockCacheable) Set(key string, value any, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheableMockRecorder) Set(key, value, expiration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacheable)(nil).Set), key, value, expiration)
}
