// Code generated by MockGen. DO NOT EDIT.
// Source: git.fogcdn.top/axe/protos/goout/auth (interfaces: AuthClient)

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	auth "git.fogcdn.top/axe/protos/goout/auth"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAuthClient is a mock of AuthClient interface
type MockAuthClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthClientMockRecorder
}

// MockAuthClientMockRecorder is the mock recorder for MockAuthClient
type MockAuthClientMockRecorder struct {
	mock *MockAuthClient
}

// NewMockAuthClient creates a new mock instance
func NewMockAuthClient(ctrl *gomock.Controller) *MockAuthClient {
	mock := &MockAuthClient{ctrl: ctrl}
	mock.recorder = &MockAuthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthClient) EXPECT() *MockAuthClientMockRecorder {
	return m.recorder
}

// Callback mocks base method
func (m *MockAuthClient) Callback(arg0 context.Context, arg1 *auth.CallbackRequest, arg2 ...grpc.CallOption) (*auth.CallbackResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Callback", varargs...)
	ret0, _ := ret[0].(*auth.CallbackResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Callback indicates an expected call of Callback
func (mr *MockAuthClientMockRecorder) Callback(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Callback", reflect.TypeOf((*MockAuthClient)(nil).Callback), varargs...)
}

// Check mocks base method
func (m *MockAuthClient) Check(arg0 context.Context, arg1 *auth.CheckRequest, arg2 ...grpc.CallOption) (*auth.CheckResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Check", varargs...)
	ret0, _ := ret[0].(*auth.CheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check
func (mr *MockAuthClientMockRecorder) Check(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockAuthClient)(nil).Check), varargs...)
}

// GetActionList mocks base method
func (m *MockAuthClient) GetActionList(arg0 context.Context, arg1 *auth.GetActionListRequest, arg2 ...grpc.CallOption) (*auth.GetActionListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetActionList", varargs...)
	ret0, _ := ret[0].(*auth.GetActionListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionList indicates an expected call of GetActionList
func (mr *MockAuthClientMockRecorder) GetActionList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionList", reflect.TypeOf((*MockAuthClient)(nil).GetActionList), varargs...)
}

// GetResourceList mocks base method
func (m *MockAuthClient) GetResourceList(arg0 context.Context, arg1 *auth.GetResourceListRequest, arg2 ...grpc.CallOption) (*auth.GetResourceListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourceList", varargs...)
	ret0, _ := ret[0].(*auth.GetResourceListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceList indicates an expected call of GetResourceList
func (mr *MockAuthClientMockRecorder) GetResourceList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceList", reflect.TypeOf((*MockAuthClient)(nil).GetResourceList), varargs...)
}

// GetUserInfo mocks base method
func (m *MockAuthClient) GetUserInfo(arg0 context.Context, arg1 *auth.GetUserInfoRequest, arg2 ...grpc.CallOption) (*auth.GetUserInfoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserInfo", varargs...)
	ret0, _ := ret[0].(*auth.GetUserInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo
func (mr *MockAuthClientMockRecorder) GetUserInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockAuthClient)(nil).GetUserInfo), varargs...)
}

// Login mocks base method
func (m *MockAuthClient) Login(arg0 context.Context, arg1 *auth.LoginRequest, arg2 ...grpc.CallOption) (*auth.LoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*auth.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockAuthClientMockRecorder) Login(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthClient)(nil).Login), varargs...)
}

// Logout mocks base method
func (m *MockAuthClient) Logout(arg0 context.Context, arg1 *auth.LogoutRequest, arg2 ...grpc.CallOption) (*auth.LogoutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Logout", varargs...)
	ret0, _ := ret[0].(*auth.LogoutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logout indicates an expected call of Logout
func (mr *MockAuthClientMockRecorder) Logout(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthClient)(nil).Logout), varargs...)
}
