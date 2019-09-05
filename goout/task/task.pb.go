// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task/task.proto

package task

import (
	context "context"
	fmt "fmt"
	cmdb "git.fogcdn.top/axe/protos/goout/cmdb"
	common "git.fogcdn.top/axe/protos/goout/common"
	schedule "git.fogcdn.top/axe/protos/goout/schedule"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
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

// 定时任务类型 0-undefined 1-单次任务 2-定时任务
type ScheduleType int32

const (
	// 0-undefined
	ScheduleType_UNDEFINED ScheduleType = 0
	// 1-单次任务
	ScheduleType_SINGLE ScheduleType = 1
	// 2-定时任务
	ScheduleType_CRONTAB ScheduleType = 2
)

var ScheduleType_name = map[int32]string{
	0: "UNDEFINED",
	1: "SINGLE",
	2: "CRONTAB",
}

var ScheduleType_value = map[string]int32{
	"UNDEFINED": 0,
	"SINGLE":    1,
	"CRONTAB":   2,
}

func (x ScheduleType) String() string {
	return proto.EnumName(ScheduleType_name, int32(x))
}

func (ScheduleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{0}
}

// 作业任务实例
type TaskObject struct {
	// 作业任务ID
	TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// 模板ID
	TemplateId int32 `protobuf:"varint,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// 定时任务ID
	ScheduleId int32 `protobuf:"varint,3,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	// 定时任务类型
	ScheduleType ScheduleType `protobuf:"varint,4,opt,name=schedule_type,json=scheduleType,proto3,enum=task.ScheduleType" json:"schedule_type,omitempty"`
	// 任务类型
	TaskType schedule.TaskType `protobuf:"varint,5,opt,name=task_type,json=taskType,proto3,enum=schedule.TaskType" json:"task_type,omitempty"`
	// cmdb的搜索条件
	CmdbSearchRequest []*cmdb.ChooseHostRequest `protobuf:"bytes,6,rep,name=cmdb_search_request,json=cmdbSearchRequest,proto3" json:"cmdb_search_request,omitempty"`
	// 额外变量JSON String 例如： {"key":"testKey","value":"testVal","description":"测试描述"}
	ExtraVar string `protobuf:"bytes,7,opt,name=extra_var,json=extraVar,proto3" json:"extra_var,omitempty"`
	// 执行人ID
	Executor int32 `protobuf:"varint,8,opt,name=executor,proto3" json:"executor,omitempty"`
	// 总共执行多少主机
	ExecuteCount int32 `protobuf:"varint,9,opt,name=execute_count,json=executeCount,proto3" json:"execute_count,omitempty"`
	// 执行失败多少台主机
	FailCount int32 `protobuf:"varint,10,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
	// 执行成功多少台主机
	SuccessCount int32 `protobuf:"varint,11,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	// 执行开始时间
	StartTime string `protobuf:"bytes,12,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// 执行结束时间
	EndTime string `protobuf:"bytes,13,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// 作业名称
	Name                 string   `protobuf:"bytes,14,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskObject) Reset()         { *m = TaskObject{} }
func (m *TaskObject) String() string { return proto.CompactTextString(m) }
func (*TaskObject) ProtoMessage()    {}
func (*TaskObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{0}
}

func (m *TaskObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskObject.Unmarshal(m, b)
}
func (m *TaskObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskObject.Marshal(b, m, deterministic)
}
func (m *TaskObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskObject.Merge(m, src)
}
func (m *TaskObject) XXX_Size() int {
	return xxx_messageInfo_TaskObject.Size(m)
}
func (m *TaskObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskObject.DiscardUnknown(m)
}

var xxx_messageInfo_TaskObject proto.InternalMessageInfo

func (m *TaskObject) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *TaskObject) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func (m *TaskObject) GetScheduleId() int32 {
	if m != nil {
		return m.ScheduleId
	}
	return 0
}

func (m *TaskObject) GetScheduleType() ScheduleType {
	if m != nil {
		return m.ScheduleType
	}
	return ScheduleType_UNDEFINED
}

func (m *TaskObject) GetTaskType() schedule.TaskType {
	if m != nil {
		return m.TaskType
	}
	return schedule.TaskType_UNDEFINED
}

func (m *TaskObject) GetCmdbSearchRequest() []*cmdb.ChooseHostRequest {
	if m != nil {
		return m.CmdbSearchRequest
	}
	return nil
}

func (m *TaskObject) GetExtraVar() string {
	if m != nil {
		return m.ExtraVar
	}
	return ""
}

func (m *TaskObject) GetExecutor() int32 {
	if m != nil {
		return m.Executor
	}
	return 0
}

func (m *TaskObject) GetExecuteCount() int32 {
	if m != nil {
		return m.ExecuteCount
	}
	return 0
}

func (m *TaskObject) GetFailCount() int32 {
	if m != nil {
		return m.FailCount
	}
	return 0
}

func (m *TaskObject) GetSuccessCount() int32 {
	if m != nil {
		return m.SuccessCount
	}
	return 0
}

func (m *TaskObject) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *TaskObject) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *TaskObject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 创建任务实例请求
type CreateRequest struct {
	// 模板ID
	TemplateId int32 `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// cmdb的搜索条件
	CmdbSearchRequest []*cmdb.ChooseHostRequest `protobuf:"bytes,2,rep,name=cmdb_search_request,json=cmdbSearchRequest,proto3" json:"cmdb_search_request,omitempty"`
	// 任务类型
	TaskType schedule.TaskType `protobuf:"varint,3,opt,name=task_type,json=taskType,proto3,enum=schedule.TaskType" json:"task_type,omitempty"`
	// 额外变量JSON String 例如： {"key":"testKey","value":"testVal","description":"测试描述"}
	ExtraVar string `protobuf:"bytes,4,opt,name=extra_var,json=extraVar,proto3" json:"extra_var,omitempty"`
	// 定时任务ID
	ScheduleId int32 `protobuf:"varint,5,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	// 定时任务类型
	ScheduleType ScheduleType `protobuf:"varint,6,opt,name=schedule_type,json=scheduleType,proto3,enum=task.ScheduleType" json:"schedule_type,omitempty"`
	// 作业名称
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{1}
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

func (m *CreateRequest) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func (m *CreateRequest) GetCmdbSearchRequest() []*cmdb.ChooseHostRequest {
	if m != nil {
		return m.CmdbSearchRequest
	}
	return nil
}

func (m *CreateRequest) GetTaskType() schedule.TaskType {
	if m != nil {
		return m.TaskType
	}
	return schedule.TaskType_UNDEFINED
}

func (m *CreateRequest) GetExtraVar() string {
	if m != nil {
		return m.ExtraVar
	}
	return ""
}

func (m *CreateRequest) GetScheduleId() int32 {
	if m != nil {
		return m.ScheduleId
	}
	return 0
}

func (m *CreateRequest) GetScheduleType() ScheduleType {
	if m != nil {
		return m.ScheduleType
	}
	return ScheduleType_UNDEFINED
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 创建任务实例请求返回
type CreateResponse struct {
	// 任务实例ID
	TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// 返回的请求状态
	Status               *common.ResponseStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{2}
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

func (m *CreateResponse) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *CreateResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// 获取作业任务请求
type GetRequest struct {
	// 作业任务ID
	TaskId               int64    `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{3}
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

func (m *GetRequest) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

// 获取作业任务请求返回
type GetResponse struct {
	// 获取的作业任务实例
	Task *TaskObject `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	// task对于的作业版本
	PlaybookVersion string `protobuf:"bytes,2,opt,name=playbook_version,json=playbookVersion,proto3" json:"playbook_version,omitempty"`
	// task对于的主机模块
	HostModule string `protobuf:"bytes,3,opt,name=host_module,json=hostModule,proto3" json:"host_module,omitempty"`
	// 返回的请求状态
	Status               *common.ResponseStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{4}
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

func (m *GetResponse) GetTask() *TaskObject {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *GetResponse) GetPlaybookVersion() string {
	if m != nil {
		return m.PlaybookVersion
	}
	return ""
}

func (m *GetResponse) GetHostModule() string {
	if m != nil {
		return m.HostModule
	}
	return ""
}

func (m *GetResponse) GetStatus() *common.ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// 获取作业任务详细执行过程请求
type GetLogRequest struct {
	// 作业任务ID
	TaskId               int64    `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLogRequest) Reset()         { *m = GetLogRequest{} }
func (m *GetLogRequest) String() string { return proto.CompactTextString(m) }
func (*GetLogRequest) ProtoMessage()    {}
func (*GetLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{5}
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

func (m *GetLogRequest) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

// 获取作业任务详细执行过程请求返回
type GetLogResponse struct {
	// 返回的请求状态
	Status               *common.ResponseStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetLogResponse) Reset()         { *m = GetLogResponse{} }
func (m *GetLogResponse) String() string { return proto.CompactTextString(m) }
func (*GetLogResponse) ProtoMessage()    {}
func (*GetLogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{6}
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

// 筛选获取作业任务请求
type FilterRequest struct {
	// 分页信息
	Paging *common.Paging `protobuf:"bytes,1,opt,name=paging,proto3" json:"paging,omitempty"`
	// 用于筛选任务名字的关键字
	TaskName             string   `protobuf:"bytes,2,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterRequest) Reset()         { *m = FilterRequest{} }
func (m *FilterRequest) String() string { return proto.CompactTextString(m) }
func (*FilterRequest) ProtoMessage()    {}
func (*FilterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{7}
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

func (m *FilterRequest) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *FilterRequest) GetTaskType() int64 {
	if m != nil {
		return m.TaskType
	}
	return 0
}

// 筛选作业任务请求返回
type FilterResponse struct {
	// 筛选到的多个作业任务实例
	Tasks []*TaskObject `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	// 分页信息
	Paging *common.Paging `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
	// 返回的请求状态
	Status               *common.ResponseStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FilterResponse) Reset()         { *m = FilterResponse{} }
func (m *FilterResponse) String() string { return proto.CompactTextString(m) }
func (*FilterResponse) ProtoMessage()    {}
func (*FilterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f2b86464a95fe, []int{8}
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

func (m *FilterResponse) GetTasks() []*TaskObject {
	if m != nil {
		return m.Tasks
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

func init() {
	proto.RegisterEnum("task.ScheduleType", ScheduleType_name, ScheduleType_value)
	proto.RegisterType((*TaskObject)(nil), "task.TaskObject")
	proto.RegisterType((*CreateRequest)(nil), "task.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "task.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "task.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "task.GetResponse")
	proto.RegisterType((*GetLogRequest)(nil), "task.GetLogRequest")
	proto.RegisterType((*GetLogResponse)(nil), "task.GetLogResponse")
	proto.RegisterType((*FilterRequest)(nil), "task.FilterRequest")
	proto.RegisterType((*FilterResponse)(nil), "task.FilterResponse")
}

func init() { proto.RegisterFile("task/task.proto", fileDescriptor_8e8f2b86464a95fe) }

var fileDescriptor_8e8f2b86464a95fe = []byte{
	// 1042 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5b, 0x6f, 0x1b, 0x45,
	0x18, 0xcd, 0xda, 0x8e, 0x13, 0x7f, 0x8e, 0x1d, 0x67, 0x1a, 0x35, 0x8b, 0x29, 0x24, 0xda, 0x84,
	0xd6, 0x44, 0xc4, 0x2b, 0x82, 0x28, 0x52, 0x25, 0x24, 0x9a, 0x4b, 0x43, 0x50, 0x48, 0xd0, 0x26,
	0x54, 0xc0, 0x8b, 0x35, 0x59, 0x4f, 0x36, 0x4b, 0xbc, 0x3b, 0xdb, 0x9d, 0x71, 0x2e, 0x42, 0xbc,
	0x00, 0x4f, 0x48, 0x48, 0x15, 0xbc, 0x21, 0xfa, 0x50, 0x24, 0xa4, 0xf2, 0xc4, 0x0b, 0x42, 0x55,
	0x11, 0xe2, 0x3f, 0x94, 0x1f, 0x50, 0xa9, 0x8a, 0x12, 0x7e, 0x06, 0x9a, 0x9b, 0x2f, 0x25, 0x6d,
	0xd3, 0x17, 0x7b, 0xf6, 0x9c, 0x33, 0xdf, 0x7c, 0x97, 0xb3, 0xb3, 0x30, 0xca, 0x31, 0xdb, 0x73,
	0xc5, 0x4f, 0x3d, 0x49, 0x29, 0xa7, 0x28, 0x27, 0xd6, 0xd5, 0x4b, 0x01, 0xa5, 0x41, 0x8b, 0xb8,
	0x38, 0x09, 0x5d, 0x1c, 0xc7, 0x94, 0x63, 0x1e, 0xd2, 0x98, 0x29, 0x4d, 0xf5, 0x0d, 0xf9, 0xe7,
	0xcf, 0x05, 0x24, 0x9e, 0x63, 0x07, 0x38, 0x08, 0x48, 0xea, 0xd2, 0x44, 0x2a, 0xce, 0x50, 0x5f,
	0x0d, 0x42, 0xbe, 0xdb, 0xde, 0xae, 0xfb, 0x34, 0x72, 0xa3, 0x83, 0x90, 0xef, 0xd1, 0x03, 0x37,
	0xa0, 0x73, 0x92, 0x9c, 0xdb, 0xc7, 0xad, 0xb0, 0x89, 0x39, 0x4d, 0x99, 0xdb, 0x59, 0xea, 0x7d,
	0x15, 0x9f, 0x46, 0x11, 0x8d, 0x45, 0x0e, 0x1a, 0x99, 0x60, 0xfe, 0x2e, 0x69, 0xb6, 0x5b, 0xc4,
	0x35, 0x0b, 0x4d, 0x8c, 0xfa, 0x51, 0x73, 0xdb, 0x15, 0x3f, 0x0a, 0x70, 0xbe, 0xcd, 0x01, 0x6c,
	0x61, 0xb6, 0xb7, 0xb1, 0xfd, 0x39, 0xf1, 0x39, 0x9a, 0x80, 0x21, 0x51, 0x56, 0x23, 0x6c, 0xda,
	0xd6, 0x94, 0x55, 0xcb, 0x7a, 0x79, 0xf1, 0xb8, 0xda, 0x44, 0x93, 0x50, 0xe4, 0x24, 0x4a, 0x5a,
	0x98, 0x13, 0x41, 0x66, 0xa6, 0xac, 0xda, 0xa0, 0x07, 0x06, 0x52, 0x02, 0x73, 0x96, 0x10, 0x64,
	0x95, 0xc0, 0x40, 0xab, 0x4d, 0xf4, 0x0e, 0x94, 0x3a, 0x02, 0x7e, 0x94, 0x10, 0x3b, 0x37, 0x65,
	0xd5, 0xca, 0xf3, 0xa8, 0x2e, 0x7b, 0xba, 0xa9, 0xa9, 0xad, 0xa3, 0x84, 0x78, 0x23, 0xac, 0xe7,
	0x09, 0xb9, 0x50, 0x90, 0x39, 0xc9, 0x4d, 0x83, 0x7a, 0x53, 0xa7, 0x2e, 0x91, 0xbc, 0xdc, 0x34,
	0xcc, 0xf5, 0x0a, 0xad, 0xc0, 0x05, 0x51, 0x61, 0x83, 0x11, 0x9c, 0xfa, 0xbb, 0x8d, 0x94, 0xdc,
	0x6a, 0x13, 0xc6, 0xed, 0xfc, 0x54, 0xb6, 0x56, 0x9c, 0x9f, 0xa8, 0xcb, 0xea, 0x17, 0x77, 0x29,
	0x65, 0xe4, 0x7d, 0xca, 0xb8, 0xa7, 0x68, 0x6f, 0x4c, 0xe0, 0x9b, 0x72, 0x8b, 0x86, 0xd0, 0xcb,
	0x50, 0x20, 0x87, 0x3c, 0xc5, 0x8d, 0x7d, 0x9c, 0xda, 0x43, 0x53, 0x56, 0xad, 0xe0, 0x0d, 0x4b,
	0xe0, 0x26, 0x4e, 0x51, 0x15, 0x86, 0xc9, 0x21, 0xf1, 0xdb, 0x9c, 0xa6, 0xf6, 0xb0, 0xac, 0xb6,
	0xf3, 0x8c, 0xa6, 0xa1, 0xa4, 0xd6, 0xa4, 0xe1, 0xd3, 0x76, 0xcc, 0xed, 0x82, 0x14, 0x8c, 0x68,
	0x70, 0x51, 0x60, 0xe8, 0x15, 0x80, 0x1d, 0x1c, 0xb6, 0xb4, 0x02, 0xa4, 0xa2, 0x20, 0x10, 0x45,
	0x4f, 0x43, 0x89, 0xb5, 0x7d, 0x9f, 0x30, 0xa6, 0x15, 0x45, 0x15, 0x43, 0x83, 0x9d, 0x18, 0x8c,
	0xe3, 0x94, 0x37, 0x78, 0x18, 0x11, 0x7b, 0x44, 0xa6, 0x58, 0x90, 0xc8, 0x56, 0x18, 0x11, 0xf4,
	0x12, 0x0c, 0x93, 0xb8, 0xa9, 0xc8, 0x92, 0x24, 0x87, 0x48, 0xdc, 0x94, 0x14, 0x82, 0x5c, 0x8c,
	0x23, 0x62, 0x97, 0x25, 0x2c, 0xd7, 0xce, 0x2f, 0x59, 0x28, 0x2d, 0xa6, 0x04, 0x73, 0x62, 0x3a,
	0xf0, 0x6e, 0xff, 0xd8, 0x85, 0x27, 0x06, 0x17, 0x2e, 0x3d, 0x7e, 0x34, 0x69, 0x57, 0x06, 0x66,
	0xc7, 0x4f, 0xee, 0xdd, 0x3d, 0xb9, 0xf3, 0xf7, 0xea, 0xd2, 0xf1, 0xfd, 0xef, 0xfe, 0xfd, 0xfa,
	0xaf, 0xe3, 0xfb, 0x0f, 0x4e, 0x7f, 0x7d, 0xd0, 0x67, 0x8a, 0xe4, 0xec, 0x49, 0x64, 0x9e, 0x39,
	0x89, 0x85, 0x2b, 0x8f, 0x1f, 0x4d, 0x4e, 0xcf, 0x4e, 0x0a, 0xf6, 0xf4, 0xc7, 0x6f, 0x4e, 0x6e,
	0xff, 0x74, 0xfa, 0xfb, 0xcf, 0x27, 0x77, 0xee, 0x1e, 0xff, 0xf9, 0x47, 0xef, 0x41, 0x9f, 0x58,
	0x67, 0x8d, 0xac, 0xcf, 0x2c, 0xd9, 0x73, 0x98, 0xa5, 0x6f, 0xc6, 0xb9, 0x27, 0x66, 0xfc, 0x84,
	0xa9, 0x07, 0x9f, 0x6f, 0xea, 0xfc, 0x39, 0x4d, 0x6d, 0xda, 0x3f, 0xd4, 0x6d, 0xff, 0xb5, 0x2b,
	0xdf, 0x5f, 0x9f, 0x01, 0xe7, 0xa1, 0x55, 0x30, 0x4a, 0xf6, 0xd0, 0x3a, 0xab, 0x7f, 0xce, 0xa7,
	0x50, 0x36, 0x63, 0x62, 0x09, 0x8d, 0x19, 0x79, 0xfa, 0x7b, 0x5b, 0x87, 0x3c, 0xe3, 0x98, 0xb7,
	0x99, 0x6c, 0x46, 0x71, 0xfe, 0x62, 0x5d, 0x5d, 0x16, 0x75, 0xb3, 0x75, 0x53, 0xb2, 0x9e, 0x56,
	0x39, 0xaf, 0x01, 0xac, 0x10, 0x33, 0x89, 0xa7, 0x86, 0x75, 0x7e, 0xb3, 0xa0, 0x28, 0x75, 0xfa,
	0xfc, 0x19, 0x90, 0xd7, 0xa1, 0x54, 0x15, 0xe7, 0x2b, 0xaa, 0xfc, 0xee, 0xbd, 0xe2, 0x49, 0x16,
	0xbd, 0x0e, 0x95, 0xa4, 0x85, 0x8f, 0xb6, 0x29, 0xdd, 0x6b, 0xec, 0x93, 0x94, 0x85, 0x34, 0x96,
	0x37, 0x49, 0xc1, 0x1b, 0x35, 0xf8, 0x4d, 0x05, 0x8b, 0xce, 0xef, 0x52, 0xc6, 0x1b, 0x11, 0x15,
	0x8d, 0x90, 0xc9, 0x17, 0x3c, 0x10, 0xd0, 0x87, 0x12, 0xe9, 0x29, 0x2c, 0x77, 0xae, 0xc2, 0x6a,
	0x50, 0x5a, 0x21, 0x7c, 0x8d, 0x06, 0xcf, 0xad, 0xed, 0x3d, 0x28, 0x1b, 0xa5, 0xae, 0xae, 0x7b,
	0x56, 0xe6, 0x5c, 0x67, 0xdd, 0x82, 0xd2, 0x8d, 0xb0, 0xc5, 0x49, 0x6a, 0xce, 0xba, 0x0c, 0xf9,
	0x04, 0x07, 0x61, 0x1c, 0xe8, 0x06, 0x95, 0x4d, 0x80, 0x8f, 0x24, 0xea, 0x69, 0x56, 0x98, 0x51,
	0xe6, 0x24, 0xad, 0xa1, 0x3a, 0x23, 0x9d, 0xba, 0x8e, 0x23, 0xd2, 0x21, 0x3b, 0xd6, 0xce, 0x76,
	0x6d, 0xec, 0xdc, 0xb6, 0xa0, 0x6c, 0xce, 0xd4, 0x59, 0x5f, 0x86, 0x41, 0x41, 0x33, 0xdb, 0x92,
	0xaf, 0xdb, 0xff, 0x87, 0xa2, 0xe8, 0x9e, 0xe4, 0x32, 0xcf, 0x4c, 0xee, 0x05, 0xad, 0x34, 0x7b,
	0x15, 0x46, 0x7a, 0x5f, 0x00, 0x54, 0x82, 0xc2, 0xc7, 0xeb, 0x4b, 0xcb, 0x37, 0x56, 0xd7, 0x97,
	0x97, 0x2a, 0x03, 0x08, 0x20, 0xbf, 0xb9, 0xba, 0xbe, 0xb2, 0xb6, 0x5c, 0xb1, 0x50, 0x11, 0x86,
	0x16, 0xbd, 0x8d, 0xf5, 0xad, 0xeb, 0x0b, 0x95, 0xcc, 0xfc, 0xbd, 0x0c, 0xe4, 0x44, 0x96, 0x68,
	0x03, 0xf2, 0xca, 0xe6, 0xe8, 0x82, 0xca, 0xbd, 0xef, 0x6e, 0xaa, 0x8e, 0xf7, 0x83, 0x2a, 0x07,
	0xa7, 0xfa, 0xd5, 0x3f, 0xc7, 0x3f, 0x64, 0xc6, 0x9d, 0x51, 0x77, 0xff, 0x4d, 0xf9, 0xb9, 0x76,
	0x7d, 0x29, 0xb8, 0x66, 0xcd, 0xa2, 0x35, 0xc8, 0xab, 0x1e, 0x99, 0x80, 0x7d, 0x53, 0x32, 0x01,
	0xfb, 0xdb, 0xe8, 0x4c, 0xc8, 0x80, 0x63, 0xa8, 0x1b, 0x70, 0x47, 0xc5, 0xf8, 0x00, 0xb2, 0x2b,
	0x84, 0x23, 0xdd, 0xd7, 0xee, 0x5b, 0x53, 0x1d, 0xeb, 0x41, 0x74, 0x90, 0x57, 0x65, 0x10, 0x1b,
	0x5d, 0xec, 0x04, 0x09, 0x08, 0x77, 0xbf, 0xd0, 0x06, 0xfc, 0x12, 0xbd, 0x0d, 0x79, 0xe5, 0x39,
	0x93, 0x59, 0x9f, 0x57, 0x4d, 0x66, 0xfd, 0xb6, 0x74, 0x06, 0x16, 0x16, 0xa0, 0xca, 0x69, 0x52,
	0xdf, 0xa1, 0x81, 0xdf, 0x8c, 0xeb, 0xf8, 0x50, 0x7f, 0xe6, 0x99, 0x94, 0x7f, 0x36, 0x13, 0x84,
	0xdc, 0x70, 0x9c, 0x26, 0x2e, 0x3e, 0x24, 0xae, 0xe2, 0xdd, 0x80, 0xd2, 0x36, 0x97, 0x89, 0x6c,
	0xe7, 0x25, 0xf4, 0xd6, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x76, 0xae, 0xbc, 0xf6, 0xe1, 0x08,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskClient interface {
	// 创建作业任务(执行作业模板)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// 筛选作业任务
	Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error)
	// 获取作业任务
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// 获取作业任务详细执行过程
	GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogResponse, error)
}

type taskClient struct {
	cc *grpc.ClientConn
}

func NewTaskClient(cc *grpc.ClientConn) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/task.Task/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error) {
	out := new(FilterResponse)
	err := c.cc.Invoke(ctx, "/task.Task/Filter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/task.Task/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) GetLog(ctx context.Context, in *GetLogRequest, opts ...grpc.CallOption) (*GetLogResponse, error) {
	out := new(GetLogResponse)
	err := c.cc.Invoke(ctx, "/task.Task/GetLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
type TaskServer interface {
	// 创建作业任务(执行作业模板)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// 筛选作业任务
	Filter(context.Context, *FilterRequest) (*FilterResponse, error)
	// 获取作业任务
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// 获取作业任务详细执行过程
	GetLog(context.Context, *GetLogRequest) (*GetLogResponse, error)
}

// UnimplementedTaskServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (*UnimplementedTaskServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTaskServer) Filter(ctx context.Context, req *FilterRequest) (*FilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Filter not implemented")
}
func (*UnimplementedTaskServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTaskServer) GetLog(ctx context.Context, req *GetLogRequest) (*GetLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLog not implemented")
}

func RegisterTaskServer(s *grpc.Server, srv TaskServer) {
	s.RegisterService(&_Task_serviceDesc, srv)
}

func _Task_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Task/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_Filter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Filter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Task/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Filter(ctx, req.(*FilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Task/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_GetLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).GetLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Task/GetLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).GetLog(ctx, req.(*GetLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Task_serviceDesc = grpc.ServiceDesc{
	ServiceName: "task.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Task_Create_Handler,
		},
		{
			MethodName: "Filter",
			Handler:    _Task_Filter_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Task_Get_Handler,
		},
		{
			MethodName: "GetLog",
			Handler:    _Task_GetLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task/task.proto",
}
