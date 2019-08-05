// Code generated by protoc-gen-go. DO NOT EDIT.
// source: template/template.proto

package template

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

// 模板实例
type TemplateObject struct {
	TemplateId           int32    `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateObject) Reset()         { *m = TemplateObject{} }
func (m *TemplateObject) String() string { return proto.CompactTextString(m) }
func (*TemplateObject) ProtoMessage()    {}
func (*TemplateObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{0}
}

func (m *TemplateObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateObject.Unmarshal(m, b)
}
func (m *TemplateObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateObject.Marshal(b, m, deterministic)
}
func (m *TemplateObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateObject.Merge(m, src)
}
func (m *TemplateObject) XXX_Size() int {
	return xxx_messageInfo_TemplateObject.Size(m)
}
func (m *TemplateObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateObject.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateObject proto.InternalMessageInfo

func (m *TemplateObject) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
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
	return fileDescriptor_dca67df6b60706ce, []int{1}
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

// 创建模板请求返回
type CreateResponse struct {
	TemplateId           int32                  `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Status               *common.ResponseStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{2}
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

func (m *CreateResponse) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
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
	TemplateId           int32    `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{3}
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

func (m *GetRequest) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

// 获取模板请求返回
type GetResponse struct {
	Template             *TemplateObject        `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	Status               *common.ResponseStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{4}
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

func (m *GetResponse) GetTemplate() *TemplateObject {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *GetResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// 筛选模板请求
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
	return fileDescriptor_dca67df6b60706ce, []int{5}
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

// 筛选模板请求返回
type FilterResponse struct {
	Templates            []*TemplateObject      `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
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
	return fileDescriptor_dca67df6b60706ce, []int{6}
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

func (m *FilterResponse) GetTemplates() []*TemplateObject {
	if m != nil {
		return m.Templates
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

type UpdateRequest struct {
	TemplateId           int32    `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{7}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{8}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type DeleteRequest struct {
	TemplateId           int32    `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dca67df6b60706ce, []int{9}
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

func (m *DeleteRequest) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
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
	return fileDescriptor_dca67df6b60706ce, []int{10}
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
	proto.RegisterType((*TemplateObject)(nil), "proto.TemplateObject")
	proto.RegisterType((*CreateRequest)(nil), "proto.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "proto.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "proto.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "proto.GetResponse")
	proto.RegisterType((*FilterRequest)(nil), "proto.FilterRequest")
	proto.RegisterType((*FilterResponse)(nil), "proto.FilterResponse")
	proto.RegisterType((*UpdateRequest)(nil), "proto.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "proto.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "proto.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "proto.DeleteResponse")
}

func init() { proto.RegisterFile("template/template.proto", fileDescriptor_dca67df6b60706ce) }

var fileDescriptor_dca67df6b60706ce = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0xd9, 0x84, 0x2c, 0xf5, 0xc5, 0xac, 0xe9, 0xd0, 0xda, 0xb0, 0x16, 0xba, 0x2c, 0x28,
	0x21, 0x98, 0xdd, 0x24, 0x3d, 0x08, 0x3d, 0x19, 0xb5, 0x96, 0x20, 0x68, 0xd8, 0x56, 0x50, 0x11,
	0x64, 0x9b, 0x4c, 0x97, 0x2d, 0x9b, 0x9d, 0x31, 0x3b, 0x51, 0x41, 0xbc, 0xf8, 0x05, 0x14, 0x05,
	0x0f, 0x62, 0x0f, 0xde, 0x3c, 0x8a, 0x7a, 0x50, 0x44, 0xfa, 0x21, 0xfc, 0x0a, 0x69, 0xf5, 0x63,
	0x48, 0x66, 0x67, 0x36, 0xd9, 0x16, 0x93, 0xf6, 0x34, 0xe1, 0xff, 0xde, 0xfc, 0xde, 0xff, 0xbd,
	0x37, 0x59, 0x58, 0x62, 0xb8, 0x47, 0x03, 0x97, 0x61, 0x5b, 0xfe, 0xb0, 0x68, 0x9f, 0x30, 0x82,
	0x72, 0xfc, 0xd0, 0x8b, 0x1d, 0xd2, 0xeb, 0x91, 0xd0, 0x76, 0xa9, 0x1f, 0x07, 0xf4, 0x65, 0x8f,
	0x10, 0x2f, 0xc0, 0x23, 0xc5, 0x76, 0xc3, 0x90, 0x30, 0x97, 0xf9, 0x24, 0x8c, 0x44, 0xf4, 0x32,
	0x3f, 0x3a, 0x55, 0x0f, 0x87, 0xd5, 0xe8, 0xa9, 0xeb, 0x79, 0xb8, 0x6f, 0x13, 0xca, 0x33, 0x8e,
	0x67, 0x9b, 0x75, 0xd0, 0xb6, 0x44, 0xd9, 0x3b, 0xdb, 0xbb, 0xb8, 0xc3, 0xd0, 0x0a, 0xe4, 0xa5,
	0x91, 0x47, 0x7e, 0xb7, 0xa4, 0x18, 0x4a, 0x39, 0xe7, 0x80, 0x94, 0x5a, 0x5d, 0xf3, 0x1c, 0x14,
	0xae, 0xf7, 0xb1, 0xcb, 0xb0, 0x83, 0x1f, 0x0f, 0x70, 0xc4, 0x4c, 0x17, 0x34, 0x29, 0x44, 0x94,
	0x84, 0x11, 0x9e, 0xc9, 0x40, 0x16, 0xa8, 0x11, 0x73, 0xd9, 0x20, 0x2a, 0x65, 0x0c, 0xa5, 0x9c,
	0x6f, 0x9c, 0xb7, 0xe2, 0x2e, 0x2d, 0x89, 0xd8, 0xe4, 0x51, 0x47, 0x64, 0x99, 0x55, 0x80, 0x0d,
	0xcc, 0x44, 0xc1, 0xd9, 0x16, 0x29, 0xe4, 0x79, 0xba, 0xb0, 0x53, 0x87, 0x39, 0x19, 0xe4, 0xc9,
	0xf9, 0xc6, 0x62, 0xdc, 0xbe, 0x95, 0xee, 0xdd, 0x49, 0xd2, 0x4e, 0x6d, 0xf0, 0x0a, 0x14, 0x6e,
	0xfa, 0x01, 0xc3, 0x7d, 0xe9, 0xf1, 0x12, 0xa8, 0xd4, 0xf5, 0xfc, 0xd0, 0x13, 0x15, 0x35, 0x09,
	0x68, 0x73, 0xd5, 0x11, 0x51, 0x73, 0x4f, 0x01, 0x4d, 0xde, 0x14, 0x76, 0x57, 0xe1, 0x8c, 0xf4,
	0x11, 0x95, 0x14, 0x23, 0xfb, 0x7f, 0xbf, 0xe3, 0xbc, 0x89, 0x7a, 0x99, 0x69, 0xf5, 0x26, 0x1a,
	0xcb, 0x9e, 0xa8, 0xb1, 0x1a, 0x14, 0xee, 0xd2, 0xee, 0x78, 0xdb, 0xb3, 0x87, 0x5f, 0x04, 0x4d,
	0xde, 0x88, 0x89, 0x23, 0xc6, 0x0d, 0x1c, 0xe0, 0xd3, 0x31, 0xe4, 0x8d, 0x98, 0xd1, 0xf8, 0x92,
	0x85, 0x39, 0xd9, 0x3d, 0xba, 0x0d, 0x6a, 0xfc, 0xe2, 0xd0, 0x82, 0x18, 0x4c, 0xea, 0x45, 0xea,
	0x8b, 0x47, 0x54, 0xe1, 0x63, 0xe9, 0xe5, 0xef, 0xe1, 0xdb, 0xcc, 0xbc, 0x79, 0xd6, 0x7e, 0x52,
	0x4f, 0xfe, 0x6d, 0x6b, 0x4a, 0x05, 0xdd, 0x02, 0x35, 0xde, 0x41, 0xc2, 0x4b, 0x2d, 0x33, 0xe1,
	0xa5, 0x17, 0x65, 0x2e, 0x70, 0x9e, 0x86, 0x52, 0x3c, 0xd4, 0x86, 0xec, 0x06, 0x66, 0x68, 0x5e,
	0xdc, 0x19, 0xbf, 0x5b, 0x1d, 0x4d, 0x4a, 0x82, 0x61, 0x72, 0xc6, 0x32, 0xd2, 0x27, 0x19, 0xf6,
	0xf3, 0x89, 0xf1, 0xbc, 0x40, 0x0f, 0x41, 0x8d, 0x27, 0x9a, 0xd8, 0x4b, 0xad, 0x24, 0xb1, 0x77,
	0x64, 0xec, 0x17, 0x39, 0x7a, 0x45, 0x9f, 0x82, 0x1e, 0x35, 0x7f, 0x1f, 0xd4, 0x78, 0xd6, 0x09,
	0x3d, 0xb5, 0xac, 0x84, 0x9e, 0x5e, 0x88, 0x34, 0x5e, 0x99, 0x42, 0xbf, 0xf6, 0x4e, 0x79, 0xd3,
	0xf4, 0xd1, 0x55, 0xb8, 0xd0, 0xbc, 0xb7, 0xfe, 0x77, 0xff, 0xf5, 0x9f, 0x9f, 0xdf, 0x9a, 0xed,
	0x96, 0x51, 0x35, 0x86, 0xbf, 0x3e, 0x0c, 0xbf, 0xbf, 0x3f, 0xfc, 0xf4, 0xf1, 0x70, 0x6f, 0x1f,
	0xe9, 0x49, 0xf0, 0xe0, 0xc7, 0xd7, 0x83, 0x57, 0x9f, 0x9d, 0xf5, 0xcd, 0xad, 0x9d, 0x41, 0x60,
	0x34, 0xdb, 0xad, 0x46, 0xae, 0x6e, 0xd5, 0xac, 0x5a, 0x45, 0x51, 0x1a, 0x45, 0x97, 0xd2, 0xc0,
	0xef, 0xf0, 0x8f, 0x97, 0xbd, 0x1b, 0x91, 0x70, 0xed, 0x98, 0xf2, 0xa0, 0xec, 0xf9, 0xcc, 0xda,
	0x21, 0x5e, 0xa7, 0x1b, 0x5a, 0x8c, 0x50, 0xdb, 0x7d, 0x86, 0x6d, 0xee, 0x3d, 0xb2, 0x3d, 0x42,
	0x06, 0x2c, 0xb1, 0xb9, 0xad, 0x72, 0x79, 0xf5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x8d,
	0xe2, 0x67, 0x7d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TemplateClient is the client API for Template service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TemplateClient interface {
	// 创建作业模板
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// 筛选作业模板
	Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error)
	// 获取作业模板
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// 更新作业模板
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// 删除作业模板
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type templateClient struct {
	cc *grpc.ClientConn
}

func NewTemplateClient(cc *grpc.ClientConn) TemplateClient {
	return &templateClient{cc}
}

func (c *templateClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/proto.Template/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error) {
	out := new(FilterResponse)
	err := c.cc.Invoke(ctx, "/proto.Template/Filter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/proto.Template/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/proto.Template/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/proto.Template/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServer is the server API for Template service.
type TemplateServer interface {
	// 创建作业模板
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// 筛选作业模板
	Filter(context.Context, *FilterRequest) (*FilterResponse, error)
	// 获取作业模板
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// 更新作业模板
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// 删除作业模板
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedTemplateServer can be embedded to have forward compatible implementations.
type UnimplementedTemplateServer struct {
}

func (*UnimplementedTemplateServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTemplateServer) Filter(ctx context.Context, req *FilterRequest) (*FilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Filter not implemented")
}
func (*UnimplementedTemplateServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTemplateServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTemplateServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterTemplateServer(s *grpc.Server, srv TemplateServer) {
	s.RegisterService(&_Template_serviceDesc, srv)
}

func _Template_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Template/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_Filter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).Filter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Template/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).Filter(ctx, req.(*FilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Template/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Template/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Template/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Template_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Template",
	HandlerType: (*TemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Template_Create_Handler,
		},
		{
			MethodName: "Filter",
			Handler:    _Template_Filter_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Template_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Template_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Template_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "template/template.proto",
}
