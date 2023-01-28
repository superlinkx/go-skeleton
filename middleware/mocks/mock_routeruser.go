// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/superlinkx/go-skeleton/middleware (interfaces: RouterUser)

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRouterUser is a mock of RouterUser interface.
type MockRouterUser struct {
	ctrl     *gomock.Controller
	recorder *MockRouterUserMockRecorder
}

// MockRouterUserMockRecorder is the mock recorder for MockRouterUser.
type MockRouterUserMockRecorder struct {
	mock *MockRouterUser
}

// NewMockRouterUser creates a new mock instance.
func NewMockRouterUser(ctrl *gomock.Controller) *MockRouterUser {
	mock := &MockRouterUser{ctrl: ctrl}
	mock.recorder = &MockRouterUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouterUser) EXPECT() *MockRouterUserMockRecorder {
	return m.recorder
}

// Use mocks base method.
func (m *MockRouterUser) Use(arg0 ...func(http.Handler) http.Handler) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Use", varargs...)
}

// Use indicates an expected call of Use.
func (mr *MockRouterUserMockRecorder) Use(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockRouterUser)(nil).Use), arg0...)
}