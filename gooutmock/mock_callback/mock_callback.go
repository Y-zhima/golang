// Code generated by MockGen. DO NOT EDIT.
// Source: git.fogcdn.top/axe/protos/goout/callback (interfaces: CallbackClient)

// Package mock_callback is a generated GoMock package.
package mock_callback

import (
	context "context"
	callback "git.fogcdn.top/axe/protos/goout/callback"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockCallbackClient is a mock of CallbackClient interface
type MockCallbackClient struct {
	ctrl     *gomock.Controller
	recorder *MockCallbackClientMockRecorder
}

// MockCallbackClientMockRecorder is the mock recorder for MockCallbackClient
type MockCallbackClientMockRecorder struct {
	mock *MockCallbackClient
}

// NewMockCallbackClient creates a new mock instance
func NewMockCallbackClient(ctrl *gomock.Controller) *MockCallbackClient {
	mock := &MockCallbackClient{ctrl: ctrl}
	mock.recorder = &MockCallbackClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCallbackClient) EXPECT() *MockCallbackClientMockRecorder {
	return m.recorder
}

// CmdbEvent mocks base method
func (m *MockCallbackClient) CmdbEvent(arg0 context.Context, arg1 *callback.CmdbEventRequest, arg2 ...grpc.CallOption) (*callback.CmdbEventResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CmdbEvent", varargs...)
	ret0, _ := ret[0].(*callback.CmdbEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CmdbEvent indicates an expected call of CmdbEvent
func (mr *MockCallbackClientMockRecorder) CmdbEvent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdbEvent", reflect.TypeOf((*MockCallbackClient)(nil).CmdbEvent), varargs...)
}
