// Code generated by MockGen. DO NOT EDIT.
// Source: git.fogcdn.top/axe/protos/goout/image (interfaces: ImageClient)

// Package mock_image is a generated GoMock package.
package mock_image

import (
	context "context"
	image "git.fogcdn.top/axe/protos/goout/image"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockImageClient is a mock of ImageClient interface
type MockImageClient struct {
	ctrl     *gomock.Controller
	recorder *MockImageClientMockRecorder
}

// MockImageClientMockRecorder is the mock recorder for MockImageClient
type MockImageClientMockRecorder struct {
	mock *MockImageClient
}

// NewMockImageClient creates a new mock instance
func NewMockImageClient(ctrl *gomock.Controller) *MockImageClient {
	mock := &MockImageClient{ctrl: ctrl}
	mock.recorder = &MockImageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageClient) EXPECT() *MockImageClientMockRecorder {
	return m.recorder
}

// Health mocks base method
func (m *MockImageClient) Health(arg0 context.Context, arg1 *image.HealthRequest, arg2 ...grpc.CallOption) (*image.HealthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Health", varargs...)
	ret0, _ := ret[0].(*image.HealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health
func (mr *MockImageClientMockRecorder) Health(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockImageClient)(nil).Health), varargs...)
}

// Query mocks base method
func (m *MockImageClient) Query(arg0 context.Context, arg1 *image.QueryRequest, arg2 ...grpc.CallOption) (*image.QueryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*image.QueryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockImageClientMockRecorder) Query(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockImageClient)(nil).Query), varargs...)
}