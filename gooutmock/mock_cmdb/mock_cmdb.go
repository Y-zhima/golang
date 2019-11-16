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

// ImportAsset mocks base method
func (m *MockCmdbClient) ImportAsset(arg0 context.Context, arg1 *cmdb.ImportAssetRequest, arg2 ...grpc.CallOption) (*cmdb.ImportAssetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportAsset", varargs...)
	ret0, _ := ret[0].(*cmdb.ImportAssetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportAsset indicates an expected call of ImportAsset
func (mr *MockCmdbClientMockRecorder) ImportAsset(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportAsset", reflect.TypeOf((*MockCmdbClient)(nil).ImportAsset), varargs...)
}

// ImportCrossTable mocks base method
func (m *MockCmdbClient) ImportCrossTable(arg0 context.Context, arg1 *cmdb.ImportCrossTableRequest, arg2 ...grpc.CallOption) (*cmdb.ImportCrossTableResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportCrossTable", varargs...)
	ret0, _ := ret[0].(*cmdb.ImportCrossTableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportCrossTable indicates an expected call of ImportCrossTable
func (mr *MockCmdbClientMockRecorder) ImportCrossTable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportCrossTable", reflect.TypeOf((*MockCmdbClient)(nil).ImportCrossTable), varargs...)
}

// ImportDetail mocks base method
func (m *MockCmdbClient) ImportDetail(arg0 context.Context, arg1 *cmdb.ImportDetailRequest, arg2 ...grpc.CallOption) (*cmdb.ImportDetailResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportDetail", varargs...)
	ret0, _ := ret[0].(*cmdb.ImportDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportDetail indicates an expected call of ImportDetail
func (mr *MockCmdbClientMockRecorder) ImportDetail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportDetail", reflect.TypeOf((*MockCmdbClient)(nil).ImportDetail), varargs...)
}

// ImportHistory mocks base method
func (m *MockCmdbClient) ImportHistory(arg0 context.Context, arg1 *cmdb.ImportHistoryRequest, arg2 ...grpc.CallOption) (*cmdb.ImportHistoryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportHistory", varargs...)
	ret0, _ := ret[0].(*cmdb.ImportHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportHistory indicates an expected call of ImportHistory
func (mr *MockCmdbClientMockRecorder) ImportHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportHistory", reflect.TypeOf((*MockCmdbClient)(nil).ImportHistory), varargs...)
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

// ImportReview mocks base method
func (m *MockCmdbClient) ImportReview(arg0 context.Context, arg1 *cmdb.ImportReviewRequest, arg2 ...grpc.CallOption) (*cmdb.ImportReviewResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportReview", varargs...)
	ret0, _ := ret[0].(*cmdb.ImportReviewResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportReview indicates an expected call of ImportReview
func (mr *MockCmdbClientMockRecorder) ImportReview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportReview", reflect.TypeOf((*MockCmdbClient)(nil).ImportReview), varargs...)
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

// RoomTopology mocks base method
func (m *MockCmdbClient) RoomTopology(arg0 context.Context, arg1 *cmdb.RoomTopologyRequest, arg2 ...grpc.CallOption) (*cmdb.RoomTopologyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RoomTopology", varargs...)
	ret0, _ := ret[0].(*cmdb.RoomTopologyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoomTopology indicates an expected call of RoomTopology
func (mr *MockCmdbClientMockRecorder) RoomTopology(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoomTopology", reflect.TypeOf((*MockCmdbClient)(nil).RoomTopology), varargs...)
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

// SearchInst mocks base method
func (m *MockCmdbClient) SearchInst(arg0 context.Context, arg1 *cmdb.SearchInstRequest, arg2 ...grpc.CallOption) (*cmdb.SearchInstResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchInst", varargs...)
	ret0, _ := ret[0].(*cmdb.SearchInstResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchInst indicates an expected call of SearchInst
func (mr *MockCmdbClientMockRecorder) SearchInst(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchInst", reflect.TypeOf((*MockCmdbClient)(nil).SearchInst), varargs...)
}

// SearchLake mocks base method
func (m *MockCmdbClient) SearchLake(arg0 context.Context, arg1 *cmdb.SearchLakeRequest, arg2 ...grpc.CallOption) (*cmdb.SearchLakeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchLake", varargs...)
	ret0, _ := ret[0].(*cmdb.SearchLakeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchLake indicates an expected call of SearchLake
func (mr *MockCmdbClientMockRecorder) SearchLake(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchLake", reflect.TypeOf((*MockCmdbClient)(nil).SearchLake), varargs...)
}

// SearchLakeArea mocks base method
func (m *MockCmdbClient) SearchLakeArea(arg0 context.Context, arg1 *cmdb.SearchLakeAreaRequest, arg2 ...grpc.CallOption) (*cmdb.SearchLakeAreaResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchLakeArea", varargs...)
	ret0, _ := ret[0].(*cmdb.SearchLakeAreaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchLakeArea indicates an expected call of SearchLakeArea
func (mr *MockCmdbClientMockRecorder) SearchLakeArea(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchLakeArea", reflect.TypeOf((*MockCmdbClient)(nil).SearchLakeArea), varargs...)
}

// SearchLakeHost mocks base method
func (m *MockCmdbClient) SearchLakeHost(arg0 context.Context, arg1 *cmdb.SearchLakeHostRequest, arg2 ...grpc.CallOption) (*cmdb.SearchLakeHostResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchLakeHost", varargs...)
	ret0, _ := ret[0].(*cmdb.SearchLakeHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchLakeHost indicates an expected call of SearchLakeHost
func (mr *MockCmdbClientMockRecorder) SearchLakeHost(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchLakeHost", reflect.TypeOf((*MockCmdbClient)(nil).SearchLakeHost), varargs...)
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

// ServerList mocks base method
func (m *MockCmdbClient) ServerList(arg0 context.Context, arg1 *cmdb.ServerListRequest, arg2 ...grpc.CallOption) (*cmdb.ServerListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ServerList", varargs...)
	ret0, _ := ret[0].(*cmdb.ServerListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerList indicates an expected call of ServerList
func (mr *MockCmdbClientMockRecorder) ServerList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerList", reflect.TypeOf((*MockCmdbClient)(nil).ServerList), varargs...)
}

// UpdateInst mocks base method
func (m *MockCmdbClient) UpdateInst(arg0 context.Context, arg1 *cmdb.UpdateInstRequest, arg2 ...grpc.CallOption) (*cmdb.UpdateInstResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateInst", varargs...)
	ret0, _ := ret[0].(*cmdb.UpdateInstResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInst indicates an expected call of UpdateInst
func (mr *MockCmdbClientMockRecorder) UpdateInst(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInst", reflect.TypeOf((*MockCmdbClient)(nil).UpdateInst), varargs...)
}
