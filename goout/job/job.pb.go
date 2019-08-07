// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job/job.proto

package job

import (
	context "context"
	fmt "fmt"
	common "git.fogcdn.top/axe/protos/goout/common"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 作业实例
type JobObject struct {
	JobId                int32    `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobObject) Reset()         { *m = JobObject{} }
func (m *JobObject) String() string { return proto.CompactTextString(m) }
func (*JobObject) ProtoMessage()    {}
func (*JobObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{0}
}

func (m *JobObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobObject.Unmarshal(m, b)
}
func (m *JobObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobObject.Marshal(b, m, deterministic)
}
func (m *JobObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobObject.Merge(m, src)
}
func (m *JobObject) XXX_Size() int {
	return xxx_messageInfo_JobObject.Size(m)
}
func (m *JobObject) XXX_DiscardUnknown() {
	xxx_messageInfo_JobObject.DiscardUnknown(m)
}

var xxx_messageInfo_JobObject proto.InternalMessageInfo

func (m *JobObject) GetJobId() int32 {
	if m != nil {
		return m.JobId
	}
	return 0
}

type CreateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

// 创建作业请求返回
type CreateResponse struct {
	JobId                int32                  `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Status               *common.ResponseStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetJobId() int32 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *CreateResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetRequest struct {
	JobId                int32    `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetJobId() int32 {
	if m != nil {
		return m.JobId
	}
	return 0
}

// 获取作业请求返回
type GetResponse struct {
	Job                  *JobObject             `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Status               *common.ResponseStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{4}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetJob() *JobObject {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *GetResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetLogRequest struct {
	JobId                int32    `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLogRequest) Reset()         { *m = GetLogRequest{} }
func (m *GetLogRequest) String() string { return proto.CompactTextString(m) }
func (*GetLogRequest) ProtoMessage()    {}
func (*GetLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{5}
}

func (m *GetLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLogRequest.Unmarshal(m, b)
}
func (m *GetLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLogRequest.Marshal(b, m, deterministic)
}
func (m *GetLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLogRequest.Merge(m, src)
}
func (m *GetLogRequest) XXX_Size() int {
	return xxx_messageInfo_GetLogRequest.Size(m)
}
func (m *GetLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLogRequest proto.InternalMessageInfo

func (m *GetLogRequest) GetJobId() int32 {
	if m != nil {
		return m.JobId
	}
	return 0
}

// 获取作业请求返回
type GetLogResponse struct {
	Status               *common.ResponseStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetLogResponse) Reset()         { *m = GetLogResponse{} }
func (m *GetLogResponse) String() string { return proto.CompactTextString(m) }
func (*GetLogResponse) ProtoMessage()    {}
func (*GetLogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{6}
}

func (m *GetLogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLogResponse.Unmarshal(m, b)
}
func (m *GetLogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLogResponse.Marshal(b, m, deterministic)
}
func (m *GetLogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLogResponse.Merge(m, src)
}
func (m *GetLogResponse) XXX_Size() int {
	return xxx_messageInfo_GetLogResponse.Size(m)
}
func (m *GetLogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLogResponse proto.InternalMessageInfo

func (m *GetLogResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// 筛选作业请求
type FilterRequest struct {
	Paging               *common.Paging `protobuf:"bytes,1,opt,name=paging,proto3" json:"paging,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FilterRequest) Reset()         { *m = FilterRequest{} }
func (m *FilterRequest) String() string { return proto.CompactTextString(m) }
func (*FilterRequest) ProtoMessage()    {}
func (*FilterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{7}
}

func (m *FilterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterRequest.Unmarshal(m, b)
}
func (m *FilterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterRequest.Marshal(b, m, deterministic)
}
func (m *FilterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterRequest.Merge(m, src)
}
func (m *FilterRequest) XXX_Size() int {
	return xxx_messageInfo_FilterRequest.Size(m)
}
func (m *FilterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilterRequest proto.InternalMessageInfo

func (m *FilterRequest) GetPaging() *common.Paging {
	if m != nil {
		return m.Paging
	}
	return nil
}

// 筛选作业请求返回
type FilterResponse struct {
	Jobs                 []*JobObject           `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	Paging               *common.Paging         `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
	Status               *common.ResponseStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FilterResponse) Reset()         { *m = FilterResponse{} }
func (m *FilterResponse) String() string { return proto.CompactTextString(m) }
func (*FilterResponse) ProtoMessage()    {}
func (*FilterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{8}
}

func (m *FilterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterResponse.Unmarshal(m, b)
}
func (m *FilterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterResponse.Marshal(b, m, deterministic)
}
func (m *FilterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterResponse.Merge(m, src)
}
func (m *FilterResponse) XXX_Size() int {
	return xxx_messageInfo_FilterResponse.Size(m)
}
func (m *FilterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FilterResponse proto.InternalMessageInfo

func (m *FilterResponse) GetJobs() []*JobObject {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *FilterResponse) GetPaging() *common.Paging {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *FilterResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ScheduleRequest struct {
	JobId                int32    `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduleRequest) Reset()         { *m = ScheduleRequest{} }
func (m *ScheduleRequest) String() string { return proto.CompactTextString(m) }
func (*ScheduleRequest) ProtoMessage()    {}
func (*ScheduleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{9}
}

func (m *ScheduleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleRequest.Unmarshal(m, b)
}
func (m *ScheduleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleRequest.Marshal(b, m, deterministic)
}
func (m *ScheduleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleRequest.Merge(m, src)
}
func (m *ScheduleRequest) XXX_Size() int {
	return xxx_messageInfo_ScheduleRequest.Size(m)
}
func (m *ScheduleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleRequest proto.InternalMessageInfo

func (m *ScheduleRequest) GetJobId() int32 {
	if m != nil {
		return m.JobId
	}
	return 0
}

type ScheduleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduleResponse) Reset()         { *m = ScheduleResponse{} }
func (m *ScheduleResponse) String() string { return proto.CompactTextString(m) }
func (*ScheduleResponse) ProtoMessage()    {}
func (*ScheduleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{10}
}

func (m *ScheduleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleResponse.Unmarshal(m, b)
}
func (m *ScheduleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleResponse.Marshal(b, m, deterministic)
}
func (m *ScheduleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleResponse.Merge(m, src)
}
func (m *ScheduleResponse) XXX_Size() int {
	return xxx_messageInfo_ScheduleResponse.Size(m)
}
func (m *ScheduleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleResponse proto.InternalMessageInfo

type ScheduleListRequest struct {
	JobId                int32    `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduleListRequest) Reset()         { *m = ScheduleListRequest{} }
func (m *ScheduleListRequest) String() string { return proto.CompactTextString(m) }
func (*ScheduleListRequest) ProtoMessage()    {}
func (*ScheduleListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{11}
}

func (m *ScheduleListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleListRequest.Unmarshal(m, b)
}
func (m *ScheduleListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleListRequest.Marshal(b, m, deterministic)
}
func (m *ScheduleListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleListRequest.Merge(m, src)
}
func (m *ScheduleListRequest) XXX_Size() int {
	return xxx_messageInfo_ScheduleListRequest.Size(m)
}
func (m *ScheduleListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleListRequest proto.InternalMessageInfo

func (m *ScheduleListRequest) GetJobId() int32 {
	if m != nil {
		return m.JobId
	}
	return 0
}

type ScheduleListResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduleListResponse) Reset()         { *m = ScheduleListResponse{} }
func (m *ScheduleListResponse) String() string { return proto.CompactTextString(m) }
func (*ScheduleListResponse) ProtoMessage()    {}
func (*ScheduleListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{12}
}

func (m *ScheduleListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleListResponse.Unmarshal(m, b)
}
func (m *ScheduleListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleListResponse.Marshal(b, m, deterministic)
}
func (m *ScheduleListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleListResponse.Merge(m, src)
}
func (m *ScheduleListResponse) XXX_Size() int {
	return xxx_messageInfo_ScheduleListResponse.Size(m)
}
func (m *ScheduleListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleListResponse proto.InternalMessageInfo

type DeleteRequest struct {
	JobId                int32    `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{13}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetJobId() int32 {
	if m != nil {
		return m.JobId
	}
	return 0
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{14}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*JobObject)(nil), "job.JobObject")
	proto.RegisterType((*CreateRequest)(nil), "job.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "job.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "job.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "job.GetResponse")
	proto.RegisterType((*GetLogRequest)(nil), "job.GetLogRequest")
	proto.RegisterType((*GetLogResponse)(nil), "job.GetLogResponse")
	proto.RegisterType((*FilterRequest)(nil), "job.FilterRequest")
	proto.RegisterType((*FilterResponse)(nil), "job.FilterResponse")
	proto.RegisterType((*ScheduleRequest)(nil), "job.ScheduleRequest")
	proto.RegisterType((*ScheduleResponse)(nil), "job.ScheduleResponse")
	proto.RegisterType((*ScheduleListRequest)(nil), "job.ScheduleListRequest")
	proto.RegisterType((*ScheduleListResponse)(nil), "job.ScheduleListResponse")
	proto.RegisterType((*DeleteRequest)(nil), "job.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "job.DeleteResponse")
}

func init() { proto.RegisterFile("job/job.proto", fileDescriptor_c3f97a06ee6e4937) }

var fileDescriptor_c3f97a06ee6e4937 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x4e, 0xd4, 0x5c,
	0x14, 0x4d, 0x19, 0xe8, 0xf7, 0xb1, 0x61, 0x7e, 0x72, 0xf8, 0x71, 0x28, 0x6a, 0x6a, 0x31, 0x84,
	0x10, 0x68, 0x61, 0xbc, 0x30, 0xe1, 0x8a, 0x41, 0x91, 0x80, 0x24, 0x4e, 0x06, 0x13, 0x8d, 0x89,
	0x21, 0xed, 0xcc, 0xa1, 0xb6, 0x29, 0xdd, 0x75, 0xce, 0x19, 0x35, 0x31, 0xde, 0x78, 0x6d, 0xe2,
	0xdf, 0x9d, 0xf1, 0x21, 0x4c, 0xf4, 0xc2, 0x9f, 0x18, 0x1f, 0xc2, 0x57, 0x80, 0xe8, 0x63, 0x98,
	0x9e, 0x9e, 0xc2, 0x94, 0x61, 0x1c, 0xb9, 0x9a, 0x64, 0xed, 0xb5, 0xd7, 0x5e, 0x7b, 0xcd, 0x3e,
	0x85, 0xbc, 0x8f, 0x8e, 0xe5, 0xa3, 0x63, 0x46, 0x2d, 0xe4, 0x48, 0x72, 0x3e, 0x3a, 0x5a, 0xa9,
	0x81, 0xfb, 0xfb, 0x18, 0x5a, 0x76, 0xe4, 0x25, 0xb0, 0x76, 0xde, 0x45, 0x74, 0x03, 0x1a, 0x23,
	0x96, 0x1d, 0x86, 0xc8, 0x6d, 0xee, 0x61, 0xc8, 0x64, 0x75, 0x41, 0xfc, 0x34, 0x16, 0x5d, 0x1a,
	0x2e, 0xb2, 0xc7, 0xb6, 0xeb, 0xd2, 0x96, 0x85, 0x91, 0x60, 0x74, 0xb3, 0x0d, 0x03, 0x86, 0xb7,
	0xd0, 0xb9, 0xe5, 0xf8, 0xb4, 0xc1, 0xc9, 0x04, 0xa8, 0x3e, 0x3a, 0xbb, 0x5e, 0xb3, 0xac, 0xe8,
	0xca, 0xdc, 0x50, 0x7d, 0xc8, 0x47, 0x67, 0xb3, 0x69, 0x14, 0x21, 0x7f, 0xad, 0x45, 0x6d, 0x4e,
	0xeb, 0xf4, 0x61, 0x9b, 0x32, 0x6e, 0xdc, 0x81, 0x42, 0x0a, 0xb0, 0x08, 0x43, 0x46, 0x7b, 0x74,
	0x12, 0x13, 0x54, 0xc6, 0x6d, 0xde, 0x66, 0xe5, 0x01, 0x5d, 0x99, 0x1b, 0xa9, 0x4c, 0x9a, 0xc9,
	0x32, 0x66, 0xda, 0xb8, 0x23, 0xaa, 0x75, 0xc9, 0x32, 0x66, 0x00, 0x36, 0x28, 0x97, 0x63, 0x7a,
	0xd9, 0xd9, 0x85, 0x11, 0x41, 0x92, 0xa3, 0x75, 0x88, 0x63, 0x12, 0x94, 0x91, 0x4a, 0xc1, 0x8c,
	0xd3, 0x3b, 0xda, 0xa8, 0x1e, 0x97, 0xce, 0xec, 0x62, 0x16, 0xf2, 0x1b, 0x94, 0x6f, 0xa3, 0xdb,
	0xc7, 0xc8, 0x2a, 0x14, 0x52, 0x9e, 0xf4, 0x72, 0xd6, 0x49, 0x57, 0x21, 0x7f, 0xc3, 0x0b, 0x38,
	0x6d, 0xa5, 0x93, 0x66, 0x41, 0x8d, 0x6c, 0xd7, 0x0b, 0xdd, 0xa3, 0x7d, 0xa4, 0x40, 0x4d, 0xa0,
	0x75, 0x59, 0x35, 0x5e, 0x28, 0x50, 0x48, 0x3b, 0xe5, 0x6c, 0x03, 0x06, 0x7d, 0x74, 0x58, 0x59,
	0xd1, 0x73, 0xa7, 0x04, 0x21, 0x6a, 0x1d, 0xf2, 0x03, 0x7f, 0x93, 0xef, 0xd8, 0x23, 0xf7, 0x4f,
	0x7b, 0xcc, 0x41, 0x71, 0xa7, 0xf1, 0x80, 0x36, 0xdb, 0x01, 0xed, 0x93, 0x19, 0x81, 0xd2, 0x31,
	0x33, 0xd1, 0x32, 0x16, 0x60, 0x2c, 0xc5, 0xb6, 0x3d, 0xd6, 0xef, 0xef, 0x9f, 0x84, 0xf1, 0x2c,
	0x5b, 0xaa, 0xcc, 0x42, 0xfe, 0x3a, 0x0d, 0x28, 0xef, 0xe7, 0xa0, 0x04, 0x85, 0x94, 0x97, 0x74,
	0x56, 0xbe, 0x0e, 0x42, 0x6e, 0x0b, 0x1d, 0xb2, 0x0e, 0x6a, 0x72, 0xd6, 0x84, 0x88, 0xf4, 0x32,
	0x47, 0xaf, 0x8d, 0x65, 0x30, 0x39, 0x94, 0x3c, 0xff, 0x79, 0xf0, 0x76, 0x60, 0xd4, 0xf8, 0xcf,
	0x7a, 0xb4, 0x1c, 0x3f, 0xdc, 0x15, 0x65, 0x9e, 0xdc, 0x87, 0xff, 0x53, 0x83, 0x64, 0x5c, 0x34,
	0x9d, 0xc8, 0x46, 0x9b, 0x38, 0x81, 0x4a, 0xb1, 0xcb, 0x42, 0xec, 0xa2, 0x31, 0x25, 0xc5, 0xac,
	0xa7, 0x89, 0xff, 0x67, 0x16, 0x93, 0xd4, 0x58, 0x7e, 0x0f, 0x46, 0x3b, 0xf7, 0x27, 0xe5, 0x8c,
	0x58, 0x47, 0x80, 0xda, 0xd4, 0x29, 0x15, 0x39, 0xea, 0x92, 0x18, 0x35, 0x4d, 0x7a, 0x8f, 0x22,
	0x6b, 0xa0, 0x26, 0x17, 0x26, 0xd3, 0xc8, 0x1c, 0xaa, 0x4c, 0x23, 0x7b, 0x82, 0x46, 0x51, 0xa8,
	0x0e, 0x93, 0x34, 0x0d, 0xb2, 0x06, 0xb9, 0x0d, 0xca, 0x49, 0x51, 0x90, 0x8f, 0x5f, 0xb6, 0x56,
	0x3a, 0x06, 0x64, 0x6b, 0x59, 0xb4, 0x12, 0x52, 0x3a, 0x69, 0x88, 0xd4, 0x40, 0x4d, 0x5e, 0x99,
	0xf4, 0x91, 0x79, 0x9a, 0xd2, 0x47, 0xf6, 0x19, 0x1a, 0x17, 0x84, 0xd8, 0x39, 0x32, 0xd1, 0xb5,
	0x5d, 0x80, 0x2e, 0x23, 0x37, 0x41, 0x4d, 0x2e, 0x40, 0x2a, 0x66, 0xce, 0x46, 0x2a, 0x66, 0x4f,
	0x24, 0xb5, 0x37, 0xdf, 0x65, 0x6f, 0xed, 0xb5, 0xf2, 0xa6, 0xea, 0x91, 0x55, 0x98, 0xae, 0xde,
	0x5d, 0xff, 0xfd, 0xe3, 0xd5, 0xaf, 0x6f, 0x9f, 0xaa, 0xb5, 0x4d, 0x7d, 0x51, 0x3f, 0xf8, 0xfe,
	0xfe, 0xe0, 0xf3, 0xbb, 0xc3, 0x2f, 0x1f, 0x0f, 0x5f, 0x7e, 0x20, 0xda, 0x51, 0x31, 0x01, 0xea,
	0xeb, 0x3b, 0xb7, 0xf7, 0xda, 0x81, 0x5e, 0xad, 0x6d, 0x56, 0x86, 0x96, 0xcd, 0x25, 0x73, 0x69,
	0x5e, 0x51, 0x2a, 0x25, 0x3b, 0x8a, 0x02, 0xaf, 0x21, 0xbe, 0xcf, 0x96, 0xcf, 0x30, 0x5c, 0xe9,
	0x42, 0xee, 0xcd, 0xb8, 0x1e, 0x37, 0xf7, 0xd0, 0x6d, 0x34, 0x43, 0x93, 0x63, 0x64, 0xd9, 0x4f,
	0xa8, 0x25, 0xbe, 0xe4, 0xcc, 0x72, 0x11, 0xdb, 0x3c, 0x36, 0xe7, 0xa8, 0x02, 0xb9, 0xf2, 0x27,
	0x00, 0x00, 0xff, 0xff, 0xd8, 0x67, 0xa7, 0xeb, 0x4f, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobClient interface {
	// 创建作业任务(执行作业模板)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// 创建定时作业任务
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
	// 查看定时作业任务
	ScheduleList(ctx context.Context, in *ScheduleListRequest, opts ...grpc.CallOption) (*ScheduleListResponse, error)
	// 筛选作业任务
	Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error)
	// 获取作业任务
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// 获取作业任务详细执行过程
	GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogResponse, error)
	// 删除作业任务
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type jobClient struct {
	cc *grpc.ClientConn
}

func NewJobClient(cc *grpc.ClientConn) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/job.Job/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, "/job.Job/Schedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) ScheduleList(ctx context.Context, in *ScheduleListRequest, opts ...grpc.CallOption) (*ScheduleListResponse, error) {
	out := new(ScheduleListResponse)
	err := c.cc.Invoke(ctx, "/job.Job/ScheduleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error) {
	out := new(FilterResponse)
	err := c.cc.Invoke(ctx, "/job.Job/Filter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/job.Job/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogResponse, error) {
	out := new(GetLogResponse)
	err := c.cc.Invoke(ctx, "/job.Job/GetLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/job.Job/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
type JobServer interface {
	// 创建作业任务(执行作业模板)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// 创建定时作业任务
	Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
	// 查看定时作业任务
	ScheduleList(context.Context, *ScheduleListRequest) (*ScheduleListResponse, error)
	// 筛选作业任务
	Filter(context.Context, *FilterRequest) (*FilterResponse, error)
	// 获取作业任务
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// 获取作业任务详细执行过程
	GetLog(context.Context, *GetLogRequest) (*GetLogResponse, error)
	// 删除作业任务
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedJobServer can be embedded to have forward compatible implementations.
type UnimplementedJobServer struct {
}

func (*UnimplementedJobServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedJobServer) Schedule(ctx context.Context, req *ScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Schedule not implemented")
}
func (*UnimplementedJobServer) ScheduleList(ctx context.Context, req *ScheduleListRequest) (*ScheduleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleList not implemented")
}
func (*UnimplementedJobServer) Filter(ctx context.Context, req *FilterRequest) (*FilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Filter not implemented")
}
func (*UnimplementedJobServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedJobServer) GetLog(ctx context.Context, req *GetLogRequest) (*GetLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLog not implemented")
}
func (*UnimplementedJobServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterJobServer(s *grpc.Server, srv JobServer) {
	s.RegisterService(&_Job_serviceDesc, srv)
}

func _Job_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/Schedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Schedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_ScheduleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).ScheduleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/ScheduleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).ScheduleList(ctx, req.(*ScheduleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_Filter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Filter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Filter(ctx, req.(*FilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/GetLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetLog(ctx, req.(*GetLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Job_serviceDesc = grpc.ServiceDesc{
	ServiceName: "job.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Job_Create_Handler,
		},
		{
			MethodName: "Schedule",
			Handler:    _Job_Schedule_Handler,
		},
		{
			MethodName: "ScheduleList",
			Handler:    _Job_ScheduleList_Handler,
		},
		{
			MethodName: "Filter",
			Handler:    _Job_Filter_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Job_Get_Handler,
		},
		{
			MethodName: "GetLog",
			Handler:    _Job_GetLog_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Job_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job/job.proto",
}
