// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/superlinkx/go-skeleton/handlers (interfaces: RouterHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRouterHandler is a mock of RouterHandler interface.
type MockRouterHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRouterHandlerMockRecorder
}

// MockRouterHandlerMockRecorder is the mock recorder for MockRouterHandler.
type MockRouterHandlerMockRecorder struct {
	mock *MockRouterHandler
}

// NewMockRouterHandler creates a new mock instance.
func NewMockRouterHandler(ctrl *gomock.Controller) *MockRouterHandler {
	mock := &MockRouterHandler{ctrl: ctrl}
	mock.recorder = &MockRouterHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouterHandler) EXPECT() *MockRouterHandlerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRouterHandler) Get(arg0 string, arg1 http.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get", arg0, arg1)
}

// Get indicates an expected call of Get.
func (mr *MockRouterHandlerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRouterHandler)(nil).Get), arg0, arg1)
}

// Post mocks base method.
func (m *MockRouterHandler) Post(arg0 string, arg1 http.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Post", arg0, arg1)
}

// Post indicates an expected call of Post.
func (mr *MockRouterHandlerMockRecorder) Post(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockRouterHandler)(nil).Post), arg0, arg1)
}
