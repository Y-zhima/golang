// Code generated by MockGen. DO NOT EDIT.
// Source: git.fogcdn.top/axe/protos/goout/cmdb (interfaces: CmdbClient)

// Package mock_cmdb is a generated GoMock package.
package mock_cmdb

import (
	context "context"
	cmdb "git.fogcdn.top/axe/protos/goout/cmdb"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockCmdbClient is a mock of CmdbClient interface
type MockCmdbClient struct {
	ctrl     *gomock.Controller
	recorder *MockCmdbClientMockRecorder
}

// MockCmdbClientMockRecorder is the mock recorder for MockCmdbClient
type MockCmdbClientMockRecorder struct {
	mock *MockCmdbClient
}

// NewMockCmdbClient creates a new mock instance
func NewMockCmdbClient(ctrl *gomock.Controller) *MockCmdbClient {
	mock := &MockCmdbClient{ctrl: ctrl}
	mock.recorder = &MockCmdbClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCmdbClient) EXPECT() *MockCmdbClientMockRecorder {
	return m.recorder
}

// ImportHost mocks base method
func (m *MockCmdbClient) ImportHost(arg0 context.Context, arg1 *cmdb.ImportHostRequest, arg2 ...grpc.CallOption) (*cmdb.ImportHostResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportHost", varargs...)
	ret0, _ := ret[0].(*cmdb.ImportHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportHost indicates an expected call of ImportHost
func (mr *MockCmdbClientMockRecorder) ImportHost(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportHost", reflect.TypeOf((*MockCmdbClient)(nil).ImportHost), varargs...)
}

// ImportLake mocks base method
func (m *MockCmdbClient) ImportLake(arg0 context.Context, arg1 *cmdb.ImportLakeRequest, arg2 ...grpc.CallOption) (*cmdb.ImportLakeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportLake", varargs...)
	ret0, _ := ret[0].(*cmdb.ImportLakeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportLake indicates an expected call of ImportLake
func (mr *MockCmdbClientMockRecorder) ImportLake(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportLake", reflect.TypeOf((*MockCmdbClient)(nil).ImportLake), varargs...)
}

// ImportSwitch mocks base method
func (m *MockCmdbClient) ImportSwitch(arg0 context.Context, arg1 *cmdb.ImportSwitchRequest, arg2 ...grpc.CallOption) (*cmdb.ImportSwitchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportSwitch", varargs...)
	ret0, _ := ret[0].(*cmdb.ImportSwitchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportSwitch indicates an expected call of ImportSwitch
func (mr *MockCmdbClientMockRecorder) ImportSwitch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportSwitch", reflect.TypeOf((*MockCmdbClient)(nil).ImportSwitch), varargs...)
}

// InstanceTopology mocks base method
func (m *MockCmdbClient) InstanceTopology(arg0 context.Context, arg1 *cmdb.InstanceTopologyRequest, arg2 ...grpc.CallOption) (*cmdb.InstanceTopologyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InstanceTopology", varargs...)
	ret0, _ := ret[0].(*cmdb.InstanceTopologyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceTopology indicates an expected call of InstanceTopology
func (mr *MockCmdbClientMockRecorder) InstanceTopology(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceTopology", reflect.TypeOf((*MockCmdbClient)(nil).InstanceTopology), varargs...)
}

// SearchHost mocks base method
func (m *MockCmdbClient) SearchHost(arg0 context.Context, arg1 *cmdb.SearchHostRequest, arg2 ...grpc.CallOption) (*cmdb.SearchHostResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchHost", varargs...)
	ret0, _ := ret[0].(*cmdb.SearchHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchHost indicates an expected call of SearchHost
func (mr *MockCmdbClientMockRecorder) SearchHost(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchHost", reflect.TypeOf((*MockCmdbClient)(nil).SearchHost), varargs...)
}

// SearchModule mocks base method
func (m *MockCmdbClient) SearchModule(arg0 context.Context, arg1 *cmdb.SearchMoudleRequest, arg2 ...grpc.CallOption) (*cmdb.SearchMoudleResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchModule", varargs...)
	ret0, _ := ret[0].(*cmdb.SearchMoudleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchModule indicates an expected call of SearchModule
func (mr *MockCmdbClientMockRecorder) SearchModule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchModule", reflect.TypeOf((*MockCmdbClient)(nil).SearchModule), varargs...)
}
