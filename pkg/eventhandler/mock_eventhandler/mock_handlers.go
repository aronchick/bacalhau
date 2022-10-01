// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_eventhandler is a generated GoMock package.
package mock_eventhandler

import (
	context "context"
	reflect "reflect"

	model "github.com/filecoin-project/bacalhau/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockLocalEventHandler is a mock of LocalEventHandler interface.
type MockLocalEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockLocalEventHandlerMockRecorder
}

// MockLocalEventHandlerMockRecorder is the mock recorder for MockLocalEventHandler.
type MockLocalEventHandlerMockRecorder struct {
	mock *MockLocalEventHandler
}

// NewMockLocalEventHandler creates a new mock instance.
func NewMockLocalEventHandler(ctrl *gomock.Controller) *MockLocalEventHandler {
	mock := &MockLocalEventHandler{ctrl: ctrl}
	mock.recorder = &MockLocalEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalEventHandler) EXPECT() *MockLocalEventHandlerMockRecorder {
	return m.recorder
}

// HandleLocalEvent mocks base method.
func (m *MockLocalEventHandler) HandleLocalEvent(ctx context.Context, event model.JobLocalEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleLocalEvent", ctx, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleLocalEvent indicates an expected call of HandleLocalEvent.
func (mr *MockLocalEventHandlerMockRecorder) HandleLocalEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleLocalEvent", reflect.TypeOf((*MockLocalEventHandler)(nil).HandleLocalEvent), ctx, event)
}

// MockJobEventHandler is a mock of JobEventHandler interface.
type MockJobEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockJobEventHandlerMockRecorder
}

// MockJobEventHandlerMockRecorder is the mock recorder for MockJobEventHandler.
type MockJobEventHandlerMockRecorder struct {
	mock *MockJobEventHandler
}

// NewMockJobEventHandler creates a new mock instance.
func NewMockJobEventHandler(ctrl *gomock.Controller) *MockJobEventHandler {
	mock := &MockJobEventHandler{ctrl: ctrl}
	mock.recorder = &MockJobEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobEventHandler) EXPECT() *MockJobEventHandlerMockRecorder {
	return m.recorder
}

// HandleJobEvent mocks base method.
func (m *MockJobEventHandler) HandleJobEvent(ctx context.Context, event model.JobEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleJobEvent", ctx, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleJobEvent indicates an expected call of HandleJobEvent.
func (mr *MockJobEventHandlerMockRecorder) HandleJobEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleJobEvent", reflect.TypeOf((*MockJobEventHandler)(nil).HandleJobEvent), ctx, event)
}