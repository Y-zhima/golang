// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmdb/cmdb.proto

package cmdb

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

// 业务拓扑对象
type TopologyObject struct {
	BkInstId             int32             `protobuf:"varint,1,opt,name=bk_inst_id,json=bkInstId,proto3" json:"bk_inst_id,omitempty"`
	BkInstName           string            `protobuf:"bytes,2,opt,name=bk_inst_name,json=bkInstName,proto3" json:"bk_inst_name,omitempty"`
	BkObjId              string            `protobuf:"bytes,3,opt,name=bk_obj_id,json=bkObjId,proto3" json:"bk_obj_id,omitempty"`
	BkObjName            string            `protobuf:"bytes,4,opt,name=bk_obj_name,json=bkObjName,proto3" json:"bk_obj_name,omitempty"`
	Child                []*TopologyObject `protobuf:"bytes,5,rep,name=child,proto3" json:"child,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TopologyObject) Reset()         { *m = TopologyObject{} }
func (m *TopologyObject) String() string { return proto.CompactTextString(m) }
func (*TopologyObject) ProtoMessage()    {}
func (*TopologyObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e1f4eaffd389ae, []int{0}
}

func (m *TopologyObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopologyObject.Unmarshal(m, b)
}
func (m *TopologyObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopologyObject.Marshal(b, m, deterministic)
}
func (m *TopologyObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopologyObject.Merge(m, src)
}
func (m *TopologyObject) XXX_Size() int {
	return xxx_messageInfo_TopologyObject.Size(m)
}
func (m *TopologyObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TopologyObject.DiscardUnknown(m)
}

var xxx_messageInfo_TopologyObject proto.InternalMessageInfo

func (m *TopologyObject) GetBkInstId() int32 {
	if m != nil {
		return m.BkInstId
	}
	return 0
}

func (m *TopologyObject) GetBkInstName() string {
	if m != nil {
		return m.BkInstName
	}
	return ""
}

func (m *TopologyObject) GetBkObjId() string {
	if m != nil {
		return m.BkObjId
	}
	return ""
}

func (m *TopologyObject) GetBkObjName() string {
	if m != nil {
		return m.BkObjName
	}
	return ""
}

func (m *TopologyObject) GetChild() []*TopologyObject {
	if m != nil {
		return m.Child
	}
	return nil
}

// 拓扑实例请求
type InstanceTopologyRequest struct {
	Level                int32    `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceTopologyRequest) Reset()         { *m = InstanceTopologyRequest{} }
func (m *InstanceTopologyRequest) String() string { return proto.CompactTextString(m) }
func (*InstanceTopologyRequest) ProtoMessage()    {}
func (*InstanceTopologyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e1f4eaffd389ae, []int{1}
}

func (m *InstanceTopologyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceTopologyRequest.Unmarshal(m, b)
}
func (m *InstanceTopologyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceTopologyRequest.Marshal(b, m, deterministic)
}
func (m *InstanceTopologyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceTopologyRequest.Merge(m, src)
}
func (m *InstanceTopologyRequest) XXX_Size() int {
	return xxx_messageInfo_InstanceTopologyRequest.Size(m)
}
func (m *InstanceTopologyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceTopologyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceTopologyRequest proto.InternalMessageInfo

func (m *InstanceTopologyRequest) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

// 拓扑实例请求返回
type InstanceTopologyResponse struct {
	Instance             *TopologyObject        `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Status               *common.ResponseStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *InstanceTopologyResponse) Reset()         { *m = InstanceTopologyResponse{} }
func (m *InstanceTopologyResponse) String() string { return proto.CompactTextString(m) }
func (*InstanceTopologyResponse) ProtoMessage()    {}
func (*InstanceTopologyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e1f4eaffd389ae, []int{2}
}

func (m *InstanceTopologyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceTopologyResponse.Unmarshal(m, b)
}
func (m *InstanceTopologyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceTopologyResponse.Marshal(b, m, deterministic)
}
func (m *InstanceTopologyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceTopologyResponse.Merge(m, src)
}
func (m *InstanceTopologyResponse) XXX_Size() int {
	return xxx_messageInfo_InstanceTopologyResponse.Size(m)
}
func (m *InstanceTopologyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceTopologyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceTopologyResponse proto.InternalMessageInfo

func (m *InstanceTopologyResponse) GetInstance() *TopologyObject {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *InstanceTopologyResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type HostObject struct {
	BkHostInnerip        string   `protobuf:"bytes,1,opt,name=bk_host_innerip,json=bkHostInnerip,proto3" json:"bk_host_innerip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostObject) Reset()         { *m = HostObject{} }
func (m *HostObject) String() string { return proto.CompactTextString(m) }
func (*HostObject) ProtoMessage()    {}
func (*HostObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e1f4eaffd389ae, []int{3}
}

func (m *HostObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostObject.Unmarshal(m, b)
}
func (m *HostObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostObject.Marshal(b, m, deterministic)
}
func (m *HostObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostObject.Merge(m, src)
}
func (m *HostObject) XXX_Size() int {
	return xxx_messageInfo_HostObject.Size(m)
}
func (m *HostObject) XXX_DiscardUnknown() {
	xxx_messageInfo_HostObject.DiscardUnknown(m)
}

var xxx_messageInfo_HostObject proto.InternalMessageInfo

func (m *HostObject) GetBkHostInnerip() string {
	if m != nil {
		return m.BkHostInnerip
	}
	return ""
}

type ModuleObject struct {
	BkModuleName         string   `protobuf:"bytes,1,opt,name=bk_module_name,json=bkModuleName,proto3" json:"bk_module_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModuleObject) Reset()         { *m = ModuleObject{} }
func (m *ModuleObject) String() string { return proto.CompactTextString(m) }
func (*ModuleObject) ProtoMessage()    {}
func (*ModuleObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e1f4eaffd389ae, []int{4}
}

func (m *ModuleObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModuleObject.Unmarshal(m, b)
}
func (m *ModuleObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModuleObject.Marshal(b, m, deterministic)
}
func (m *ModuleObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleObject.Merge(m, src)
}
func (m *ModuleObject) XXX_Size() int {
	return xxx_messageInfo_ModuleObject.Size(m)
}
func (m *ModuleObject) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleObject.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleObject proto.InternalMessageInfo

func (m *ModuleObject) GetBkModuleName() string {
	if m != nil {
		return m.BkModuleName
	}
	return ""
}

type SetObject struct {
	BkSetName            string   `protobuf:"bytes,1,opt,name=bk_set_name,json=bkSetName,proto3" json:"bk_set_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetObject) Reset()         { *m = SetObject{} }
func (m *SetObject) String() string { return proto.CompactTextString(m) }
func (*SetObject) ProtoMessage()    {}
func (*SetObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e1f4eaffd389ae, []int{5}
}

func (m *SetObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetObject.Unmarshal(m, b)
}
func (m *SetObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetObject.Marshal(b, m, deterministic)
}
func (m *SetObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetObject.Merge(m, src)
}
func (m *SetObject) XXX_Size() int {
	return xxx_messageInfo_SetObject.Size(m)
}
func (m *SetObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SetObject.DiscardUnknown(m)
}

var xxx_messageInfo_SetObject proto.InternalMessageInfo

func (m *SetObject) GetBkSetName() string {
	if m != nil {
		return m.BkSetName
	}
	return ""
}

type BizObject struct {
	BkBizName            string   `protobuf:"bytes,1,opt,name=bk_biz_name,json=bkBizName,proto3" json:"bk_biz_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BizObject) Reset()         { *m = BizObject{} }
func (m *BizObject) String() string { return proto.CompactTextString(m) }
func (*BizObject) ProtoMessage()    {}
func (*BizObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e1f4eaffd389ae, []int{6}
}

func (m *BizObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BizObject.Unmarshal(m, b)
}
func (m *BizObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BizObject.Marshal(b, m, deterministic)
}
func (m *BizObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BizObject.Merge(m, src)
}
func (m *BizObject) XXX_Size() int {
	return xxx_messageInfo_BizObject.Size(m)
}
func (m *BizObject) XXX_DiscardUnknown() {
	xxx_messageInfo_BizObject.DiscardUnknown(m)
}

var xxx_messageInfo_BizObject proto.InternalMessageInfo

func (m *BizObject) GetBkBizName() string {
	if m != nil {
		return m.BkBizName
	}
	return ""
}

type HostInfoObject struct {
	Host                 *HostObject     `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Module               []*ModuleObject `protobuf:"bytes,2,rep,name=module,proto3" json:"module,omitempty"`
	Set                  []*SetObject    `protobuf:"bytes,3,rep,name=set,proto3" json:"set,omitempty"`
	Biz                  []*BizObject    `protobuf:"bytes,4,rep,name=biz,proto3" json:"biz,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *HostInfoObject) Reset()         { *m = HostInfoObject{} }
func (m *HostInfoObject) String() string { return proto.CompactTextString(m) }
func (*HostInfoObject) ProtoMessage()    {}
func (*HostInfoObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e1f4eaffd389ae, []int{7}
}

func (m *HostInfoObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostInfoObject.Unmarshal(m, b)
}
func (m *HostInfoObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostInfoObject.Marshal(b, m, deterministic)
}
func (m *HostInfoObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostInfoObject.Merge(m, src)
}
func (m *HostInfoObject) XXX_Size() int {
	return xxx_messageInfo_HostInfoObject.Size(m)
}
func (m *HostInfoObject) XXX_DiscardUnknown() {
	xxx_messageInfo_HostInfoObject.DiscardUnknown(m)
}

var xxx_messageInfo_HostInfoObject proto.InternalMessageInfo

func (m *HostInfoObject) GetHost() *HostObject {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *HostInfoObject) GetModule() []*ModuleObject {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *HostInfoObject) GetSet() []*SetObject {
	if m != nil {
		return m.Set
	}
	return nil
}

func (m *HostInfoObject) GetBiz() []*BizObject {
	if m != nil {
		return m.Biz
	}
	return nil
}

// 查找主机请求
type SearchHostRequest struct {
	Host                 *HostObject   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Module               *ModuleObject `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	Set                  *SetObject    `protobuf:"bytes,3,opt,name=set,proto3" json:"set,omitempty"`
	Biz                  *BizObject    `protobuf:"bytes,4,opt,name=biz,proto3" json:"biz,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SearchHostRequest) Reset()         { *m = SearchHostRequest{} }
func (m *SearchHostRequest) String() string { return proto.CompactTextString(m) }
func (*SearchHostRequest) ProtoMessage()    {}
func (*SearchHostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e1f4eaffd389ae, []int{8}
}

func (m *SearchHostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchHostRequest.Unmarshal(m, b)
}
func (m *SearchHostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchHostRequest.Marshal(b, m, deterministic)
}
func (m *SearchHostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchHostRequest.Merge(m, src)
}
func (m *SearchHostRequest) XXX_Size() int {
	return xxx_messageInfo_SearchHostRequest.Size(m)
}
func (m *SearchHostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchHostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchHostRequest proto.InternalMessageInfo

func (m *SearchHostRequest) GetHost() *HostObject {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *SearchHostRequest) GetModule() *ModuleObject {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *SearchHostRequest) GetSet() *SetObject {
	if m != nil {
		return m.Set
	}
	return nil
}

func (m *SearchHostRequest) GetBiz() *BizObject {
	if m != nil {
		return m.Biz
	}
	return nil
}

// 查找主机请求返回
type SearchHostResponse struct {
	Count                int32                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Info                 []*HostInfoObject      `protobuf:"bytes,2,rep,name=info,proto3" json:"info,omitempty"`
	Status               *common.ResponseStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SearchHostResponse) Reset()         { *m = SearchHostResponse{} }
func (m *SearchHostResponse) String() string { return proto.CompactTextString(m) }
func (*SearchHostResponse) ProtoMessage()    {}
func (*SearchHostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e1f4eaffd389ae, []int{9}
}

func (m *SearchHostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchHostResponse.Unmarshal(m, b)
}
func (m *SearchHostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchHostResponse.Marshal(b, m, deterministic)
}
func (m *SearchHostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchHostResponse.Merge(m, src)
}
func (m *SearchHostResponse) XXX_Size() int {
	return xxx_messageInfo_SearchHostResponse.Size(m)
}
func (m *SearchHostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchHostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchHostResponse proto.InternalMessageInfo

func (m *SearchHostResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SearchHostResponse) GetInfo() []*HostInfoObject {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *SearchHostResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*TopologyObject)(nil), "cmdb.TopologyObject")
	proto.RegisterType((*InstanceTopologyRequest)(nil), "cmdb.InstanceTopologyRequest")
	proto.RegisterType((*InstanceTopologyResponse)(nil), "cmdb.InstanceTopologyResponse")
	proto.RegisterType((*HostObject)(nil), "cmdb.HostObject")
	proto.RegisterType((*ModuleObject)(nil), "cmdb.ModuleObject")
	proto.RegisterType((*SetObject)(nil), "cmdb.SetObject")
	proto.RegisterType((*BizObject)(nil), "cmdb.BizObject")
	proto.RegisterType((*HostInfoObject)(nil), "cmdb.HostInfoObject")
	proto.RegisterType((*SearchHostRequest)(nil), "cmdb.SearchHostRequest")
	proto.RegisterType((*SearchHostResponse)(nil), "cmdb.SearchHostResponse")
}

func init() { proto.RegisterFile("cmdb/cmdb.proto", fileDescriptor_02e1f4eaffd389ae) }

var fileDescriptor_02e1f4eaffd389ae = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xd1, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0x95, 0xb5, 0xdd, 0x7f, 0x3d, 0xdb, 0x7f, 0x1b, 0xa6, 0xb0, 0xd0, 0x8d, 0xa9, 0x8b,
	0x26, 0x54, 0x0d, 0x96, 0xc0, 0xd8, 0x13, 0x94, 0x1b, 0x7a, 0x01, 0x93, 0x52, 0xae, 0x10, 0x52,
	0x15, 0x27, 0x5e, 0xe6, 0xa5, 0xf1, 0x09, 0xb5, 0x3b, 0x46, 0xc5, 0x15, 0xe2, 0x0d, 0x78, 0x09,
	0x24, 0x5e, 0x80, 0xf7, 0xe0, 0x8e, 0x6b, 0x1e, 0x04, 0xc5, 0x4e, 0xda, 0xae, 0x55, 0xc5, 0xb8,
	0xd9, 0x1a, 0x7f, 0xbf, 0xe3, 0xf3, 0x9d, 0xcf, 0x4e, 0x60, 0x2b, 0x4c, 0x23, 0xea, 0xe5, 0x7f,
	0xdc, 0x6c, 0x88, 0x0a, 0x49, 0x35, 0xff, 0xdd, 0xdc, 0x0e, 0x31, 0x4d, 0x51, 0x78, 0x41, 0xc6,
	0xcd, 0x7a, 0x73, 0x2f, 0x46, 0x8c, 0x07, 0x2c, 0x5f, 0xf1, 0x02, 0x21, 0x50, 0x05, 0x8a, 0xa3,
	0x90, 0x85, 0xfa, 0x44, 0xff, 0x0b, 0x8f, 0x63, 0x26, 0x8e, 0xe5, 0x87, 0x20, 0x8e, 0xd9, 0xd0,
	0xc3, 0x4c, 0x13, 0x8b, 0xb4, 0xf3, 0xc3, 0x82, 0xcd, 0x37, 0x98, 0xe1, 0x00, 0xe3, 0x8f, 0x67,
	0xf4, 0x92, 0x85, 0x8a, 0xec, 0x01, 0xd0, 0xa4, 0xcf, 0x85, 0x54, 0x7d, 0x1e, 0xd9, 0x56, 0xcb,
	0x6a, 0xd7, 0xfc, 0x35, 0x9a, 0x74, 0x85, 0x54, 0xdd, 0x88, 0xb4, 0x60, 0xa3, 0x54, 0x45, 0x90,
	0x32, 0x7b, 0xa5, 0x65, 0xb5, 0xeb, 0x3e, 0x18, 0xfd, 0x75, 0x90, 0x32, 0xd2, 0x84, 0x3a, 0x4d,
	0xfa, 0x48, 0x2f, 0xf3, 0xf2, 0x8a, 0x96, 0xff, 0xa3, 0xc9, 0x19, 0xbd, 0xec, 0x46, 0x64, 0x1f,
	0xd6, 0x0b, 0x4d, 0x17, 0x57, 0xb5, 0x5a, 0xd7, 0xaa, 0xae, 0x3d, 0x82, 0x5a, 0x78, 0xc1, 0x07,
	0x91, 0x5d, 0x6b, 0x55, 0xda, 0xeb, 0x27, 0x0d, 0x57, 0xc7, 0x71, 0xd3, 0xa0, 0x6f, 0x10, 0xc7,
	0x83, 0x9d, 0xbc, 0x67, 0x20, 0x42, 0x56, 0x02, 0x3e, 0x7b, 0x3f, 0x62, 0x52, 0x91, 0x06, 0xd4,
	0x06, 0xec, 0x8a, 0x0d, 0x0a, 0xf7, 0xe6, 0xc1, 0xf9, 0x04, 0xf6, 0x62, 0x81, 0xcc, 0x50, 0x48,
	0x46, 0x9e, 0xc2, 0x1a, 0x2f, 0x34, 0x5d, 0xb4, 0xac, 0xf7, 0x84, 0x22, 0x2e, 0xac, 0x4a, 0x15,
	0xa8, 0x91, 0xd4, 0x11, 0xac, 0x9f, 0xdc, 0x77, 0xcd, 0x41, 0xb9, 0xe5, 0x9e, 0x3d, 0xad, 0xfa,
	0x05, 0xe5, 0x9c, 0x02, 0xbc, 0x44, 0xa9, 0x8a, 0x90, 0x1f, 0xc1, 0x16, 0x4d, 0xfa, 0x17, 0x98,
	0x87, 0x2c, 0x04, 0x1b, 0xf2, 0x4c, 0xb7, 0xad, 0xfb, 0xff, 0xd3, 0x24, 0xc7, 0xba, 0x66, 0xd1,
	0x39, 0x85, 0x8d, 0x57, 0x18, 0x8d, 0x06, 0xac, 0xa8, 0x3b, 0x84, 0x4d, 0x9a, 0xf4, 0x53, 0xbd,
	0x64, 0x32, 0x34, 0x65, 0x1b, 0x34, 0x31, 0x5c, 0x1e, 0xa3, 0xf3, 0x18, 0xea, 0x3d, 0x56, 0xb6,
	0x32, 0x99, 0x4b, 0xa6, 0x66, 0xf9, 0x3a, 0x4d, 0x7a, 0x4c, 0x95, 0x70, 0x87, 0x8f, 0x6f, 0xc0,
	0x94, 0x8f, 0xe7, 0xe0, 0x0e, 0x1f, 0x6b, 0xf8, 0x9b, 0x05, 0x9b, 0xc6, 0xdf, 0x39, 0x4e, 0x2c,
	0x55, 0xf3, 0x39, 0x8a, 0xd8, 0xb6, 0x4d, 0x6c, 0xd3, 0x51, 0x7d, 0xad, 0x92, 0x23, 0x58, 0x35,
	0xae, 0xed, 0x15, 0x7d, 0xb4, 0xc4, 0x70, 0xb3, 0xc3, 0xf9, 0x05, 0x41, 0x0e, 0xa0, 0x22, 0x99,
	0xb2, 0x2b, 0x1a, 0xdc, 0x32, 0xe0, 0x64, 0x1e, 0x3f, 0xd7, 0x72, 0x84, 0xf2, 0xb1, 0x5d, 0x9d,
	0x45, 0x26, 0x53, 0xf8, 0xb9, 0xe6, 0x7c, 0xb7, 0xe0, 0x4e, 0x8f, 0x05, 0xc3, 0xf0, 0x22, 0x37,
	0x53, 0x5e, 0x8d, 0x7f, 0x77, 0x6b, 0xdd, 0xd6, 0xad, 0xf5, 0x77, 0xb7, 0xd6, 0x52, 0xb7, 0x5f,
	0x2c, 0x20, 0xb3, 0x6e, 0x8b, 0x7b, 0xd9, 0x80, 0x5a, 0x88, 0x23, 0xa1, 0xca, 0x9b, 0xac, 0x1f,
	0x48, 0x1b, 0xaa, 0x5c, 0x9c, 0x63, 0x11, 0x65, 0x63, 0x3a, 0xc4, 0xf4, 0x58, 0x7c, 0x4d, 0xcc,
	0xdc, 0xd2, 0xca, 0x6d, 0x6e, 0xe9, 0xc9, 0x2f, 0x0b, 0xaa, 0x2f, 0xd2, 0x88, 0x12, 0x05, 0xdb,
	0xf3, 0x2f, 0x0b, 0x79, 0x68, 0x1a, 0x2d, 0x79, 0xeb, 0x9a, 0xfb, 0xcb, 0x64, 0xd3, 0xc9, 0x39,
	0xf8, 0xfc, 0xf3, 0xf7, 0xd7, 0x95, 0x5d, 0xf2, 0xc0, 0xbb, 0x7a, 0xa6, 0xbf, 0x73, 0x1e, 0x9f,
	0xef, 0xf0, 0x0e, 0x60, 0x1a, 0x02, 0xd9, 0x29, 0xc3, 0x9c, 0x3b, 0xc4, 0xa6, 0xbd, 0x28, 0x14,
	0x3d, 0x76, 0x75, 0x8f, 0x7b, 0xe4, 0xee, 0xa4, 0x87, 0x9c, 0x40, 0x9d, 0x0e, 0x34, 0x15, 0x66,
	0xee, 0x39, 0xc6, 0x61, 0x24, 0xdc, 0xe0, 0x9a, 0x99, 0x8f, 0xa0, 0xd4, 0xbb, 0xbd, 0x3d, 0x8c,
	0xb9, 0x2a, 0x35, 0x85, 0x99, 0x17, 0x5c, 0x33, 0xcf, 0xe8, 0x5e, 0x8c, 0x38, 0x52, 0x7a, 0x3b,
	0xba, 0xaa, 0x97, 0x9e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xaa, 0x81, 0x35, 0xae, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CmdbClient is the client API for Cmdb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CmdbClient interface {
	// 实例拓扑
	InstanceTopology(ctx context.Context, in *InstanceTopologyRequest, opts ...grpc.CallOption) (*InstanceTopologyResponse, error)
	// 查找主机
	SearchHost(ctx context.Context, in *SearchHostRequest, opts ...grpc.CallOption) (*SearchHostResponse, error)
}

type cmdbClient struct {
	cc *grpc.ClientConn
}

func NewCmdbClient(cc *grpc.ClientConn) CmdbClient {
	return &cmdbClient{cc}
}

func (c *cmdbClient) InstanceTopology(ctx context.Context, in *InstanceTopologyRequest, opts ...grpc.CallOption) (*InstanceTopologyResponse, error) {
	out := new(InstanceTopologyResponse)
	err := c.cc.Invoke(ctx, "/cmdb.Cmdb/InstanceTopology", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdbClient) SearchHost(ctx context.Context, in *SearchHostRequest, opts ...grpc.CallOption) (*SearchHostResponse, error) {
	out := new(SearchHostResponse)
	err := c.cc.Invoke(ctx, "/cmdb.Cmdb/SearchHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmdbServer is the server API for Cmdb service.
type CmdbServer interface {
	// 实例拓扑
	InstanceTopology(context.Context, *InstanceTopologyRequest) (*InstanceTopologyResponse, error)
	// 查找主机
	SearchHost(context.Context, *SearchHostRequest) (*SearchHostResponse, error)
}

// UnimplementedCmdbServer can be embedded to have forward compatible implementations.
type UnimplementedCmdbServer struct {
}

func (*UnimplementedCmdbServer) InstanceTopology(ctx context.Context, req *InstanceTopologyRequest) (*InstanceTopologyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstanceTopology not implemented")
}
func (*UnimplementedCmdbServer) SearchHost(ctx context.Context, req *SearchHostRequest) (*SearchHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchHost not implemented")
}

func RegisterCmdbServer(s *grpc.Server, srv CmdbServer) {
	s.RegisterService(&_Cmdb_serviceDesc, srv)
}

func _Cmdb_InstanceTopology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceTopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).InstanceTopology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cmdb.Cmdb/InstanceTopology",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).InstanceTopology(ctx, req.(*InstanceTopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cmdb_SearchHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdbServer).SearchHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cmdb.Cmdb/SearchHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdbServer).SearchHost(ctx, req.(*SearchHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cmdb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cmdb.Cmdb",
	HandlerType: (*CmdbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InstanceTopology",
			Handler:    _Cmdb_InstanceTopology_Handler,
		},
		{
			MethodName: "SearchHost",
			Handler:    _Cmdb_SearchHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdb/cmdb.proto",
}
