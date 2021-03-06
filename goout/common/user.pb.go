// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/user.proto

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

type User struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ea6bce22462b6ad, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "common.User")
}

func init() { proto.RegisterFile("common/user.proto", fileDescriptor_9ea6bce22462b6ad) }

var fileDescriptor_9ea6bce22462b6ad = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83,
	0x08, 0x29, 0xd9, 0x70, 0xb1, 0x84, 0x16, 0xa7, 0x16, 0x09, 0x89, 0x73, 0xb1, 0x83, 0x64, 0xe3,
	0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xd8, 0x40, 0x5c, 0xcf, 0x14, 0x21, 0x69,
	0x2e, 0x4e, 0xb0, 0x44, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x07,
	0x48, 0xc0, 0x2f, 0x31, 0x37, 0xd5, 0xc9, 0x8d, 0x4b, 0xa6, 0x24, 0xbf, 0x40, 0x2f, 0x2d, 0x3f,
	0x3d, 0x39, 0x25, 0x4f, 0x2f, 0xb1, 0x22, 0x15, 0x62, 0x7a, 0xb1, 0x1e, 0xc4, 0xf4, 0x28, 0xb5,
	0xf4, 0xcc, 0x12, 0x98, 0x6c, 0x49, 0x7e, 0x81, 0x7e, 0x62, 0x45, 0xaa, 0x3e, 0x44, 0x85, 0x7e,
	0x7a, 0x7e, 0x7e, 0x69, 0x89, 0x3e, 0x44, 0x5d, 0x12, 0x1b, 0x58, 0xd0, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x34, 0x28, 0x58, 0x2f, 0xa9, 0x00, 0x00, 0x00,
}
