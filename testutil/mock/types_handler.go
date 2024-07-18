// Code generated by MockGen.
// Source: types/handler.go
// Manual changes:
// + AnteHandler(...): calling `next` at the end of the function to run all ante handler chain.
// + PostHandler(...): calling `next` at the end of the function to run all post handler chain.

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/airchains-network/cosmos-sdk/types"
)

// MockAnteDecorator is a mock of AnteDecorator interface.
type MockAnteDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockAnteDecoratorMockRecorder
}

// MockAnteDecoratorMockRecorder is the mock recorder for MockAnteDecorator.
type MockAnteDecoratorMockRecorder struct {
	mock *MockAnteDecorator
}

// NewMockAnteDecorator creates a new mock instance.
func NewMockAnteDecorator(ctrl *gomock.Controller) *MockAnteDecorator {
	mock := &MockAnteDecorator{ctrl: ctrl}
	mock.recorder = &MockAnteDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnteDecorator) EXPECT() *MockAnteDecoratorMockRecorder {
	return m.recorder
}

// AnteHandle mocks base method.
func (m *MockAnteDecorator) AnteHandle(ctx types.Context, tx types.Tx, simulate bool, next types.AnteHandler) (types.Context, error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AnteHandle", ctx, tx, simulate, next)
	// NOTE: we need to edit a generated code to call the "next handler"
	return next(ctx, tx, simulate)
}

// AnteHandle indicates an expected call of AnteHandle.
func (mr *MockAnteDecoratorMockRecorder) AnteHandle(ctx, tx, simulate, next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnteHandle", reflect.TypeOf((*MockAnteDecorator)(nil).AnteHandle), ctx, tx, simulate, next)
}

// MockPostDecorator is a mock of PostDecorator interface.
type MockPostDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockPostDecoratorMockRecorder
}

// MockPostDecoratorMockRecorder is the mock recorder for MockPostDecorator.
type MockPostDecoratorMockRecorder struct {
	mock *MockPostDecorator
}

// NewMockPostDecorator creates a new mock instance.
func NewMockPostDecorator(ctrl *gomock.Controller) *MockPostDecorator {
	mock := &MockPostDecorator{ctrl: ctrl}
	mock.recorder = &MockPostDecoratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostDecorator) EXPECT() *MockPostDecoratorMockRecorder {
	return m.recorder
}

// PostHandle mocks base method.
func (m *MockPostDecorator) PostHandle(ctx types.Context, tx types.Tx, simulate, success bool, next types.PostHandler) (types.Context, error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostHandle", ctx, tx, simulate, success, next)
	// NOTE: we need to edit the generated code to call the "next handler"
	return next(ctx, tx, simulate, success)
}

// PostHandle indicates an expected call of PostHandle.
func (mr *MockPostDecoratorMockRecorder) PostHandle(ctx, tx, simulate, success, next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostHandle", reflect.TypeOf((*MockPostDecorator)(nil).PostHandle), ctx, tx, simulate, success, next)
}
