// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file/file.proto

package file

import (
	context "context"
	fmt "fmt"
	common "git.fogcdn.top/axe/protos/goout/common"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type UploadPlaybookRequest struct {
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
	Url                  string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Md5                  string                 `protobuf:"bytes,2,opt,name=md5,proto3" json:"md5,omitempty"`
	Filesize             string                 `protobuf:"bytes,3,opt,name=filesize,proto3" json:"filesize,omitempty"`
	Entrypoint           []string               `protobuf:"bytes,4,rep,name=entrypoint,proto3" json:"entrypoint,omitempty"`
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

func init() {
	proto.RegisterType((*UploadPlaybookRequest)(nil), "file.UploadPlaybookRequest")
	proto.RegisterType((*UploadPlaybookResponse)(nil), "file.UploadPlaybookResponse")
}

func init() { proto.RegisterFile("file/file.proto", fileDescriptor_ad806f8986a0c3f6) }

var fileDescriptor_ad806f8986a0c3f6 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0xb6, 0x56, 0x1b, 0x45, 0x4b, 0xc0, 0xba, 0xac, 0x45, 0xca, 0x22, 0xb2, 0x88,
	0x4d, 0x68, 0xa5, 0x17, 0x8f, 0x1e, 0x3c, 0xcb, 0x8a, 0x17, 0x6f, 0xe9, 0x6e, 0x1a, 0x82, 0xbb,
	0x99, 0x74, 0x93, 0x55, 0xeb, 0xd1, 0x57, 0xf0, 0x25, 0x7c, 0x1f, 0x5f, 0xc1, 0x07, 0x91, 0xcd,
	0xb6, 0x62, 0xb5, 0x97, 0x64, 0xe6, 0x9b, 0x7f, 0xc8, 0xcc, 0x1f, 0x74, 0x30, 0x93, 0x19, 0xa7,
	0xd5, 0x41, 0x74, 0x01, 0x16, 0x70, 0xab, 0x8a, 0x83, 0x6e, 0x02, 0x79, 0x0e, 0x8a, 0x32, 0x2d,
	0x6b, 0x1e, 0xf4, 0x05, 0x80, 0xc8, 0x78, 0x45, 0x28, 0x53, 0x0a, 0x2c, 0xb3, 0x12, 0x94, 0x59,
	0x56, 0x2f, 0xdc, 0x95, 0x0c, 0x05, 0x57, 0x43, 0xf3, 0xcc, 0x84, 0xe0, 0x05, 0x05, 0xed, 0x14,
	0xff, 0xd5, 0xe1, 0x08, 0x1d, 0xde, 0xeb, 0x0c, 0x58, 0x7a, 0x9b, 0xb1, 0xc5, 0x14, 0xe0, 0x31,
	0xe6, 0xf3, 0x92, 0x1b, 0x8b, 0x7d, 0xb4, 0x9d, 0x80, 0xb2, 0x5c, 0x59, 0xdf, 0x1b, 0x78, 0xd1,
	0x5e, 0xbc, 0x4a, 0xc3, 0x0f, 0x0f, 0xf5, 0xfe, 0xf6, 0x18, 0x0d, 0xca, 0x70, 0xdc, 0x45, 0xcd,
	0xb2, 0xc8, 0x5c, 0x43, 0x27, 0xae, 0xc2, 0x8a, 0xe4, 0xe9, 0xc4, 0x6f, 0xd4, 0x24, 0x4f, 0x27,
	0x38, 0x40, 0x3b, 0xd5, 0x5e, 0x46, 0xbe, 0x72, 0xbf, 0xe9, 0xf0, 0x4f, 0x8e, 0x4f, 0x10, 0xe2,
	0xca, 0x16, 0x0b, 0x0d, 0x52, 0x59, 0xbf, 0x35, 0x68, 0x46, 0x9d, 0xf8, 0x17, 0xc1, 0x04, 0xb5,
	0x8d, 0x65, 0xb6, 0x34, 0xfe, 0xd6, 0xc0, 0x8b, 0x76, 0xc7, 0x3d, 0x52, 0x9b, 0x43, 0x56, 0x13,
	0xdc, 0xb9, 0x6a, 0xbc, 0x54, 0x8d, 0x17, 0xa8, 0x75, 0x23, 0x33, 0x8e, 0xe7, 0x68, 0x7f, 0x7d,
	0x62, 0x7c, 0x4c, 0x9c, 0xd1, 0x1b, 0x77, 0x0f, 0xfa, 0x9b, 0x8b, 0xf5, 0x13, 0x61, 0xf8, 0xf6,
	0xf9, 0xf5, 0xde, 0xe8, 0x87, 0x47, 0xf4, 0x69, 0xe4, 0xbe, 0x8b, 0x96, 0x6b, 0xc2, 0x2b, 0xef,
	0x3c, 0xf2, 0xae, 0xcf, 0x1e, 0x4e, 0x85, 0xb4, 0x64, 0x06, 0x22, 0x49, 0x15, 0xb1, 0xa0, 0x29,
	0x7b, 0xe1, 0xd4, 0x99, 0x6e, 0xa8, 0x00, 0x28, 0xad, 0xeb, 0x9d, 0xb6, 0x1d, 0xba, 0xfc, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xd7, 0x8f, 0x0e, 0x79, 0xfe, 0x01, 0x00, 0x00,
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
	UploadPlaybook(ctx context.Context, opts ...grpc.CallOption) (File_UploadPlaybookClient, error)
}

type fileClient struct {
	cc *grpc.ClientConn
}

func NewFileClient(cc *grpc.ClientConn) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) UploadPlaybook(ctx context.Context, opts ...grpc.CallOption) (File_UploadPlaybookClient, error) {
	stream, err := c.cc.NewStream(ctx, &_File_serviceDesc.Streams[0], "/file.File/UploadPlaybook", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileUploadPlaybookClient{stream}
	return x, nil
}

type File_UploadPlaybookClient interface {
	Send(*UploadPlaybookRequest) error
	CloseAndRecv() (*UploadPlaybookResponse, error)
	grpc.ClientStream
}

type fileUploadPlaybookClient struct {
	grpc.ClientStream
}

func (x *fileUploadPlaybookClient) Send(m *UploadPlaybookRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileUploadPlaybookClient) CloseAndRecv() (*UploadPlaybookResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadPlaybookResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServer is the server API for File service.
type FileServer interface {
	// 上传playbook压缩包并且解析入口yml文件
	UploadPlaybook(File_UploadPlaybookServer) error
}

func RegisterFileServer(s *grpc.Server, srv FileServer) {
	s.RegisterService(&_File_serviceDesc, srv)
}

func _File_UploadPlaybook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServer).UploadPlaybook(&fileUploadPlaybookServer{stream})
}

type File_UploadPlaybookServer interface {
	SendAndClose(*UploadPlaybookResponse) error
	Recv() (*UploadPlaybookRequest, error)
	grpc.ServerStream
}

type fileUploadPlaybookServer struct {
	grpc.ServerStream
}

func (x *fileUploadPlaybookServer) SendAndClose(m *UploadPlaybookResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileUploadPlaybookServer) Recv() (*UploadPlaybookRequest, error) {
	m := new(UploadPlaybookRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _File_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file.File",
	HandlerType: (*FileServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadPlaybook",
			Handler:       _File_UploadPlaybook_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "file/file.proto",
}
