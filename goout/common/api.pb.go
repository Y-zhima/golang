// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/api.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 请求返回状态码
type StatusCode int32

const (
	// 成功
	StatusCode_SUCCESS StatusCode = 0
	// 参数错误
	StatusCode_INVALID_ARGUMENT StatusCode = 400
	// 访问拒绝
	StatusCode_ACCESS_DENIED StatusCode = 403
)

var StatusCode_name = map[int32]string{
	0:   "SUCCESS",
	400: "INVALID_ARGUMENT",
	403: "ACCESS_DENIED",
}

var StatusCode_value = map[string]int32{
	"SUCCESS":          0,
	"INVALID_ARGUMENT": 400,
	"ACCESS_DENIED":    403,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}

func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71edda3932b293bf, []int{0}
}

// 分页
type Paging struct {
	// 总页数
	TotalPage int32 `protobuf:"varint,1,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
	// 当前页数
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// 每页显示的记录条数
	PerPage int32 `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	// 总记录数
	TotalRecord          int32    `protobuf:"varint,4,opt,name=total_record,json=totalRecord,proto3" json:"total_record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Paging) Reset()         { *m = Paging{} }
func (m *Paging) String() string { return proto.CompactTextString(m) }
func (*Paging) ProtoMessage()    {}
func (*Paging) Descriptor() ([]byte, []int) {
	return fileDescriptor_71edda3932b293bf, []int{0}
}

func (m *Paging) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paging.Unmarshal(m, b)
}
func (m *Paging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paging.Marshal(b, m, deterministic)
}
func (m *Paging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paging.Merge(m, src)
}
func (m *Paging) XXX_Size() int {
	return xxx_messageInfo_Paging.Size(m)
}
func (m *Paging) XXX_DiscardUnknown() {
	xxx_messageInfo_Paging.DiscardUnknown(m)
}

var xxx_messageInfo_Paging proto.InternalMessageInfo

func (m *Paging) GetTotalPage() int32 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

func (m *Paging) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Paging) GetPerPage() int32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *Paging) GetTotalRecord() int32 {
	if m != nil {
		return m.TotalRecord
	}
	return 0
}

// 请求返回状态
type ResponseStatus struct {
	// 状态码
	Code StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=common.StatusCode" json:"code,omitempty"`
	// 信息
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseStatus) Reset()         { *m = ResponseStatus{} }
func (m *ResponseStatus) String() string { return proto.CompactTextString(m) }
func (*ResponseStatus) ProtoMessage()    {}
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_71edda3932b293bf, []int{1}
}

func (m *ResponseStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseStatus.Unmarshal(m, b)
}
func (m *ResponseStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseStatus.Marshal(b, m, deterministic)
}
func (m *ResponseStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseStatus.Merge(m, src)
}
func (m *ResponseStatus) XXX_Size() int {
	return xxx_messageInfo_ResponseStatus.Size(m)
}
func (m *ResponseStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseStatus proto.InternalMessageInfo

func (m *ResponseStatus) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_SUCCESS
}

func (m *ResponseStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("common.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterType((*Paging)(nil), "common.Paging")
	proto.RegisterType((*ResponseStatus)(nil), "common.ResponseStatus")
}

func init() { proto.RegisterFile("common/api.proto", fileDescriptor_71edda3932b293bf) }

var fileDescriptor_71edda3932b293bf = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x14, 0x80, 0xff, 0xfd, 0x9b, 0x9b, 0x3b, 0xd3, 0x51, 0x0e, 0x08, 0x13, 0x14, 0x74, 0x17, 0x43,
	0xbc, 0x68, 0x41, 0x9f, 0x60, 0xae, 0x55, 0x06, 0x5a, 0x46, 0xea, 0xbc, 0xf0, 0xa6, 0xc4, 0x36,
	0x86, 0x82, 0xed, 0x09, 0x4d, 0x06, 0xf3, 0x2d, 0x04, 0x5f, 0x58, 0x38, 0xb1, 0x78, 0x97, 0x7c,
	0xdf, 0x97, 0xe4, 0x10, 0x08, 0x0a, 0xaa, 0x6b, 0x6a, 0x22, 0x69, 0xaa, 0xd0, 0xb4, 0xe4, 0x08,
	0x87, 0x9e, 0xcc, 0x3f, 0x61, 0xb8, 0x91, 0xba, 0x6a, 0x34, 0x9e, 0x03, 0x38, 0x72, 0xf2, 0x23,
	0x37, 0x52, 0xab, 0x59, 0xef, 0xa2, 0x77, 0x75, 0x20, 0xc6, 0x4c, 0x36, 0x52, 0x2b, 0x44, 0x18,
	0xb0, 0xf8, 0xcf, 0x82, 0xd7, 0x78, 0x0a, 0x87, 0x46, 0xb5, 0xfe, 0x40, 0x9f, 0xf9, 0xc8, 0xa8,
	0x96, 0xf3, 0x4b, 0x38, 0xf2, 0xb7, 0xb5, 0xaa, 0xa0, 0xb6, 0x9c, 0x0d, 0x58, 0x4f, 0x98, 0x09,
	0x46, 0x73, 0x01, 0x53, 0xa1, 0xac, 0xa1, 0xc6, 0xaa, 0xcc, 0x49, 0xb7, 0xb3, 0xb8, 0x80, 0x41,
	0x41, 0xa5, 0x7f, 0x7c, 0x7a, 0x83, 0xa1, 0x9f, 0x31, 0xf4, 0x76, 0x45, 0xa5, 0x12, 0xec, 0x71,
	0x06, 0xa3, 0x5a, 0x59, 0xdb, 0x8d, 0x33, 0x16, 0xdd, 0xf6, 0x3a, 0x06, 0xf8, 0xab, 0x71, 0x02,
	0xa3, 0x6c, 0xbb, 0x5a, 0x25, 0x59, 0x16, 0xfc, 0xc3, 0x13, 0x08, 0xd6, 0xe9, 0xcb, 0xf2, 0x71,
	0x1d, 0xe7, 0x4b, 0xf1, 0xb0, 0x7d, 0x4a, 0xd2, 0xe7, 0xe0, 0xab, 0x8f, 0x08, 0xc7, 0x4b, 0x4e,
	0xf2, 0x38, 0x49, 0xd7, 0x49, 0x1c, 0x7c, 0xf7, 0xef, 0xee, 0xe1, 0xcc, 0x91, 0x09, 0xdf, 0x49,
	0x17, 0x65, 0x13, 0xca, 0xbd, 0xf2, 0x9f, 0x66, 0x7f, 0x07, 0x7a, 0x5d, 0xe8, 0xca, 0x75, 0xd6,
	0x91, 0x89, 0xe4, 0x5e, 0x45, 0xbe, 0x88, 0x34, 0xd1, 0xce, 0x45, 0xbe, 0x7b, 0x1b, 0x32, 0xbc,
	0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x1a, 0x30, 0x3a, 0x7f, 0x01, 0x00, 0x00,
}
