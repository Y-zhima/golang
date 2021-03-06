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

// 下载模板类型
type TemplateType int32

const (
	// 0-undefined,指除了关系链模板之外的导入模板，也可不传
	TemplateType_UNDEFINED TemplateType = 0
	// 1-导入主机业务拓扑
	TemplateType_HOSTCHAIN TemplateType = 1
	// 2-导入lake节点关系链
	TemplateType_LAKECHAIN TemplateType = 2
	// 3-导入交维表
	TemplateType_CROSSTABLE TemplateType = 3
)

var TemplateType_name = map[int32]string{
	0: "UNDEFINED",
	1: "HOSTCHAIN",
	2: "LAKECHAIN",
	3: "CROSSTABLE",
}

var TemplateType_value = map[string]int32{
	"UNDEFINED":  0,
	"HOSTCHAIN":  1,
	"LAKECHAIN":  2,
	"CROSSTABLE": 3,
}

func (x TemplateType) String() string {
	return proto.EnumName(TemplateType_name, int32(x))
}

func (TemplateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{0}
}

type DownloadTemplateRequest struct {
	// 导入类型
	TemplateType TemplateType `protobuf:"varint,1,opt,name=template_type,json=templateType,proto3,enum=file.TemplateType" json:"template_type,omitempty"`
	// 非关系链的导入模板的类型
	AssetType            string   `protobuf:"bytes,2,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadTemplateRequest) Reset()         { *m = DownloadTemplateRequest{} }
func (m *DownloadTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadTemplateRequest) ProtoMessage()    {}
func (*DownloadTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{0}
}

func (m *DownloadTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadTemplateRequest.Unmarshal(m, b)
}
func (m *DownloadTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadTemplateRequest.Marshal(b, m, deterministic)
}
func (m *DownloadTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadTemplateRequest.Merge(m, src)
}
func (m *DownloadTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadTemplateRequest.Size(m)
}
func (m *DownloadTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadTemplateRequest proto.InternalMessageInfo

func (m *DownloadTemplateRequest) GetTemplateType() TemplateType {
	if m != nil {
		return m.TemplateType
	}
	return TemplateType_UNDEFINED
}

func (m *DownloadTemplateRequest) GetAssetType() string {
	if m != nil {
		return m.AssetType
	}
	return ""
}

type DownloadTemplateResponse struct {
	// 返回的请求状态
	Status *common.ResponseStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 返回的excel文件
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadTemplateResponse) Reset()         { *m = DownloadTemplateResponse{} }
func (m *DownloadTemplateResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadTemplateResponse) ProtoMessage()    {}
func (*DownloadTemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{1}
}

func (m *DownloadTemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadTemplateResponse.Unmarshal(m, b)
}
func (m *DownloadTemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadTemplateResponse.Marshal(b, m, deterministic)
}
func (m *DownloadTemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadTemplateResponse.Merge(m, src)
}
func (m *DownloadTemplateResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadTemplateResponse.Size(m)
}
func (m *DownloadTemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadTemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadTemplateResponse proto.InternalMessageInfo

func (m *DownloadTemplateResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DownloadTemplateResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

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
	return fileDescriptor_ad806f8986a0c3f6, []int{2}
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
	return fileDescriptor_ad806f8986a0c3f6, []int{3}
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
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRequest) Reset()         { *m = UploadRequest{} }
func (m *UploadRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()    {}
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{4}
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

func (m *UploadRequest) GetKey() string {
	if m != nil {
		return m.Key
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
	return fileDescriptor_ad806f8986a0c3f6, []int{5}
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
	proto.RegisterEnum("file.TemplateType", TemplateType_name, TemplateType_value)
	proto.RegisterType((*DownloadTemplateRequest)(nil), "file.DownloadTemplateRequest")
	proto.RegisterType((*DownloadTemplateResponse)(nil), "file.DownloadTemplateResponse")
	proto.RegisterType((*UploadPlaybookRequest)(nil), "file.UploadPlaybookRequest")
	proto.RegisterType((*UploadPlaybookResponse)(nil), "file.UploadPlaybookResponse")
	proto.RegisterType((*UploadRequest)(nil), "file.UploadRequest")
	proto.RegisterType((*UploadResponse)(nil), "file.UploadResponse")
}

func init() { proto.RegisterFile("file/file.proto", fileDescriptor_ad806f8986a0c3f6) }

var fileDescriptor_ad806f8986a0c3f6 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0x12, 0x41,
	0x14, 0x76, 0x01, 0x51, 0x8e, 0x40, 0x37, 0x63, 0xa5, 0xeb, 0x96, 0x36, 0xb8, 0xf1, 0x82, 0xf4,
	0x62, 0x37, 0xc5, 0x18, 0x13, 0xef, 0xa0, 0xd0, 0xb4, 0x69, 0x03, 0x66, 0xa1, 0x37, 0xde, 0x98,
	0x01, 0xa6, 0x64, 0xc3, 0x32, 0xb3, 0x65, 0x67, 0xb5, 0x78, 0xe9, 0x2b, 0x78, 0xeb, 0x03, 0xf8,
	0x14, 0xbe, 0x84, 0xaf, 0xe0, 0x83, 0x98, 0xf9, 0xd9, 0xba, 0x60, 0x1b, 0x35, 0xf1, 0x86, 0xcc,
	0xf9, 0xbe, 0x73, 0xbe, 0xef, 0x7c, 0x30, 0x03, 0x6c, 0x5d, 0x06, 0x21, 0xf1, 0xc4, 0x87, 0x1b,
	0x2d, 0x19, 0x67, 0xa8, 0x20, 0xce, 0xb6, 0x39, 0x61, 0x8b, 0x05, 0xa3, 0x1e, 0x8e, 0x02, 0x85,
	0xdb, 0xf5, 0x19, 0x63, 0xb3, 0x90, 0x08, 0xc4, 0xc3, 0x94, 0x32, 0x8e, 0x79, 0xc0, 0x68, 0xac,
	0x58, 0xe7, 0x0a, 0x76, 0xba, 0xec, 0x03, 0x0d, 0x19, 0x9e, 0x8e, 0xc8, 0x22, 0x0a, 0x31, 0x27,
	0x3e, 0xb9, 0x4a, 0x48, 0xcc, 0xd1, 0x2b, 0xa8, 0x70, 0x0d, 0xbd, 0xe3, 0xab, 0x88, 0x58, 0x46,
	0xc3, 0x68, 0x56, 0x5b, 0xc8, 0x95, 0xa6, 0x69, 0xf7, 0x68, 0x15, 0x11, 0xbf, 0xcc, 0x33, 0x15,
	0xda, 0x03, 0xc0, 0x71, 0x4c, 0xb8, 0x9a, 0xca, 0x35, 0x8c, 0x66, 0xc9, 0x2f, 0x49, 0x44, 0xd0,
	0xce, 0x14, 0xac, 0xdf, 0x2d, 0xe3, 0x88, 0xd1, 0x98, 0x20, 0x17, 0x8a, 0x31, 0xc7, 0x3c, 0x89,
	0xa5, 0xd9, 0xa3, 0x56, 0xcd, 0x55, 0x79, 0xdc, 0xb4, 0x63, 0x28, 0x59, 0x5f, 0x77, 0x21, 0x0b,
	0x1e, 0x4c, 0x18, 0xe5, 0x84, 0x72, 0xe9, 0x53, 0xf6, 0xd3, 0xd2, 0x39, 0x84, 0x27, 0x17, 0x91,
	0xf0, 0x78, 0x13, 0xe2, 0xd5, 0x98, 0xb1, 0x79, 0x1a, 0x2b, 0x33, 0x62, 0xac, 0x8f, 0x7c, 0x35,
	0xa0, 0xb6, 0x39, 0xa3, 0xf7, 0x32, 0x21, 0x9f, 0x2c, 0x43, 0x39, 0x50, 0xf2, 0xc5, 0x51, 0x20,
	0x8b, 0xe9, 0x4b, 0x9d, 0x4e, 0x1c, 0x91, 0x0d, 0x0f, 0xc5, 0x37, 0x13, 0x07, 0x1f, 0x89, 0x95,
	0x97, 0xf0, 0x4d, 0x8d, 0xf6, 0x01, 0x08, 0xe5, 0xcb, 0x55, 0xc4, 0x02, 0xca, 0xad, 0x42, 0x23,
	0xdf, 0x2c, 0xf9, 0x19, 0x24, 0x93, 0xfb, 0xfe, 0xdf, 0xe4, 0x76, 0x86, 0x50, 0x51, 0x9b, 0xfe,
	0x31, 0x15, 0xaa, 0x41, 0x71, 0x9c, 0x4c, 0xe6, 0x84, 0xeb, 0x5d, 0x75, 0x25, 0x02, 0xcc, 0xc9,
	0x4a, 0x6f, 0x2a, 0x8e, 0xce, 0x17, 0x03, 0xaa, 0xa9, 0xea, 0x7f, 0xca, 0xad, 0x39, 0x8a, 0x17,
	0xc4, 0x2a, 0xfc, 0xe2, 0x44, 0xfd, 0xaf, 0x99, 0x0f, 0xce, 0xa0, 0x9c, 0xbd, 0x74, 0xa8, 0x02,
	0xa5, 0x8b, 0x7e, 0xb7, 0x77, 0x7c, 0xda, 0xef, 0x75, 0xcd, 0x7b, 0xa2, 0x3c, 0x19, 0x0c, 0x47,
	0x47, 0x27, 0xed, 0xd3, 0xbe, 0x69, 0x88, 0xf2, 0xbc, 0x7d, 0xd6, 0x53, 0x65, 0x0e, 0x55, 0x01,
	0x8e, 0xfc, 0xc1, 0x70, 0x38, 0x6a, 0x77, 0xce, 0x7b, 0x66, 0xbe, 0xf5, 0x2d, 0x07, 0x85, 0xe3,
	0x20, 0x24, 0x88, 0xa5, 0x99, 0xd3, 0xdf, 0x1c, 0xed, 0xaa, 0x0b, 0x7e, 0xeb, 0xed, 0xb1, 0xeb,
	0xb7, 0x93, 0x6a, 0x61, 0xc7, 0xf9, 0xf4, 0xfd, 0xc7, 0xe7, 0x5c, 0xdd, 0xd9, 0xf1, 0xde, 0x1f,
	0xca, 0xb7, 0xe9, 0x25, 0x6b, 0x8d, 0xaf, 0x8d, 0x03, 0x34, 0x80, 0xa2, 0x9a, 0x46, 0x8f, 0xb3,
	0x5a, 0xa9, 0xc1, 0xf6, 0x3a, 0xa8, 0x85, 0x6d, 0x29, 0xbc, 0xed, 0x6c, 0x6d, 0x08, 0x0b, 0x41,
	0x0e, 0xe6, 0xe6, 0x7b, 0x42, 0x7b, 0x4a, 0xe5, 0x8e, 0xa7, 0x6d, 0xef, 0xdf, 0x45, 0x6b, 0xbb,
	0x67, 0xd2, 0x6e, 0x17, 0x3d, 0xbd, 0xb1, 0x9b, 0x6e, 0xb4, 0x76, 0x3a, 0x60, 0x73, 0x16, 0xb9,
	0x97, 0x6c, 0x36, 0x99, 0x52, 0x17, 0x5f, 0xeb, 0xbf, 0xa1, 0x58, 0x2a, 0xbf, 0x7d, 0x3e, 0x0b,
	0x78, 0xca, 0x71, 0x16, 0x79, 0xf8, 0x9a, 0x78, 0x8a, 0xf7, 0x66, 0x8c, 0x25, 0x5c, 0x8a, 0x8e,
	0x8b, 0x12, 0x7a, 0xf1, 0x33, 0x00, 0x00, 0xff, 0xff, 0x01, 0x41, 0x32, 0x47, 0xcc, 0x04, 0x00,
	0x00,
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
	// 获取导入模板
	DownloadTemplate(ctx context.Context, in *DownloadTemplateRequest, opts ...grpc.CallOption) (*DownloadTemplateResponse, error)
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

func (c *fileClient) DownloadTemplate(ctx context.Context, in *DownloadTemplateRequest, opts ...grpc.CallOption) (*DownloadTemplateResponse, error) {
	out := new(DownloadTemplateResponse)
	err := c.cc.Invoke(ctx, "/file.File/DownloadTemplate", in, out, opts...)
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
	// 获取导入模板
	DownloadTemplate(context.Context, *DownloadTemplateRequest) (*DownloadTemplateResponse, error)
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
func (*UnimplementedFileServer) DownloadTemplate(ctx context.Context, req *DownloadTemplateRequest) (*DownloadTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadTemplate not implemented")
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

func _File_DownloadTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).DownloadTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.File/DownloadTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).DownloadTemplate(ctx, req.(*DownloadTemplateRequest))
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
		{
			MethodName: "DownloadTemplate",
			Handler:    _File_DownloadTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file/file.proto",
}
