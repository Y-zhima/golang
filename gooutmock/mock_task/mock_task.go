// Code generated by MockGen. DO NOT EDIT.
// Source: git.fogcdn.top/axe/protos/goout/task (interfaces: TaskClient)

// Package mock_task is a generated GoMock package.
package mock_task

import (
	context "context"
	task "git.fogcdn.top/axe/protos/goout/task"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockTaskClient is a mock of TaskClient interface
type MockTaskClient struct {
	ctrl     *gomock.Controller
	recorder *MockTaskClientMockRecorder
}

// MockTaskClientMockRecorder is the mock recorder for MockTaskClient
type MockTaskClientMockRecorder struct {
	mock *MockTaskClient
}

// NewMockTaskClient creates a new mock instance
func NewMockTaskClient(ctrl *gomock.Controller) *MockTaskClient {
	mock := &MockTaskClient{ctrl: ctrl}
	mock.recorder = &MockTaskClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskClient) EXPECT() *MockTaskClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTaskClient) Create(arg0 context.Context, arg1 *task.CreateRequest, arg2 ...grpc.CallOption) (*task.CreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*task.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTaskClientMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskClient)(nil).Create), varargs...)
}

// Filter mocks base method
func (m *MockTaskClient) Filter(arg0 context.Context, arg1 *task.FilterRequest, arg2 ...grpc.CallOption) (*task.FilterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].(*task.FilterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *MockTaskClientMockRecorder) Filter(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockTaskClient)(nil).Filter), varargs...)
}

// Get mocks base method
func (m *MockTaskClient) Get(arg0 context.Context, arg1 *task.GetRequest, arg2 ...grpc.CallOption) (*task.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*task.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTaskClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTaskClient)(nil).Get), varargs...)
}

// GetLog mocks base method
func (m *MockTaskClient) GetLog(arg0 context.Context, arg1 *task.GetLogRequest, arg2 ...grpc.CallOption) (*task.GetLogResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLog", varargs...)
	ret0, _ := ret[0].(*task.GetLogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLog indicates an expected call of GetLog
func (mr *MockTaskClientMockRecorder) GetLog(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLog", reflect.TypeOf((*MockTaskClient)(nil).GetLog), varargs...)
}