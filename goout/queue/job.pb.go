// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queue/job.proto

package queue

import (
	fmt "fmt"
	cmdb "git.fogcdn.top/axe/protos/goout/cmdb"
	schedule "git.fogcdn.top/axe/protos/goout/schedule"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 下发任务对象
type DeliverJobObject struct {
	// 任务类型
	TaskType schedule.TaskType `protobuf:"varint,1,opt,name=task_type,json=taskType,proto3,enum=schedule.TaskType" json:"task_type,omitempty"`
	// cmdb的搜索条件
	CmdbSearchRequest *cmdb.ChooseHostRequest `protobuf:"bytes,2,opt,name=cmdb_search_request,json=cmdbSearchRequest,proto3" json:"cmdb_search_request,omitempty"`
	// playbook 文件的URL
	PlaybookFileUrl string `protobuf:"bytes,3,opt,name=playbook_file_url,json=playbookFileUrl,proto3" json:"playbook_file_url,omitempty"`
	// playbook 文件的MD5
	PlaybookFileMd5 string `protobuf:"bytes,4,opt,name=playbook_file_md5,json=playbookFileMd5,proto3" json:"playbook_file_md5,omitempty"`
	// playbook 入口文件
	PlaybookYmlName string `protobuf:"bytes,5,opt,name=playbook_yml_name,json=playbookYmlName,proto3" json:"playbook_yml_name,omitempty"`
	// 额外变量
	ExtraVar             string   `protobuf:"bytes,6,opt,name=extra_var,json=extraVar,proto3" json:"extra_var,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliverJobObject) Reset()         { *m = DeliverJobObject{} }
func (m *DeliverJobObject) String() string { return proto.CompactTextString(m) }
func (*DeliverJobObject) ProtoMessage()    {}
func (*DeliverJobObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aceba9723f0c267, []int{0}
}

func (m *DeliverJobObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliverJobObject.Unmarshal(m, b)
}
func (m *DeliverJobObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliverJobObject.Marshal(b, m, deterministic)
}
func (m *DeliverJobObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliverJobObject.Merge(m, src)
}
func (m *DeliverJobObject) XXX_Size() int {
	return xxx_messageInfo_DeliverJobObject.Size(m)
}
func (m *DeliverJobObject) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliverJobObject.DiscardUnknown(m)
}

var xxx_messageInfo_DeliverJobObject proto.InternalMessageInfo

func (m *DeliverJobObject) GetTaskType() schedule.TaskType {
	if m != nil {
		return m.TaskType
	}
	return schedule.TaskType_UNDEFINED
}

func (m *DeliverJobObject) GetCmdbSearchRequest() *cmdb.ChooseHostRequest {
	if m != nil {
		return m.CmdbSearchRequest
	}
	return nil
}

func (m *DeliverJobObject) GetPlaybookFileUrl() string {
	if m != nil {
		return m.PlaybookFileUrl
	}
	return ""
}

func (m *DeliverJobObject) GetPlaybookFileMd5() string {
	if m != nil {
		return m.PlaybookFileMd5
	}
	return ""
}

func (m *DeliverJobObject) GetPlaybookYmlName() string {
	if m != nil {
		return m.PlaybookYmlName
	}
	return ""
}

func (m *DeliverJobObject) GetExtraVar() string {
	if m != nil {
		return m.ExtraVar
	}
	return ""
}

func init() {
	proto.RegisterType((*DeliverJobObject)(nil), "queue.DeliverJobObject")
}

func init() { proto.RegisterFile("queue/job.proto", fileDescriptor_5aceba9723f0c267) }

var fileDescriptor_5aceba9723f0c267 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x51, 0x6b, 0xea, 0x30,
	0x18, 0x86, 0xa9, 0xe7, 0x28, 0xda, 0x03, 0xc7, 0x63, 0xcf, 0x85, 0x45, 0x77, 0x21, 0x83, 0x81,
	0x8c, 0xd9, 0x80, 0xc3, 0x3f, 0xb0, 0xc9, 0x36, 0x06, 0xdb, 0xa0, 0x73, 0x83, 0xed, 0xa6, 0xa4,
	0xed, 0x67, 0xac, 0xa6, 0xfd, 0x62, 0x92, 0x3a, 0xfb, 0xa3, 0xf6, 0x1f, 0x87, 0x69, 0x1d, 0x88,
	0xbb, 0x49, 0x3e, 0x9e, 0xf7, 0x79, 0x49, 0x48, 0xec, 0xf6, 0x3a, 0x87, 0x1c, 0xc8, 0x12, 0x43,
	0x4f, 0x48, 0xd4, 0xe8, 0xd4, 0x0d, 0xe8, 0x9d, 0x30, 0x44, 0xc6, 0x81, 0x50, 0x91, 0x10, 0x9a,
	0x65, 0xa8, 0xa9, 0x4e, 0x30, 0x53, 0xa5, 0xd4, 0xbb, 0x30, 0x5b, 0x34, 0x62, 0x90, 0x8d, 0xd4,
	0x07, 0x65, 0x0c, 0x24, 0x41, 0x61, 0x8c, 0x1f, 0xec, 0x76, 0x94, 0xc6, 0x21, 0xd9, 0x2d, 0x15,
	0xe8, 0xaa, 0x68, 0x01, 0x71, 0xce, 0x81, 0xec, 0x87, 0x32, 0x38, 0xfd, 0xac, 0xd9, 0xff, 0xa6,
	0xc0, 0x93, 0x0d, 0xc8, 0x7b, 0x0c, 0x9f, 0xc2, 0x25, 0x44, 0xda, 0x21, 0x76, 0x4b, 0x53, 0xb5,
	0x0a, 0x74, 0x21, 0xc0, 0xb5, 0x06, 0xd6, 0xf0, 0xef, 0xd8, 0xf1, 0xbe, 0x8b, 0x33, 0xaa, 0x56,
	0xb3, 0x42, 0x80, 0xdf, 0xd4, 0xd5, 0xe4, 0xdc, 0xda, 0xff, 0x77, 0x87, 0x05, 0x0a, 0xa8, 0x8c,
	0x16, 0x81, 0x84, 0x75, 0x0e, 0x4a, 0xbb, 0xb5, 0x81, 0x35, 0xfc, 0x33, 0xee, 0x7a, 0xe6, 0x22,
	0xd7, 0x0b, 0x44, 0x05, 0x77, 0xa8, 0xb4, 0x5f, 0xc6, 0x7e, 0x67, 0xc7, 0x9f, 0x4d, 0xa5, 0x42,
	0xce, 0xb9, 0xdd, 0x11, 0x9c, 0x16, 0x21, 0xe2, 0x2a, 0x98, 0x27, 0x1c, 0x82, 0x5c, 0x72, 0xf7,
	0xd7, 0xc0, 0x1a, 0xb6, 0xfc, 0xf6, 0x3e, 0xb8, 0x49, 0x38, 0xbc, 0x48, 0x7e, 0xec, 0xa6, 0xf1,
	0xc4, 0xfd, 0x7d, 0xec, 0x3e, 0xc4, 0x93, 0x03, 0xb7, 0x48, 0x79, 0x90, 0xd1, 0x14, 0xdc, 0xfa,
	0xa1, 0xfb, 0x96, 0xf2, 0x47, 0x9a, 0x82, 0xd3, 0xb7, 0x5b, 0xb0, 0xd5, 0x92, 0x06, 0x1b, 0x2a,
	0xdd, 0x86, 0x71, 0x9a, 0x06, 0xbc, 0x52, 0x79, 0x35, 0xb5, 0xfb, 0x1a, 0x85, 0x37, 0x47, 0x16,
	0xc5, 0x99, 0x47, 0xb7, 0xd5, 0x3b, 0x2a, 0xcf, 0x7c, 0xe2, 0xfb, 0x19, 0x4b, 0xf4, 0x3e, 0xd4,
	0x28, 0x08, 0xdd, 0x02, 0x29, 0x05, 0xc2, 0x10, 0x73, 0x4d, 0x8c, 0x16, 0x36, 0x0c, 0xbb, 0xfc,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x46, 0xc6, 0x23, 0x0c, 0x02, 0x00, 0x00,
}