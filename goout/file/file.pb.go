// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file/file.proto

package file

import (
	context "context"
	fmt "fmt"
	common "git.fogcdn.top/axe/protos/goout/common"
	proto "github.com/golang/protobuf/proto"
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

// 上传playbook压缩包并且解析入口yml文件的请求体
type UploadPlaybookRequest struct {
	// 上传的playbook压缩包
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadPlaybookRequest) Reset()         { *m = UploadPlaybookRequest{} }
func (m *UploadPlaybookRequest) String() string { return proto.CompactTextString(m) }
func (*UploadPlaybookRequest) ProtoMessage()    {}
func (*UploadPlaybookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{0}
}

func (m *UploadPlaybookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadPlaybookRequest.Unmarshal(m, b)
}
func (m *UploadPlaybookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadPlaybookRequest.Marshal(b, m, deterministic)
}
func (m *UploadPlaybookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadPlaybookRequest.Merge(m, src)
}
func (m *UploadPlaybookRequest) XXX_Size() int {
	return xxx_messageInfo_UploadPlaybookRequest.Size(m)
}
func (m *UploadPlaybookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadPlaybookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadPlaybookRequest proto.InternalMessageInfo

func (m *UploadPlaybookRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type UploadPlaybookResponse struct {
	// 解析出来的playbook项目存储在对象存储的url
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// 解析出来的playbook项目文件md5
	Md5 string `protobuf:"bytes,2,opt,name=md5,proto3" json:"md5,omitempty"`
	// 解析出来的playbook项目文件大小
	Filesize string `protobuf:"bytes,3,opt,name=filesize,proto3" json:"filesize,omitempty"`
	// 解析出来的playbook项目入口yml文件，有多个
	Entrypoint []string `protobuf:"bytes,4,rep,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	// 返回的请求状态
	Status               *common.ResponseStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UploadPlaybookResponse) Reset()         { *m = UploadPlaybookResponse{} }
func (m *UploadPlaybookResponse) String() string { return proto.CompactTextString(m) }
func (*UploadPlaybookResponse) ProtoMessage()    {}
func (*UploadPlaybookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{1}
}

func (m *UploadPlaybookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadPlaybookResponse.Unmarshal(m, b)
}
func (m *UploadPlaybookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadPlaybookResponse.Marshal(b, m, deterministic)
}
func (m *UploadPlaybookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadPlaybookResponse.Merge(m, src)
}
func (m *UploadPlaybookResponse) XXX_Size() int {
	return xxx_messageInfo_UploadPlaybookResponse.Size(m)
}
func (m *UploadPlaybookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadPlaybookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadPlaybookResponse proto.InternalMessageInfo

func (m *UploadPlaybookResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *UploadPlaybookResponse) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *UploadPlaybookResponse) GetFilesize() string {
	if m != nil {
		return m.Filesize
	}
	return ""
}

func (m *UploadPlaybookResponse) GetEntrypoint() []string {
	if m != nil {
		return m.Entrypoint
	}
	return nil
}

func (m *UploadPlaybookResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// 上传csv等通用文件
type UploadRequest struct {
	// 文件字段
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// 指定上传bucket_name（可选）
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// 指定上传存储路径（可选）
	Job                  string   `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRequest) Reset()         { *m = UploadRequest{} }
func (m *UploadRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()    {}
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{2}
}

func (m *UploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRequest.Unmarshal(m, b)
}
func (m *UploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRequest.Marshal(b, m, deterministic)
}
func (m *UploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRequest.Merge(m, src)
}
func (m *UploadRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRequest.Size(m)
}
func (m *UploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRequest proto.InternalMessageInfo

func (m *UploadRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *UploadRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *UploadRequest) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

type UploadResponse struct {
	// 通用文件对象存储的url
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// 解析出来的通用文件md5
	Md5 string `protobuf:"bytes,2,opt,name=md5,proto3" json:"md5,omitempty"`
	// 解析出来的通用文件大小
	Filesize string `protobuf:"bytes,3,opt,name=filesize,proto3" json:"filesize,omitempty"`
	// 解析出来的通用文件name
	Filename string `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	// 返回的请求状态
	Status               *common.ResponseStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{3}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

func (m *UploadResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *UploadResponse) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *UploadResponse) GetFilesize() string {
	if m != nil {
		return m.Filesize
	}
	return ""
}

func (m *UploadResponse) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *UploadResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadPlaybookRequest)(nil), "file.UploadPlaybookRequest")
	proto.RegisterType((*UploadPlaybookResponse)(nil), "file.UploadPlaybookResponse")
	proto.RegisterType((*UploadRequest)(nil), "file.UploadRequest")
	proto.RegisterType((*UploadResponse)(nil), "file.UploadResponse")
}

func init() { proto.RegisterFile("file/file.proto", fileDescriptor_ad806f8986a0c3f6) }

var fileDescriptor_ad806f8986a0c3f6 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xbd, 0xce, 0xd3, 0x30,
	0x14, 0x95, 0xbf, 0x86, 0xc0, 0x67, 0x7e, 0x5a, 0x99, 0x52, 0xa2, 0x50, 0xa1, 0xca, 0x62, 0xa8,
	0x3a, 0xc4, 0x6a, 0x51, 0x97, 0x8e, 0x1d, 0x58, 0x41, 0xa9, 0x58, 0xd8, 0x9c, 0xc4, 0x8d, 0x42,
	0x13, 0xdf, 0x50, 0x3b, 0xa8, 0x65, 0xe4, 0x15, 0x58, 0x79, 0x00, 0x1e, 0x85, 0x9d, 0x57, 0xe0,
	0x41, 0x90, 0xed, 0x04, 0xda, 0xaa, 0x12, 0x42, 0x62, 0xb1, 0xee, 0x3d, 0xc7, 0xf7, 0xf8, 0x1c,
	0xf9, 0xe2, 0xfe, 0xb6, 0x28, 0x05, 0x33, 0x47, 0x54, 0xef, 0x41, 0x03, 0xf1, 0x4c, 0x1d, 0x0e,
	0x52, 0xa8, 0x2a, 0x90, 0x8c, 0xd7, 0x85, 0xc3, 0xc3, 0x71, 0x0e, 0x90, 0x97, 0xc2, 0x20, 0x8c,
	0x4b, 0x09, 0x9a, 0xeb, 0x02, 0xa4, 0x72, 0x2c, 0x9d, 0xe3, 0x27, 0x6f, 0xeb, 0x12, 0x78, 0xf6,
	0xa6, 0xe4, 0xc7, 0x04, 0x60, 0x17, 0x8b, 0x0f, 0x8d, 0x50, 0x9a, 0x04, 0xf8, 0x6e, 0x0a, 0x52,
	0x0b, 0xa9, 0x03, 0x34, 0x41, 0xd3, 0x07, 0x71, 0xd7, 0xd2, 0x6f, 0x08, 0x8f, 0x2e, 0x67, 0x54,
	0x0d, 0x52, 0x09, 0x32, 0xc0, 0xbd, 0x66, 0x5f, 0xda, 0x81, 0xdb, 0xd8, 0x94, 0x06, 0xa9, 0xb2,
	0x65, 0x70, 0xe3, 0x90, 0x2a, 0x5b, 0x92, 0x10, 0xdf, 0x33, 0x4e, 0x55, 0xf1, 0x49, 0x04, 0x3d,
	0x0b, 0xff, 0xee, 0xc9, 0x73, 0x8c, 0x85, 0xd4, 0xfb, 0x63, 0x0d, 0x85, 0xd4, 0x81, 0x37, 0xe9,
	0x4d, 0x6f, 0xe3, 0x13, 0x84, 0x44, 0xd8, 0x57, 0x9a, 0xeb, 0x46, 0x05, 0x77, 0x26, 0x68, 0x7a,
	0x7f, 0x31, 0x8a, 0x5c, 0xdc, 0xa8, 0x73, 0xb0, 0xb1, 0x6c, 0xdc, 0xde, 0xa2, 0x1b, 0xfc, 0xd0,
	0x39, 0xfd, 0x6b, 0x2a, 0x32, 0xc2, 0x7e, 0xd2, 0xa4, 0x3b, 0xa1, 0x5b, 0xaf, 0x6d, 0x67, 0x02,
	0xbc, 0x87, 0xa4, 0x75, 0x6a, 0x4a, 0xfa, 0x15, 0xe1, 0x47, 0x9d, 0xea, 0x7f, 0xca, 0xdd, 0x72,
	0x92, 0x57, 0x22, 0xf0, 0xfe, 0x70, 0xa6, 0xff, 0xd7, 0xcc, 0x8b, 0xef, 0x08, 0x7b, 0xaf, 0x8a,
	0x52, 0x10, 0xe8, 0x6c, 0x76, 0xdf, 0x44, 0x9e, 0x45, 0x76, 0x5f, 0xae, 0x7e, 0x78, 0x38, 0xbe,
	0x4e, 0xba, 0x37, 0x28, 0xfd, 0xfc, 0xe3, 0xe7, 0x97, 0x9b, 0x31, 0x7d, 0xca, 0x3e, 0xce, 0xed,
	0xd6, 0xb1, 0xe6, 0xec, 0xe2, 0x0a, 0xcd, 0xc8, 0x6b, 0xec, 0xbb, 0x69, 0xf2, 0xf8, 0x54, 0xab,
	0x7b, 0x60, 0x78, 0x0e, 0xb6, 0xc2, 0xa1, 0x15, 0x1e, 0xd2, 0xfe, 0x85, 0xf0, 0x0a, 0xcd, 0xd6,
	0x6b, 0x1c, 0x6a, 0xa8, 0xa3, 0x2d, 0xe4, 0x69, 0x26, 0x23, 0x7e, 0x68, 0x57, 0x5d, 0x59, 0xa1,
	0x77, 0x2f, 0xf2, 0x42, 0x77, 0x9c, 0x86, 0x9a, 0xf1, 0x83, 0x60, 0x8e, 0x67, 0x39, 0x40, 0xa3,
	0xad, 0x5a, 0xe2, 0x5b, 0xe8, 0xe5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0xde, 0x48, 0x7b,
	0x30, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileClient interface {
	// 上传playbook压缩包并且解析入口yml文件
	UploadPlaybook(ctx context.Context, in *UploadPlaybookRequest, opts ...grpc.CallOption) (*UploadPlaybookResponse, error)
	// 上传csv等通用文件
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
}

type fileClient struct {
	cc *grpc.ClientConn
}

func NewFileClient(cc *grpc.ClientConn) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) UploadPlaybook(ctx context.Context, in *UploadPlaybookRequest, opts ...grpc.CallOption) (*UploadPlaybookResponse, error) {
	out := new(UploadPlaybookResponse)
	err := c.cc.Invoke(ctx, "/file.File/UploadPlaybook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, "/file.File/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
type FileServer interface {
	// 上传playbook压缩包并且解析入口yml文件
	UploadPlaybook(context.Context, *UploadPlaybookRequest) (*UploadPlaybookResponse, error)
	// 上传csv等通用文件
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
}

// UnimplementedFileServer can be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (*UnimplementedFileServer) UploadPlaybook(ctx context.Context, req *UploadPlaybookRequest) (*UploadPlaybookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPlaybook not implemented")
}
func (*UnimplementedFileServer) Upload(ctx context.Context, req *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterFileServer(s *grpc.Server, srv FileServer) {
	s.RegisterService(&_File_serviceDesc, srv)
}

func _File_UploadPlaybook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPlaybookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).UploadPlaybook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.File/UploadPlaybook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).UploadPlaybook(ctx, req.(*UploadPlaybookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.File/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _File_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file.File",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadPlaybook",
			Handler:    _File_UploadPlaybook_Handler,
		},
		{
			MethodName: "Upload",
			Handler:    _File_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file/file.proto",
}
