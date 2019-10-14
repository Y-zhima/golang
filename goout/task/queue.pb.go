// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task/queue.proto

package task

import (
	fmt "fmt"
	cmdb "git.fogcdn.top/axe/protos/goout/cmdb"
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

// 作业模板执行任务
type TemplateExecuteTask struct {
	// 子任务实例ID
	SubTaskId int64 `protobuf:"varint,1,opt,name=sub_task_id,json=subTaskId,proto3" json:"sub_task_id,omitempty"`
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

func (m *TemplateExecuteTask) Reset()         { *m = TemplateExecuteTask{} }
func (m *TemplateExecuteTask) String() string { return proto.CompactTextString(m) }
func (*TemplateExecuteTask) ProtoMessage()    {}
func (*TemplateExecuteTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_840e7c8beea44e4e, []int{0}
}

func (m *TemplateExecuteTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateExecuteTask.Unmarshal(m, b)
}
func (m *TemplateExecuteTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateExecuteTask.Marshal(b, m, deterministic)
}
func (m *TemplateExecuteTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateExecuteTask.Merge(m, src)
}
func (m *TemplateExecuteTask) XXX_Size() int {
	return xxx_messageInfo_TemplateExecuteTask.Size(m)
}
func (m *TemplateExecuteTask) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateExecuteTask.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateExecuteTask proto.InternalMessageInfo

func (m *TemplateExecuteTask) GetSubTaskId() int64 {
	if m != nil {
		return m.SubTaskId
	}
	return 0
}

func (m *TemplateExecuteTask) GetCmdbSearchRequest() *cmdb.ChooseHostRequest {
	if m != nil {
		return m.CmdbSearchRequest
	}
	return nil
}

func (m *TemplateExecuteTask) GetPlaybookFileUrl() string {
	if m != nil {
		return m.PlaybookFileUrl
	}
	return ""
}

func (m *TemplateExecuteTask) GetPlaybookFileMd5() string {
	if m != nil {
		return m.PlaybookFileMd5
	}
	return ""
}

func (m *TemplateExecuteTask) GetPlaybookYmlName() string {
	if m != nil {
		return m.PlaybookYmlName
	}
	return ""
}

func (m *TemplateExecuteTask) GetExtraVar() string {
	if m != nil {
		return m.ExtraVar
	}
	return ""
}

// 交维表导入后巡检任务
type ServerCompareTask struct {
	// 任务实例ID
	SubTaskId int64 `protobuf:"varint,1,opt,name=sub_task_id,json=subTaskId,proto3" json:"sub_task_id,omitempty"`
	// 查询条件
	CmdbSearchRequest    *cmdb.ChooseServerCompareRequest `protobuf:"bytes,2,opt,name=cmdb_search_request,json=cmdbSearchRequest,proto3" json:"cmdb_search_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ServerCompareTask) Reset()         { *m = ServerCompareTask{} }
func (m *ServerCompareTask) String() string { return proto.CompactTextString(m) }
func (*ServerCompareTask) ProtoMessage()    {}
func (*ServerCompareTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_840e7c8beea44e4e, []int{1}
}

func (m *ServerCompareTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerCompareTask.Unmarshal(m, b)
}
func (m *ServerCompareTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerCompareTask.Marshal(b, m, deterministic)
}
func (m *ServerCompareTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerCompareTask.Merge(m, src)
}
func (m *ServerCompareTask) XXX_Size() int {
	return xxx_messageInfo_ServerCompareTask.Size(m)
}
func (m *ServerCompareTask) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerCompareTask.DiscardUnknown(m)
}

var xxx_messageInfo_ServerCompareTask proto.InternalMessageInfo

func (m *ServerCompareTask) GetSubTaskId() int64 {
	if m != nil {
		return m.SubTaskId
	}
	return 0
}

func (m *ServerCompareTask) GetCmdbSearchRequest() *cmdb.ChooseServerCompareRequest {
	if m != nil {
		return m.CmdbSearchRequest
	}
	return nil
}

// 裸金属状态查询任务
type BareMetalSearchTask struct {
	// 子任务实例ID
	SubTaskId int64 `protobuf:"varint,1,opt,name=sub_task_id,json=subTaskId,proto3" json:"sub_task_id,omitempty"`
	// cmdb的搜索条件
	CmdbSearchRequest    *cmdb.ChooseServerRequest `protobuf:"bytes,2,opt,name=cmdb_search_request,json=cmdbSearchRequest,proto3" json:"cmdb_search_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BareMetalSearchTask) Reset()         { *m = BareMetalSearchTask{} }
func (m *BareMetalSearchTask) String() string { return proto.CompactTextString(m) }
func (*BareMetalSearchTask) ProtoMessage()    {}
func (*BareMetalSearchTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_840e7c8beea44e4e, []int{2}
}

func (m *BareMetalSearchTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalSearchTask.Unmarshal(m, b)
}
func (m *BareMetalSearchTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalSearchTask.Marshal(b, m, deterministic)
}
func (m *BareMetalSearchTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalSearchTask.Merge(m, src)
}
func (m *BareMetalSearchTask) XXX_Size() int {
	return xxx_messageInfo_BareMetalSearchTask.Size(m)
}
func (m *BareMetalSearchTask) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalSearchTask.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalSearchTask proto.InternalMessageInfo

func (m *BareMetalSearchTask) GetSubTaskId() int64 {
	if m != nil {
		return m.SubTaskId
	}
	return 0
}

func (m *BareMetalSearchTask) GetCmdbSearchRequest() *cmdb.ChooseServerRequest {
	if m != nil {
		return m.CmdbSearchRequest
	}
	return nil
}

// 裸金属电源开关任务
type BareMetalPowerTask struct {
	// 子任务实例ID
	SubTaskId int64 `protobuf:"varint,1,opt,name=sub_task_id,json=subTaskId,proto3" json:"sub_task_id,omitempty"`
	// 电源开关操作
	State cmdb.ServerPowerState `protobuf:"varint,2,opt,name=state,proto3,enum=cmdb.ServerPowerState" json:"state,omitempty"`
	// cmdb的搜索条件
	CmdbSearchRequest    *cmdb.ChooseServerRequest `protobuf:"bytes,3,opt,name=cmdb_search_request,json=cmdbSearchRequest,proto3" json:"cmdb_search_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BareMetalPowerTask) Reset()         { *m = BareMetalPowerTask{} }
func (m *BareMetalPowerTask) String() string { return proto.CompactTextString(m) }
func (*BareMetalPowerTask) ProtoMessage()    {}
func (*BareMetalPowerTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_840e7c8beea44e4e, []int{3}
}

func (m *BareMetalPowerTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalPowerTask.Unmarshal(m, b)
}
func (m *BareMetalPowerTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalPowerTask.Marshal(b, m, deterministic)
}
func (m *BareMetalPowerTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalPowerTask.Merge(m, src)
}
func (m *BareMetalPowerTask) XXX_Size() int {
	return xxx_messageInfo_BareMetalPowerTask.Size(m)
}
func (m *BareMetalPowerTask) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalPowerTask.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalPowerTask proto.InternalMessageInfo

func (m *BareMetalPowerTask) GetSubTaskId() int64 {
	if m != nil {
		return m.SubTaskId
	}
	return 0
}

func (m *BareMetalPowerTask) GetState() cmdb.ServerPowerState {
	if m != nil {
		return m.State
	}
	return cmdb.ServerPowerState_UNKNOWN
}

func (m *BareMetalPowerTask) GetCmdbSearchRequest() *cmdb.ChooseServerRequest {
	if m != nil {
		return m.CmdbSearchRequest
	}
	return nil
}

// 裸金属创建任务
type BareMetalCreateTask struct {
	// 子任务实例ID
	SubTaskId int64 `protobuf:"varint,1,opt,name=sub_task_id,json=subTaskId,proto3" json:"sub_task_id,omitempty"`
	// cmdb的搜索条件
	CmdbSearchRequest    *cmdb.ChooseServerRequest `protobuf:"bytes,2,opt,name=cmdb_search_request,json=cmdbSearchRequest,proto3" json:"cmdb_search_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BareMetalCreateTask) Reset()         { *m = BareMetalCreateTask{} }
func (m *BareMetalCreateTask) String() string { return proto.CompactTextString(m) }
func (*BareMetalCreateTask) ProtoMessage()    {}
func (*BareMetalCreateTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_840e7c8beea44e4e, []int{4}
}

func (m *BareMetalCreateTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalCreateTask.Unmarshal(m, b)
}
func (m *BareMetalCreateTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalCreateTask.Marshal(b, m, deterministic)
}
func (m *BareMetalCreateTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalCreateTask.Merge(m, src)
}
func (m *BareMetalCreateTask) XXX_Size() int {
	return xxx_messageInfo_BareMetalCreateTask.Size(m)
}
func (m *BareMetalCreateTask) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalCreateTask.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalCreateTask proto.InternalMessageInfo

func (m *BareMetalCreateTask) GetSubTaskId() int64 {
	if m != nil {
		return m.SubTaskId
	}
	return 0
}

func (m *BareMetalCreateTask) GetCmdbSearchRequest() *cmdb.ChooseServerRequest {
	if m != nil {
		return m.CmdbSearchRequest
	}
	return nil
}

// 裸金属安装任务
type BareMetalInstallTask struct {
	// 子任务实例ID
	SubTaskId int64 `protobuf:"varint,1,opt,name=sub_task_id,json=subTaskId,proto3" json:"sub_task_id,omitempty"`
	// cmdb的搜索条件
	CmdbSearchRequest *cmdb.ChooseServerRequest `protobuf:"bytes,2,opt,name=cmdb_search_request,json=cmdbSearchRequest,proto3" json:"cmdb_search_request,omitempty"`
	// 安装镜像文件的URL
	ImageFileUrl string `protobuf:"bytes,3,opt,name=image_file_url,json=imageFileUrl,proto3" json:"image_file_url,omitempty"`
	// 安装镜像文件的MD5
	ImageFileMd5 string `protobuf:"bytes,4,opt,name=image_file_md5,json=imageFileMd5,proto3" json:"image_file_md5,omitempty"`
	// 额外变量
	ExtraVar             string   `protobuf:"bytes,6,opt,name=extra_var,json=extraVar,proto3" json:"extra_var,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BareMetalInstallTask) Reset()         { *m = BareMetalInstallTask{} }
func (m *BareMetalInstallTask) String() string { return proto.CompactTextString(m) }
func (*BareMetalInstallTask) ProtoMessage()    {}
func (*BareMetalInstallTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_840e7c8beea44e4e, []int{5}
}

func (m *BareMetalInstallTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalInstallTask.Unmarshal(m, b)
}
func (m *BareMetalInstallTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalInstallTask.Marshal(b, m, deterministic)
}
func (m *BareMetalInstallTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalInstallTask.Merge(m, src)
}
func (m *BareMetalInstallTask) XXX_Size() int {
	return xxx_messageInfo_BareMetalInstallTask.Size(m)
}
func (m *BareMetalInstallTask) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalInstallTask.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalInstallTask proto.InternalMessageInfo

func (m *BareMetalInstallTask) GetSubTaskId() int64 {
	if m != nil {
		return m.SubTaskId
	}
	return 0
}

func (m *BareMetalInstallTask) GetCmdbSearchRequest() *cmdb.ChooseServerRequest {
	if m != nil {
		return m.CmdbSearchRequest
	}
	return nil
}

func (m *BareMetalInstallTask) GetImageFileUrl() string {
	if m != nil {
		return m.ImageFileUrl
	}
	return ""
}

func (m *BareMetalInstallTask) GetImageFileMd5() string {
	if m != nil {
		return m.ImageFileMd5
	}
	return ""
}

func (m *BareMetalInstallTask) GetExtraVar() string {
	if m != nil {
		return m.ExtraVar
	}
	return ""
}

func init() {
	proto.RegisterType((*TemplateExecuteTask)(nil), "task.TemplateExecuteTask")
	proto.RegisterType((*ServerCompareTask)(nil), "task.ServerCompareTask")
	proto.RegisterType((*BareMetalSearchTask)(nil), "task.BareMetalSearchTask")
	proto.RegisterType((*BareMetalPowerTask)(nil), "task.BareMetalPowerTask")
	proto.RegisterType((*BareMetalCreateTask)(nil), "task.BareMetalCreateTask")
	proto.RegisterType((*BareMetalInstallTask)(nil), "task.BareMetalInstallTask")
}

func init() { proto.RegisterFile("task/queue.proto", fileDescriptor_840e7c8beea44e4e) }

var fileDescriptor_840e7c8beea44e4e = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0xe5, 0xa6, 0xad, 0xc8, 0x14, 0xb5, 0xc4, 0x41, 0x60, 0x02, 0x42, 0x51, 0xd4, 0x45,
	0x84, 0xda, 0x58, 0x2a, 0xea, 0x0b, 0x24, 0xe2, 0x27, 0x8b, 0xa2, 0xca, 0x29, 0x48, 0xb0, 0xb1,
	0xae, 0xed, 0x5b, 0xd7, 0xca, 0x8c, 0xaf, 0x3b, 0x33, 0x6e, 0xd3, 0x1d, 0x1b, 0x9e, 0x83, 0x27,
	0xe0, 0xc5, 0x78, 0x0a, 0x34, 0xe3, 0xa6, 0xc5, 0x84, 0x46, 0x11, 0x42, 0xdd, 0xd8, 0xd2, 0x39,
	0x9f, 0xe7, 0x1e, 0x1f, 0xf9, 0x9a, 0x3d, 0xd2, 0xa0, 0xa6, 0xfe, 0x79, 0x89, 0x25, 0x0e, 0x0a,
	0x49, 0x9a, 0xdc, 0x75, 0xa3, 0x74, 0x5e, 0xa4, 0x44, 0x29, 0x47, 0x1f, 0x8a, 0xcc, 0x87, 0x3c,
	0x27, 0x0d, 0x3a, 0xa3, 0x5c, 0x55, 0x4c, 0x67, 0xcf, 0xde, 0xe2, 0xfd, 0x14, 0xf3, 0x7d, 0x75,
	0x09, 0x69, 0x8a, 0xd2, 0xa7, 0xc2, 0x12, 0x7f, 0xa1, 0x77, 0x62, 0x91, 0x44, 0xbe, 0xb9, 0x54,
	0x42, 0xef, 0xfb, 0x1a, 0x6b, 0x9f, 0xa0, 0x28, 0x38, 0x68, 0x7c, 0x33, 0xc3, 0xb8, 0xd4, 0x78,
	0x02, 0x6a, 0xea, 0xbe, 0x64, 0x5b, 0xaa, 0x8c, 0x42, 0x13, 0x20, 0xcc, 0x12, 0xcf, 0xe9, 0x3a,
	0xfd, 0x46, 0xd0, 0x54, 0x65, 0x64, 0xdc, 0x71, 0xe2, 0xbe, 0x63, 0x6d, 0x73, 0x4a, 0xa8, 0x10,
	0x64, 0x7c, 0x16, 0x4a, 0x3c, 0x2f, 0x51, 0x69, 0x6f, 0xad, 0xeb, 0xf4, 0xb7, 0x0e, 0x9e, 0x0e,
	0xec, 0x84, 0xd1, 0x19, 0x91, 0xc2, 0xf7, 0xa4, 0x74, 0x50, 0xd9, 0x41, 0xcb, 0xe8, 0x13, 0xfb,
	0xc8, 0xb5, 0xe4, 0xbe, 0x62, 0xad, 0x82, 0xc3, 0x55, 0x44, 0x34, 0x0d, 0x4f, 0x33, 0x8e, 0x61,
	0x29, 0xb9, 0xd7, 0xe8, 0x3a, 0xfd, 0x66, 0xb0, 0x33, 0x37, 0xde, 0x66, 0x1c, 0x3f, 0x4a, 0xbe,
	0xc8, 0x8a, 0xe4, 0xd0, 0x5b, 0x5f, 0x64, 0x8f, 0x92, 0xc3, 0x1a, 0x7b, 0x25, 0x78, 0x98, 0x83,
	0x40, 0x6f, 0xa3, 0xce, 0x7e, 0x16, 0xfc, 0x03, 0x08, 0x74, 0x9f, 0xb3, 0x26, 0xce, 0xb4, 0x84,
	0xf0, 0x02, 0xa4, 0xb7, 0x69, 0x99, 0x07, 0x56, 0xf8, 0x04, 0xb2, 0xf7, 0xcd, 0x61, 0xad, 0x09,
	0xca, 0x0b, 0x94, 0x23, 0x12, 0x05, 0xc8, 0xd5, 0xfa, 0x39, 0x5e, 0xd6, 0x4f, 0xf7, 0xf7, 0x7e,
	0x6a, 0x67, 0xdf, 0x5d, 0x54, 0xef, 0xab, 0xc3, 0xda, 0x43, 0x90, 0x78, 0x84, 0x1a, 0x78, 0x65,
	0xad, 0x94, 0x64, 0xbc, 0x2c, 0xc9, 0xb3, 0xc5, 0x24, 0x4b, 0x22, 0xfc, 0x70, 0x98, 0x7b, 0x13,
	0xe1, 0x98, 0x2e, 0x51, 0xae, 0x94, 0x60, 0x8f, 0x6d, 0x28, 0x0d, 0x1a, 0xed, 0xcc, 0xed, 0x83,
	0x27, 0xd5, 0xcc, 0x6a, 0x9a, 0x3d, 0x65, 0x62, 0xdc, 0xa0, 0x82, 0xee, 0xca, 0xdb, 0xf8, 0x87,
	0xbc, 0xb5, 0xca, 0x46, 0x12, 0x61, 0xc5, 0x8f, 0xfb, 0x3f, 0x56, 0xf6, 0xd3, 0x61, 0x8f, 0x6f,
	0x22, 0x8c, 0x73, 0xa5, 0x81, 0xf3, 0x7b, 0xce, 0xe0, 0xee, 0xb2, 0xed, 0x4c, 0x40, 0x8a, 0x7f,
	0xee, 0xd7, 0x43, 0xab, 0xce, 0x97, 0xab, 0x4e, 0xdd, 0x6e, 0xd6, 0x2d, 0x65, 0xd6, 0x6a, 0xd9,
	0xaa, 0x0c, 0x87, 0xac, 0xa3, 0xa9, 0x18, 0x9c, 0x52, 0x1a, 0x27, 0xf9, 0x00, 0x66, 0xd7, 0xff,
	0x31, 0x35, 0x30, 0x6f, 0xf9, 0x65, 0x37, 0xcd, 0xf4, 0xdc, 0xd3, 0x54, 0xf8, 0x30, 0x43, 0xbf,
	0xf2, 0xfd, 0x94, 0xa8, 0xd4, 0xbe, 0xa1, 0xa2, 0x4d, 0x2b, 0xbd, 0xfe, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x73, 0x60, 0x85, 0x90, 0x0e, 0x05, 0x00, 0x00,
}
