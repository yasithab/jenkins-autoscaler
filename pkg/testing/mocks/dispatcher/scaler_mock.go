// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/dispatcher/dispatcher.go

// Package mock_dispatcher is a generated GoMock package.
package mock_dispatcher

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockScalerer is a mock of Scalerer interface.
type MockScalerer struct {
	ctrl     *gomock.Controller
	recorder *MockScalererMockRecorder
}

// MockScalererMockRecorder is the mock recorder for MockScalerer.
type MockScalererMockRecorder struct {
	mock *MockScalerer
}

// NewMockScalerer creates a new mock instance.
func NewMockScalerer(ctrl *gomock.Controller) *MockScalerer {
	mock := &MockScalerer{ctrl: ctrl}
	mock.recorder = &MockScalererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScalerer) EXPECT() *MockScalererMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockScalerer) Do(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Do", ctx)
}

// Do indicates an expected call of Do.
func (mr *MockScalererMockRecorder) Do(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockScalerer)(nil).Do), ctx)
}

// GC mocks base method.
func (m *MockScalerer) GC(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GC", ctx)
}

// GC indicates an expected call of GC.
func (mr *MockScalererMockRecorder) GC(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GC", reflect.TypeOf((*MockScalerer)(nil).GC), ctx)
}
