# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [auth/auth.proto](#auth/auth.proto)
    - [CallbackRequest](#auth.CallbackRequest)
    - [CallbackResponse](#auth.CallbackResponse)
    - [CheckRequest](#auth.CheckRequest)
    - [CheckResponse](#auth.CheckResponse)
    - [GetUserInfoRequest](#auth.GetUserInfoRequest)
    - [GetUserInfoResponse](#auth.GetUserInfoResponse)
    - [LoginRequest](#auth.LoginRequest)
    - [LoginResponse](#auth.LoginResponse)
    - [LogoutRequest](#auth.LogoutRequest)
    - [LogoutResponse](#auth.LogoutResponse)
    - [UserInfo](#auth.UserInfo)
  
    - [SourceCode](#auth.SourceCode)
  
  
    - [Auth](#auth.Auth)
  

- [callback/callback.proto](#callback/callback.proto)
    - [CmdbEventData](#callback.CmdbEventData)
    - [CmdbEventData.CurDataEntry](#callback.CmdbEventData.CurDataEntry)
    - [CmdbEventData.PreDataEntry](#callback.CmdbEventData.PreDataEntry)
    - [CmdbEventRequest](#callback.CmdbEventRequest)
    - [CmdbEventResponse](#callback.CmdbEventResponse)
  
  
  
    - [Callback](#callback.Callback)
  

- [cmdb/cmdb.proto](#cmdb/cmdb.proto)
    - [AreaObject](#cmdb.AreaObject)
    - [AssociationObject](#cmdb.AssociationObject)
    - [BizObject](#cmdb.BizObject)
    - [CabinetObject](#cmdb.CabinetObject)
    - [ChooseHostRequest](#cmdb.ChooseHostRequest)
    - [ChooseServerCompareRequest](#cmdb.ChooseServerCompareRequest)
    - [ChooseServerRequest](#cmdb.ChooseServerRequest)
    - [CommonObject](#cmdb.CommonObject)
    - [CreateAssociationRequest](#cmdb.CreateAssociationRequest)
    - [HostInfoObject](#cmdb.HostInfoObject)
    - [HostObject](#cmdb.HostObject)
    - [ImportCrossTableRequest](#cmdb.ImportCrossTableRequest)
    - [ImportCrossTableResponse](#cmdb.ImportCrossTableResponse)
    - [ImportHistoryObject](#cmdb.ImportHistoryObject)
    - [ImportHistoryRequest](#cmdb.ImportHistoryRequest)
    - [ImportHistoryResponse](#cmdb.ImportHistoryResponse)
    - [ImportHostRequest](#cmdb.ImportHostRequest)
    - [ImportHostResponse](#cmdb.ImportHostResponse)
    - [ImportLakeRequest](#cmdb.ImportLakeRequest)
    - [ImportLakeResponse](#cmdb.ImportLakeResponse)
    - [ImportServerRequest](#cmdb.ImportServerRequest)
    - [ImportServerResponse](#cmdb.ImportServerResponse)
    - [ImportSwitchRequest](#cmdb.ImportSwitchRequest)
    - [ImportSwitchResponse](#cmdb.ImportSwitchResponse)
    - [InstanceTopologyRequest](#cmdb.InstanceTopologyRequest)
    - [InstanceTopologyResponse](#cmdb.InstanceTopologyResponse)
    - [IspObject](#cmdb.IspObject)
    - [LakeAreaObject](#cmdb.LakeAreaObject)
    - [LakeHost](#cmdb.LakeHost)
    - [LakeObject](#cmdb.LakeObject)
    - [ModuleObject](#cmdb.ModuleObject)
    - [RoomObject](#cmdb.RoomObject)
    - [RoomTopologyRequest](#cmdb.RoomTopologyRequest)
    - [RoomTopologyResponse](#cmdb.RoomTopologyResponse)
    - [SearchHostInLakeRequest](#cmdb.SearchHostInLakeRequest)
    - [SearchHostInLakeResponse](#cmdb.SearchHostInLakeResponse)
    - [SearchHostRequest](#cmdb.SearchHostRequest)
    - [SearchHostResponse](#cmdb.SearchHostResponse)
    - [SearchLakeAreaRequest](#cmdb.SearchLakeAreaRequest)
    - [SearchLakeAreaResponse](#cmdb.SearchLakeAreaResponse)
    - [SearchLakeRequest](#cmdb.SearchLakeRequest)
    - [SearchLakeResponse](#cmdb.SearchLakeResponse)
    - [SearchMoudleRequest](#cmdb.SearchMoudleRequest)
    - [SearchMoudleResponse](#cmdb.SearchMoudleResponse)
    - [Server](#cmdb.Server)
    - [ServerListRequest](#cmdb.ServerListRequest)
    - [ServerListResponse](#cmdb.ServerListResponse)
    - [ServerObject](#cmdb.ServerObject)
    - [ServerRoomObject](#cmdb.ServerRoomObject)
    - [SetObject](#cmdb.SetObject)
    - [TopologyObject](#cmdb.TopologyObject)
    - [VipObject](#cmdb.VipObject)
    - [ZoneObject](#cmdb.ZoneObject)
  
    - [AreaLevel](#cmdb.AreaLevel)
    - [ImportStatus](#cmdb.ImportStatus)
    - [ImportType](#cmdb.ImportType)
    - [ServerInstallState](#cmdb.ServerInstallState)
    - [ServerPowerState](#cmdb.ServerPowerState)
  
  
    - [Cmdb](#cmdb.Cmdb)
  

- [common/api.proto](#common/api.proto)
    - [Paging](#common.Paging)
    - [ResponseStatus](#common.ResponseStatus)
  
    - [StatusCode](#common.StatusCode)
  
  
  

- [common/user.proto](#common/user.proto)
    - [User](#common.User)
  
  
  
  

- [file/file.proto](#file/file.proto)
    - [UploadPlaybookRequest](#file.UploadPlaybookRequest)
    - [UploadPlaybookResponse](#file.UploadPlaybookResponse)
    - [UploadRequest](#file.UploadRequest)
    - [UploadResponse](#file.UploadResponse)
  
  
  
    - [File](#file.File)
  

- [greeter/greeter.proto](#greeter/greeter.proto)
    - [HelloReply](#greeter.HelloReply)
    - [HelloRequest](#greeter.HelloRequest)
  
  
  
    - [Greeter](#greeter.Greeter)
  

- [ironic/ironicCom.proto](#ironic/ironicCom.proto)
    - [PageInfo](#ironicCom.PageInfo)
    - [ScopeInfo](#ironicCom.ScopeInfo)
    - [TcpContReq](#ironicCom.TcpContReq)
    - [TcpContRes](#ironicCom.TcpContRes)
    - [TimeScope](#ironicCom.TimeScope)
  
    - [AppKey](#ironicCom.AppKey)
    - [ResultCode](#ironicCom.ResultCode)
    - [SvcCode](#ironicCom.SvcCode)
  
  
  

- [ironic/ironicServer.proto](#ironic/ironicServer.proto)
    - [CreateNodesRootReq](#ironicServer.CreateNodesRootReq)
    - [CreateNodesRootReq.ContractRootReq](#ironicServer.CreateNodesRootReq.ContractRootReq)
    - [CreateNodesRootReq.CreateNodeInfoReq](#ironicServer.CreateNodesRootReq.CreateNodeInfoReq)
    - [CreateNodesRootReq.NodeInfoReq](#ironicServer.CreateNodesRootReq.NodeInfoReq)
    - [CreateNodesRootReq.SvcContReq](#ironicServer.CreateNodesRootReq.SvcContReq)
    - [CreateNodesRootRes](#ironicServer.CreateNodesRootRes)
    - [CreateNodesRootRes.ContractRootRes](#ironicServer.CreateNodesRootRes.ContractRootRes)
    - [CreateNodesRootRes.CreateNodeInfoRsp](#ironicServer.CreateNodesRootRes.CreateNodeInfoRsp)
    - [CreateNodesRootRes.NodeInfoRsp](#ironicServer.CreateNodesRootRes.NodeInfoRsp)
    - [CreateNodesRootRes.SvcContRes](#ironicServer.CreateNodesRootRes.SvcContRes)
    - [InstallNodeSysRootReq](#ironicServer.InstallNodeSysRootReq)
    - [InstallNodeSysRootReq.ContractRootReq](#ironicServer.InstallNodeSysRootReq.ContractRootReq)
    - [InstallNodeSysRootReq.InstallNodeSysReq](#ironicServer.InstallNodeSysRootReq.InstallNodeSysReq)
    - [InstallNodeSysRootReq.NodeInstallInfo](#ironicServer.InstallNodeSysRootReq.NodeInstallInfo)
    - [InstallNodeSysRootReq.SvcContReq](#ironicServer.InstallNodeSysRootReq.SvcContReq)
    - [InstallNodeSysRootRes](#ironicServer.InstallNodeSysRootRes)
    - [InstallNodeSysRootRes.ContractRootRes](#ironicServer.InstallNodeSysRootRes.ContractRootRes)
    - [InstallNodeSysRootRes.InstallNodeInfoRsp](#ironicServer.InstallNodeSysRootRes.InstallNodeInfoRsp)
    - [InstallNodeSysRootRes.NodeInfoRsp](#ironicServer.InstallNodeSysRootRes.NodeInfoRsp)
    - [InstallNodeSysRootRes.SvcContRes](#ironicServer.InstallNodeSysRootRes.SvcContRes)
    - [OperNodePowerRootReq](#ironicServer.OperNodePowerRootReq)
    - [OperNodePowerRootReq.ContractRootReq](#ironicServer.OperNodePowerRootReq.ContractRootReq)
    - [OperNodePowerRootReq.NodePowerOper](#ironicServer.OperNodePowerRootReq.NodePowerOper)
    - [OperNodePowerRootReq.OperNodePowerReq](#ironicServer.OperNodePowerRootReq.OperNodePowerReq)
    - [OperNodePowerRootReq.SvcContReq](#ironicServer.OperNodePowerRootReq.SvcContReq)
    - [OperNodePowerRootRes](#ironicServer.OperNodePowerRootRes)
    - [OperNodePowerRootRes.ContractRootRes](#ironicServer.OperNodePowerRootRes.ContractRootRes)
    - [OperNodePowerRootRes.NodePowerRsp](#ironicServer.OperNodePowerRootRes.NodePowerRsp)
    - [OperNodePowerRootRes.OperNodePowerRsp](#ironicServer.OperNodePowerRootRes.OperNodePowerRsp)
    - [OperNodePowerRootRes.SvcContRes](#ironicServer.OperNodePowerRootRes.SvcContRes)
    - [QryNodeInfoRootReq](#ironicServer.QryNodeInfoRootReq)
    - [QryNodeInfoRootReq.ContractRootReq](#ironicServer.QryNodeInfoRootReq.ContractRootReq)
    - [QryNodeInfoRootReq.QryNodeInfoReq](#ironicServer.QryNodeInfoRootReq.QryNodeInfoReq)
    - [QryNodeInfoRootReq.SvcContReq](#ironicServer.QryNodeInfoRootReq.SvcContReq)
    - [QryNodeInfoRootRes](#ironicServer.QryNodeInfoRootRes)
    - [QryNodeInfoRootRes.ContractRootRes](#ironicServer.QryNodeInfoRootRes.ContractRootRes)
    - [QryNodeInfoRootRes.NodeInfo](#ironicServer.QryNodeInfoRootRes.NodeInfo)
    - [QryNodeInfoRootRes.QryNodeInfoRsp](#ironicServer.QryNodeInfoRootRes.QryNodeInfoRsp)
    - [QryNodeInfoRootRes.SvcContRes](#ironicServer.QryNodeInfoRootRes.SvcContRes)
  
    - [QryNodeInfoRootRes.PowerStatus](#ironicServer.QryNodeInfoRootRes.PowerStatus)
    - [QryNodeInfoRootRes.Status](#ironicServer.QryNodeInfoRootRes.Status)
  
  
    - [IronicServer](#ironicServer.IronicServer)
  

- [playbook/playbook.proto](#playbook/playbook.proto)
    - [CreateRequest](#playbook.CreateRequest)
    - [CreateResponse](#playbook.CreateResponse)
    - [FilterRequest](#playbook.FilterRequest)
    - [FilterResponse](#playbook.FilterResponse)
    - [GetRequest](#playbook.GetRequest)
    - [GetResponse](#playbook.GetResponse)
    - [PlaybookEntrypointObject](#playbook.PlaybookEntrypointObject)
    - [ProjectObject](#playbook.ProjectObject)
    - [SaveVersionRequest](#playbook.SaveVersionRequest)
    - [SaveVersionResponse](#playbook.SaveVersionResponse)
    - [UpdateRequest](#playbook.UpdateRequest)
    - [UpdateResponse](#playbook.UpdateResponse)
  
    - [UrlType](#playbook.UrlType)
  
  
    - [Playbook](#playbook.Playbook)
  

- [schedule/schedule.proto](#schedule/schedule.proto)
    - [CreateRequest](#schedule.CreateRequest)
    - [CreateResponse](#schedule.CreateResponse)
    - [CreateScheduleObject](#schedule.CreateScheduleObject)
    - [CreateScheduleResponseObject](#schedule.CreateScheduleResponseObject)
    - [FilterRequest](#schedule.FilterRequest)
    - [FilterResponse](#schedule.FilterResponse)
    - [GetRequest](#schedule.GetRequest)
    - [GetResponse](#schedule.GetResponse)
    - [ScheduleObject](#schedule.ScheduleObject)
    - [SwitchStatusRequest](#schedule.SwitchStatusRequest)
    - [SwitchStatusResponse](#schedule.SwitchStatusResponse)
    - [UpdateRequest](#schedule.UpdateRequest)
    - [UpdateResponse](#schedule.UpdateResponse)
  
    - [ScheduleStatus](#schedule.ScheduleStatus)
    - [TaskType](#schedule.TaskType)
  
  
    - [Schedule](#schedule.Schedule)
  

- [subtask/subtask.proto](#subtask/subtask.proto)
    - [CompleteRequest](#subtask.CompleteRequest)
    - [CompleteResponse](#subtask.CompleteResponse)
    - [CreateRequest](#subtask.CreateRequest)
    - [CreateResponse](#subtask.CreateResponse)
    - [CreateServerCompareRequest](#subtask.CreateServerCompareRequest)
    - [CreateServerCompareResponse](#subtask.CreateServerCompareResponse)
    - [CreateServerRequest](#subtask.CreateServerRequest)
    - [CreateServerResponse](#subtask.CreateServerResponse)
  
    - [ServerTaskType](#subtask.ServerTaskType)
    - [SubTaskResult](#subtask.SubTaskResult)
  
  
    - [SubTask](#subtask.SubTask)
  

- [task/queue.proto](#task/queue.proto)
    - [BareMetalCreateTask](#task.BareMetalCreateTask)
    - [BareMetalInstallTask](#task.BareMetalInstallTask)
    - [BareMetalPowerTask](#task.BareMetalPowerTask)
    - [BareMetalSearchTask](#task.BareMetalSearchTask)
    - [ServerCompareTask](#task.ServerCompareTask)
    - [TemplateExecuteTask](#task.TemplateExecuteTask)
  
  
  
  

- [task/task.proto](#task/task.proto)
    - [CheckServerStateRequest](#task.CheckServerStateRequest)
    - [CheckServerStateResponse](#task.CheckServerStateResponse)
    - [CreateRequest](#task.CreateRequest)
    - [CreateResponse](#task.CreateResponse)
    - [CreateServerCompareRequest](#task.CreateServerCompareRequest)
    - [CreateServerCompareResponse](#task.CreateServerCompareResponse)
    - [CreateServerRequest](#task.CreateServerRequest)
    - [CreateServerResponse](#task.CreateServerResponse)
    - [FilterRequest](#task.FilterRequest)
    - [FilterResponse](#task.FilterResponse)
    - [GetLogRequest](#task.GetLogRequest)
    - [GetLogResponse](#task.GetLogResponse)
    - [GetRequest](#task.GetRequest)
    - [GetResponse](#task.GetResponse)
    - [InstallServerRequest](#task.InstallServerRequest)
    - [InstallServerResponse](#task.InstallServerResponse)
    - [RetryRequest](#task.RetryRequest)
    - [RetryResponse](#task.RetryResponse)
    - [ServerPowerControlRequest](#task.ServerPowerControlRequest)
    - [ServerPowerControlResponse](#task.ServerPowerControlResponse)
    - [TaskObject](#task.TaskObject)
  
    - [ScheduleType](#task.ScheduleType)
  
  
    - [Task](#task.Task)
  

- [template/template.proto](#template/template.proto)
    - [CreateRequest](#template.CreateRequest)
    - [CreateResponse](#template.CreateResponse)
    - [DeleteRequest](#template.DeleteRequest)
    - [DeleteResponse](#template.DeleteResponse)
    - [FilterRequest](#template.FilterRequest)
    - [FilterResponse](#template.FilterResponse)
    - [GetRequest](#template.GetRequest)
    - [GetResponse](#template.GetResponse)
    - [TemplateObject](#template.TemplateObject)
    - [UpdateRequest](#template.UpdateRequest)
    - [UpdateResponse](#template.UpdateResponse)
    - [UpdateStateRequest](#template.UpdateStateRequest)
    - [UpdateStateResponse](#template.UpdateStateResponse)
  
    - [TemplateState](#template.TemplateState)
  
  
    - [Template](#template.Template)
  

- [Scalar Value Types](#scalar-value-types)



<a name="auth/auth.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth/auth.proto



<a name="auth.CallbackRequest"></a>

### CallbackRequest
登陆回调接口请求






<a name="auth.CallbackResponse"></a>

### CallbackResponse
登陆回调接口响应，返回301状态码






<a name="auth.CheckRequest"></a>

### CheckRequest
服务鉴权入口请求






<a name="auth.CheckResponse"></a>

### CheckResponse
服务鉴权入口响应，返回200、401、403等状态码






<a name="auth.GetUserInfoRequest"></a>

### GetUserInfoRequest
获取用户信息请求






<a name="auth.GetUserInfoResponse"></a>

### GetUserInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [UserInfo](#auth.UserInfo) |  | 响应消息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 响应状态 |






<a name="auth.LoginRequest"></a>

### LoginRequest
用户登陆接口请求






<a name="auth.LoginResponse"></a>

### LoginResponse
用户登陆接口响应，返回301状态码






<a name="auth.LogoutRequest"></a>

### LogoutRequest
用户登出接口请求






<a name="auth.LogoutResponse"></a>

### LogoutResponse
用户登出接口响应，返回301状态码






<a name="auth.UserInfo"></a>

### UserInfo
用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  | 用户在鉴权网关中的唯一ID |
| name | [string](#string) |  | 用户名 |
| source_uid | [string](#string) |  | 用户在外部系统中的ID |
| source | [SourceCode](#auth.SourceCode) |  | 外部系统的编号 |
| avatar | [string](#string) |  | 用户头像路径 |





 


<a name="auth.SourceCode"></a>

### SourceCode
外部系统编号

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | 不依赖外部系统 |
| IAM | 1 | IAM系统 |


 

 


<a name="auth.Auth"></a>

### Auth
鉴权网关服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Check | [CheckRequest](#auth.CheckRequest) | [CheckResponse](#auth.CheckResponse) | 服务鉴权入口 |
| Login | [LoginRequest](#auth.LoginRequest) | [LoginResponse](#auth.LoginResponse) | 用户登陆接口 |
| Logout | [LogoutRequest](#auth.LogoutRequest) | [LogoutResponse](#auth.LogoutResponse) | 用户登出接口 |
| Callback | [CallbackRequest](#auth.CallbackRequest) | [CallbackResponse](#auth.CallbackResponse) | 登陆回调接口 |
| GetUserInfo | [GetUserInfoRequest](#auth.GetUserInfoRequest) | [GetUserInfoResponse](#auth.GetUserInfoResponse) | 获取用户信息接口 |

 



<a name="callback/callback.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## callback/callback.proto



<a name="callback.CmdbEventData"></a>

### CmdbEventData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cur_data | [CmdbEventData.CurDataEntry](#callback.CmdbEventData.CurDataEntry) | repeated | 时间触发前的实体数据 |
| pre_data | [CmdbEventData.PreDataEntry](#callback.CmdbEventData.PreDataEntry) | repeated | 时间触发后的实体数据 |






<a name="callback.CmdbEventData.CurDataEntry"></a>

### CmdbEventData.CurDataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="callback.CmdbEventData.PreDataEntry"></a>

### CmdbEventData.PreDataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="callback.CmdbEventRequest"></a>

### CmdbEventRequest
cmdb事件请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event_type | [string](#string) |  | 事件类型 |
| action | [string](#string) |  | 触发动作 |
| obj_type | [string](#string) |  | 事件相关的实体，如集群set |
| data | [CmdbEventData](#callback.CmdbEventData) | repeated | 事件的变更数据 |






<a name="callback.CmdbEventResponse"></a>

### CmdbEventResponse
cmdb事件请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回状态码 |





 

 

 


<a name="callback.Callback"></a>

### Callback
主机

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CmdbEvent | [CmdbEventRequest](#callback.CmdbEventRequest) | [CmdbEventResponse](#callback.CmdbEventResponse) | 接收cmdb的事件推送 |

 



<a name="cmdb/cmdb.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cmdb/cmdb.proto



<a name="cmdb.AreaObject"></a>

### AreaObject
地区对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 实例ID |
| bk_inst_name | [string](#string) |  | 区域编码 |
| area_cnname | [string](#string) |  | 区域中文名称 |
| area_enname | [string](#string) |  | 区域英文名称 |
| parent_id | [int32](#int32) |  | 父节点ID |
| level | [string](#string) |  | 地区层级 |
| area_id | [int32](#int32) |  | 地区id |






<a name="cmdb.AssociationObject"></a>

### AssociationObject
关联关系对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_obj_asst_id | [string](#string) |  | 实例关联唯一标识 |
| bk_inst_id | [int32](#int32) |  | 源实例ID |
| bk_obj_id | [string](#string) |  | 源实例对象名称 |
| bk_asst_inst_id | [int32](#int32) |  | 目标实例ID |
| bk_asst_obj_id | [string](#string) |  | 目标对象名称 |
| id | [int32](#int32) |  | 实例关联的id |






<a name="cmdb.BizObject"></a>

### BizObject
业务


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_biz_name | [string](#string) |  | 业务名 |
| bk_biz_id | [int32](#int32) |  | 业务ID |






<a name="cmdb.CabinetObject"></a>

### CabinetObject
机柜对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 实例ID |
| bk_inst_name | [string](#string) |  | 实例名 |






<a name="cmdb.ChooseHostRequest"></a>

### ChooseHostRequest
选择主机请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_biz_id | [int32](#int32) |  | 业务ID |
| bk_set_id | [int32](#int32) |  | 集群ID |
| bk_module_id | [int32](#int32) | repeated | 模块ID |






<a name="cmdb.ChooseServerCompareRequest"></a>

### ChooseServerCompareRequest
选择需要巡检校验的物理机请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int32](#int32) |  | 机房ID |






<a name="cmdb.ChooseServerRequest"></a>

### ChooseServerRequest
选择物理机请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int32](#int32) |  | 机房ID |
| server | [Server](#cmdb.Server) | repeated | 机房下对应物理机UUID及ipmi的ip |






<a name="cmdb.CommonObject"></a>

### CommonObject
通用对象，包括机房 交换机 物理机


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 实例ID |
| bk_inst_name | [string](#string) |  | 实例名称 |
| bk_obj_id | [string](#string) |  | 实例属对象 |






<a name="cmdb.CreateAssociationRequest"></a>

### CreateAssociationRequest
关联关系请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_obj_asst_id | [string](#string) |  | 实例关联唯一标识 |
| bk_inst_id | [int32](#int32) |  | 源实例ID |
| bk_asst_inst_id | [int32](#int32) |  | 目标实例ID |






<a name="cmdb.HostInfoObject"></a>

### HostInfoObject
主机内容信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [HostObject](#cmdb.HostObject) |  | 主机对象 |
| module | [ModuleObject](#cmdb.ModuleObject) | repeated | 模块对象 |
| set | [SetObject](#cmdb.SetObject) | repeated | 集群对象 |
| biz | [BizObject](#cmdb.BizObject) | repeated | 业务对象 |






<a name="cmdb.HostObject"></a>

### HostObject
主机


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_host_innerip | [string](#string) |  | 内网IP |
| bk_host_id | [int32](#int32) |  | 主机ID |
| state | [string](#string) |  | 主机状态,有上线和下线两种 |
| ipv6 | [string](#string) |  | ipv6的地址 |






<a name="cmdb.ImportCrossTableRequest"></a>

### ImportCrossTableRequest
导入交维表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | xlsx文件下载路径 |
| md5 | [string](#string) |  | xlsx文件md5 |
| filename | [string](#string) |  | 用户上传xlsx文件的文件名 |






<a name="cmdb.ImportCrossTableResponse"></a>

### ImportCrossTableResponse
导入交维表请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="cmdb.ImportHistoryObject"></a>

### ImportHistoryObject
导入历史对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | ID |
| url | [string](#string) |  | xlsx文件路径 |
| md5 | [string](#string) |  | xlsx文件md5 |
| logfile_url | [string](#string) |  | 日志文件下载路径 |
| logfile_md5 | [string](#string) |  | 日志文件md5 |
| total_count | [int32](#int32) |  | 总行数 |
| success_count | [int32](#int32) |  | 成功行数 |
| fail_count | [int32](#int32) |  | 失败行数 |
| operator | [common.User](#common.User) |  | 操作人 |
| import_type | [ImportType](#cmdb.ImportType) |  | 导入类型 |
| status | [ImportStatus](#cmdb.ImportStatus) |  | 导入状态 |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 起始时间 |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 结束时间 |
| filename | [string](#string) |  | 用户上传的文件名 |
| logfile_content | [string](#string) |  | 日志文件内容 |






<a name="cmdb.ImportHistoryRequest"></a>

### ImportHistoryRequest
查询导入历史列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |






<a name="cmdb.ImportHistoryResponse"></a>

### ImportHistoryResponse
查询导入历史列表返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| history | [ImportHistoryObject](#cmdb.ImportHistoryObject) | repeated | 导入历史对象列表 |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="cmdb.ImportHostRequest"></a>

### ImportHostRequest
导入主机请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | xlsx文件下载路径 |
| md5 | [string](#string) |  | xlsx文件md5 |
| filename | [string](#string) |  | 用户上传xlsx文件的文件名 |






<a name="cmdb.ImportHostResponse"></a>

### ImportHostResponse
导入主机请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="cmdb.ImportLakeRequest"></a>

### ImportLakeRequest
导入lake节点请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | xlsx文件下载路径 |
| md5 | [string](#string) |  | xlsx文件md5 |
| filename | [string](#string) |  | 用户上传xlsx文件的文件名 |






<a name="cmdb.ImportLakeResponse"></a>

### ImportLakeResponse
导入lake节点请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="cmdb.ImportServerRequest"></a>

### ImportServerRequest
导入物理机请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | xlsx文件下载路径 |
| md5 | [string](#string) |  | xlsx文件md5 |
| filename | [string](#string) |  | 用户上传xlsx文件的文件名 |






<a name="cmdb.ImportServerResponse"></a>

### ImportServerResponse
导入物理机请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="cmdb.ImportSwitchRequest"></a>

### ImportSwitchRequest
导入交换机请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | xlsx文件下载路径 |
| md5 | [string](#string) |  | xlsx文件md5 |
| filename | [string](#string) |  | 用户上传xlsx文件的文件名 |






<a name="cmdb.ImportSwitchResponse"></a>

### ImportSwitchResponse
导入交换机请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="cmdb.InstanceTopologyRequest"></a>

### InstanceTopologyRequest
拓扑实例请求

int32 level = 1;






<a name="cmdb.InstanceTopologyResponse"></a>

### InstanceTopologyResponse
拓扑实例请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| instance | [TopologyObject](#cmdb.TopologyObject) | repeated | 多个业务拓扑对象 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | TopologyObject instance = 1; 状态码 |






<a name="cmdb.IspObject"></a>

### IspObject
运营商对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 实例ID |
| bk_inst_name | [string](#string) |  | 实例名 |
| isp_code | [string](#string) |  | 运营商编码 |
| isp_enname | [string](#string) |  | 运营商英文名称 |
| isp_abbr | [string](#string) |  | 运营商缩写英文名 |






<a name="cmdb.LakeAreaObject"></a>

### LakeAreaObject
通过IP获取节点的地区和位置信息的返回的具体结构体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ipv4 | [string](#string) |  |  |
| ipv6 | [string](#string) |  |  |
| lake | [LakeObject](#cmdb.LakeObject) | repeated | 对应的lake节点信息,包括机房信息 |
| host_state | [string](#string) |  | 主机的状态，有上线和下线两种 |






<a name="cmdb.LakeHost"></a>

### LakeHost
lake下的主机列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lake_name | [string](#string) |  |  |
| node_id | [int32](#int32) |  | 缓存节点id |
| host | [HostObject](#cmdb.HostObject) | repeated |  |






<a name="cmdb.LakeObject"></a>

### LakeObject
LAKE节点对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 业务ID |
| bk_inst_name | [string](#string) |  | 实例名 |
| node_id | [int32](#int32) |  | 节点ID |
| vip | [VipObject](#cmdb.VipObject) | repeated | 节点VIP |
| bandwidth_limit | [float](#float) |  | 能力上限 |
| node_level | [string](#string) |  | 节点层次 |
| service_state | [string](#string) |  | 节点服务状态 |
| construct_state | [string](#string) |  | 节点建设状态 |
| construct_updatetime | [string](#string) |  | 节点建设状态变更时间 |
| room | [ServerRoomObject](#cmdb.ServerRoomObject) | repeated | 机房信息 |






<a name="cmdb.ModuleObject"></a>

### ModuleObject
模块


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_module_name | [string](#string) |  | 模块名 |
| bk_module_id | [int32](#int32) |  | 模块ID |






<a name="cmdb.RoomObject"></a>

### RoomObject
机房对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 业务ID |
| bk_inst_name | [string](#string) |  | 业务名称 |
| bk_obj_id | [string](#string) |  | 对象ID |
| room_name | [string](#string) |  | 机房名称 |
| city | [string](#string) |  | 城市 |
| prov | [string](#string) |  | 省份 |
| area_code | [string](#string) |  | 地区编码 |
| child | [RoomObject](#cmdb.RoomObject) | repeated | 子对象 |






<a name="cmdb.RoomTopologyRequest"></a>

### RoomTopologyRequest
机房拓扑请求






<a name="cmdb.RoomTopologyResponse"></a>

### RoomTopologyResponse
机房拓扑请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rooms | [RoomObject](#cmdb.RoomObject) | repeated | 机房拓扑对象 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="cmdb.SearchHostInLakeRequest"></a>

### SearchHostInLakeRequest
查询lake下的主机列表的请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lake_name | [string](#string) | repeated |  |






<a name="cmdb.SearchHostInLakeResponse"></a>

### SearchHostInLakeResponse
查询lake下的主机列表的返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lake_host | [LakeHost](#cmdb.LakeHost) | repeated |  |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="cmdb.SearchHostRequest"></a>

### SearchHostRequest
查找主机请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| host | [HostObject](#cmdb.HostObject) |  | 主机对象 |
| module | [ModuleObject](#cmdb.ModuleObject) |  | 模块对象 |
| set | [SetObject](#cmdb.SetObject) |  | 集群对象 |
| biz | [BizObject](#cmdb.BizObject) |  | 业务对象 |
| zone | [ZoneObject](#cmdb.ZoneObject) |  | 区域对象 |






<a name="cmdb.SearchHostResponse"></a>

### SearchHostResponse
查找主机请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |
| instance | [TopologyObject](#cmdb.TopologyObject) | repeated | 主机信息 |






<a name="cmdb.SearchLakeAreaRequest"></a>

### SearchLakeAreaRequest
通过IP获取节点的地区和位置信息的请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ipv4 | [string](#string) | repeated | ipv4,可传多个 |
| ipv6 | [string](#string) | repeated | ipv6,可传多个 |
| area_level | [AreaLevel](#cmdb.AreaLevel) |  | 要查询的区域的地区层级 |






<a name="cmdb.SearchLakeAreaResponse"></a>

### SearchLakeAreaResponse
通过IP获取节点的地区和位置信息的返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lake_area | [LakeAreaObject](#cmdb.LakeAreaObject) | repeated |  |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="cmdb.SearchLakeRequest"></a>

### SearchLakeRequest
查找Lake请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node_id | [int32](#int32) | repeated | 节点ID(一个或多个，可为空) |
| vip | [string](#string) | repeated | 节点VIP(一个或多个，可为空) |
| area_code | [string](#string) | repeated | 地区编码(一个或多个，可为空) |
| isp_code | [string](#string) | repeated | 运营商编码(一个或多个，可为空) |






<a name="cmdb.SearchLakeResponse"></a>

### SearchLakeResponse
查找Lake请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |
| lake | [LakeObject](#cmdb.LakeObject) | repeated | 主机信息 |






<a name="cmdb.SearchMoudleRequest"></a>

### SearchMoudleRequest
查询模块请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cmdb_search_request | [ChooseHostRequest](#cmdb.ChooseHostRequest) |  |  |






<a name="cmdb.SearchMoudleResponse"></a>

### SearchMoudleResponse
查询模块请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |
| moudule | [ModuleObject](#cmdb.ModuleObject) | repeated | 模块对象 |






<a name="cmdb.Server"></a>

### Server
物理机传入信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | 裸金属uuid |
| ipmi_ip | [string](#string) |  | 物理机ipmi的ip |






<a name="cmdb.ServerListRequest"></a>

### ServerListRequest
裸金属列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int32](#int32) |  | 机房id，相当于 RoomObject 中的 bk_inst_id |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |






<a name="cmdb.ServerListResponse"></a>

### ServerListResponse
裸金属列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| servers | [ServerObject](#cmdb.ServerObject) | repeated | 裸金属列表 |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="cmdb.ServerObject"></a>

### ServerObject
物理机对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 物理机id |
| bk_inst_name | [string](#string) |  | 实例名 |
| manage_ip | [string](#string) |  | IPMI地址 |
| account | [string](#string) |  | IPMI账号 |
| password | [string](#string) |  | IPMI密码 |
| mac_address | [string](#string) |  | MAC地址 |
| cabinet | [string](#string) |  | 机柜 |
| port | [string](#string) |  | IPMI端口 |
| power_state | [ServerPowerState](#cmdb.ServerPowerState) |  | 电源状态 |
| install_state | [ServerInstallState](#cmdb.ServerInstallState) |  | 装机状态 |
| uuid | [string](#string) |  | 裸金属uuid |






<a name="cmdb.ServerRoomObject"></a>

### ServerRoomObject
机房对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 机房ID |
| bk_inst_name | [string](#string) |  | 机房名称 |
| state | [string](#string) |  | 机房状态 |
| area_code | [string](#string) |  | 地区编码 |
| isp_code | [string](#string) |  | 运营商编码 |
| area | [AreaObject](#cmdb.AreaObject) |  | 机房的具体区域信息 |
| room_address | [string](#string) |  | 机房的具体地址 |






<a name="cmdb.SetObject"></a>

### SetObject
集群


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_set_name | [string](#string) |  | 集群名 |
| bk_set_id | [int32](#int32) |  | 集群ID |






<a name="cmdb.TopologyObject"></a>

### TopologyObject
业务拓扑对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 业务ID |
| bk_inst_name | [string](#string) |  | 业务名称 |
| bk_obj_id | [string](#string) |  | 对象ID |
| bk_obj_name | [string](#string) |  | 对象名称 |
| uuid | [int32](#int32) |  | 唯一标识 |
| child | [TopologyObject](#cmdb.TopologyObject) | repeated | 子对象 |






<a name="cmdb.VipObject"></a>

### VipObject
VIP对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 实例ID |
| bk_inst_name | [string](#string) |  | ip的ipv4地址 |
| biz_type | [string](#string) |  | vip类型 |
| vip_ipv6 | [string](#string) |  | vip的ipv6地址 |
| vip_ipv4 | [string](#string) |  | vip的ipv4地址 |






<a name="cmdb.ZoneObject"></a>

### ZoneObject
区域


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 区域ID |





 


<a name="cmdb.AreaLevel"></a>

### AreaLevel
地区层级

| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 | 0-不存在 |
| NATION | 1 | 1-国家级 |
| REGION | 2 | 2-大区级 |
| PROVINCE | 3 | 3-省级 |
| CITY | 4 | 4-市级 |
| COUNTY | 5 | 5-区县级 |



<a name="cmdb.ImportStatus"></a>

### ImportStatus
导入状态: 0-正在导入 1-导入完成 2-导入失败

| Name | Number | Description |
| ---- | ------ | ----------- |
| IMPORTING | 0 | 0-正在导入 |
| COMPLETED | 1 | 1-导入完成 |
| FAILED | 2 | 2-导入失败 |



<a name="cmdb.ImportType"></a>

### ImportType
导入类型:0-undefined 1-主机 2-物理机 3-LAKE 4-交换机

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 | 0-undefined |
| HOST | 1 | 1-主机 |
| SERVER | 2 | 2-物理机 |
| LAKE | 3 | 3-LAKE |
| SWITCH | 4 | 4-交换机 |
| CROSSTABLE | 5 | 5-交维表 |



<a name="cmdb.ServerInstallState"></a>

### ServerInstallState
装机状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNINSTALLED | 0 | 未装机 |
| INSTALLING | 1 | 装机中 |
| INSTALLED | 2 | 已装机 |



<a name="cmdb.ServerPowerState"></a>

### ServerPowerState
电源状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | 未知状态 |
| OFF | 1 | 下电 |
| ON | 2 | 上电 |


 

 


<a name="cmdb.Cmdb"></a>

### Cmdb
主机

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| InstanceTopology | [InstanceTopologyRequest](#cmdb.InstanceTopologyRequest) | [InstanceTopologyResponse](#cmdb.InstanceTopologyResponse) | 实例拓扑 |
| SearchHost | [SearchHostRequest](#cmdb.SearchHostRequest) | [SearchHostResponse](#cmdb.SearchHostResponse) | 查找主机 |
| ImportHost | [ImportHostRequest](#cmdb.ImportHostRequest) | [ImportHostResponse](#cmdb.ImportHostResponse) | 导入主机 |
| ImportCrossTable | [ImportCrossTableRequest](#cmdb.ImportCrossTableRequest) | [ImportCrossTableResponse](#cmdb.ImportCrossTableResponse) | 导入交维表 |
| ImportLake | [ImportLakeRequest](#cmdb.ImportLakeRequest) | [ImportLakeResponse](#cmdb.ImportLakeResponse) | 导入lake节点 |
| RoomTopology | [RoomTopologyRequest](#cmdb.RoomTopologyRequest) | [RoomTopologyResponse](#cmdb.RoomTopologyResponse) | 机房拓扑 |
| ServerList | [ServerListRequest](#cmdb.ServerListRequest) | [ServerListResponse](#cmdb.ServerListResponse) | 裸金属列表 |
| SearchModule | [SearchMoudleRequest](#cmdb.SearchMoudleRequest) | [SearchMoudleResponse](#cmdb.SearchMoudleResponse) | 查询模块 |
| ImportHistory | [ImportHistoryRequest](#cmdb.ImportHistoryRequest) | [ImportHistoryResponse](#cmdb.ImportHistoryResponse) | 查询导入历史记录列表 |
| SearchLake | [SearchLakeRequest](#cmdb.SearchLakeRequest) | [SearchLakeResponse](#cmdb.SearchLakeResponse) | 查询Lake节点 |
| SearchHostInLake | [SearchHostInLakeRequest](#cmdb.SearchHostInLakeRequest) | [SearchHostInLakeResponse](#cmdb.SearchHostInLakeResponse) | 查询Lake节点下的主机列表 |
| SearchLakeArea | [SearchLakeAreaRequest](#cmdb.SearchLakeAreaRequest) | [SearchLakeAreaResponse](#cmdb.SearchLakeAreaResponse) | 通过IP获取节点的地区和位置信息 |

 



<a name="common/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common/api.proto



<a name="common.Paging"></a>

### Paging
分页


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_page | [int32](#int32) |  | 总页数 |
| page | [int32](#int32) |  | 当前页数 |
| per_page | [int32](#int32) |  | 每页显示的记录条数 |
| total_record | [int32](#int32) |  | 总记录数 |






<a name="common.ResponseStatus"></a>

### ResponseStatus
请求返回状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [StatusCode](#common.StatusCode) |  | 状态码 |
| message | [string](#string) |  | 信息 |





 


<a name="common.StatusCode"></a>

### StatusCode
请求返回状态码

| Name | Number | Description |
| ---- | ------ | ----------- |
| SUCCESS | 0 | 成功 |
| INVALID_ARGUMENT | 400 | 参数错误 |
| ACCESS_DENIED | 403 | 访问拒绝 |


 

 

 



<a name="common/user.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common/user.proto



<a name="common.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int32](#int32) |  |  |
| user_name | [string](#string) |  |  |





 

 

 

 



<a name="file/file.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## file/file.proto



<a name="file.UploadPlaybookRequest"></a>

### UploadPlaybookRequest
上传playbook压缩包并且解析入口yml文件的请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [bytes](#bytes) |  | 上传的playbook压缩包 |






<a name="file.UploadPlaybookResponse"></a>

### UploadPlaybookResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | 解析出来的playbook项目存储在对象存储的url |
| md5 | [string](#string) |  | 解析出来的playbook项目文件md5 |
| filesize | [string](#string) |  | 解析出来的playbook项目文件大小 |
| entrypoint | [string](#string) | repeated | 解析出来的playbook项目入口yml文件，有多个 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="file.UploadRequest"></a>

### UploadRequest
上传csv等通用文件


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [bytes](#bytes) |  | 文件字段 |






<a name="file.UploadResponse"></a>

### UploadResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | 通用文件对象存储的url |
| md5 | [string](#string) |  | 解析出来的通用文件md5 |
| filesize | [string](#string) |  | 解析出来的通用文件大小 |
| filename | [string](#string) |  | 解析出来的通用文件name |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |





 

 

 


<a name="file.File"></a>

### File
文件服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UploadPlaybook | [UploadPlaybookRequest](#file.UploadPlaybookRequest) | [UploadPlaybookResponse](#file.UploadPlaybookResponse) | 上传playbook压缩包并且解析入口yml文件 |
| Upload | [UploadRequest](#file.UploadRequest) | [UploadResponse](#file.UploadResponse) | 上传csv等通用文件 |

 



<a name="greeter/greeter.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## greeter/greeter.proto



<a name="greeter.HelloReply"></a>

### HelloReply
The response message containing the greetings


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="greeter.HelloRequest"></a>

### HelloRequest
The request message containing the user&#39;s name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |





 

 

 


<a name="greeter.Greeter"></a>

### Greeter
The greeting service definition.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SayHello | [HelloRequest](#greeter.HelloRequest) | [HelloReply](#greeter.HelloReply) | Sends a greeting |

 



<a name="ironic/ironicCom.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ironic/ironicCom.proto



<a name="ironicCom.PageInfo"></a>

### PageInfo
分页信息(每个服务通用的控制信息，当客户端发起请求时，信息填写如下)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pageIndex | [int32](#int32) |  | 当前页 |
| pageSize | [int32](#int32) |  | 每页数量 |
| pageCount | [int32](#int32) |  | 总页数 |
| rowCount | [int32](#int32) |  | 总记录数 |






<a name="ironicCom.ScopeInfo"></a>

### ScopeInfo
查询范围信息(继承PageInfo。详情查询类服务中，用于指定具体子表的查询范围及分页信息。)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scope | [string](#string) |  | 范围编码 |
| pageIndex | [int32](#int32) |  | 当前页 |
| pageSize | [int32](#int32) |  | 每页数量 |
| pageCount | [int32](#int32) |  | 总页数 |
| rowCount | [int32](#int32) |  | 总记录数 |






<a name="ironicCom.TcpContReq"></a>

### TcpContReq
控制对象(每个服务通用的控制信息，当客户端发起请求时，信息填写如下)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| svcCode | [SvcCode](#ironicCom.SvcCode) |  | 服务编码 |
| apiCode | [string](#string) |  | 能力编码 |
| appKey | [AppKey](#ironicCom.AppKey) |  | 调用方编码 |
| dstSysID | [AppKey](#ironicCom.AppKey) |  | 提供方编码 |
| transactionID | [string](#string) |  | 交易流水号 |
| reqTime | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 请求时间 |
| sign | [string](#string) |  | 签名字符串 |
| version | [string](#string) |  | 服务版本 |






<a name="ironicCom.TcpContRes"></a>

### TcpContRes
控制对象(每个服务通用的控制信息，当服务端反馈应答时，信息填写如下)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transactionID | [string](#string) |  | 交易流水号 |
| reqTime | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 应答时间 |
| sign | [string](#string) |  | 签名字符串 |






<a name="ironicCom.TimeScope"></a>

### TimeScope
时间范围信息(用于指定时间范围。如按受理开始时间、受理结束时间查询订单列表。按某个具体时间查询，不使用此对象。)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beginDate | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 起始时间 |
| endDate | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 结束时间 |





 


<a name="ironicCom.AppKey"></a>

### AppKey
调用方编码

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 |  |
| IRONICAPP | 1000010000 | 裸金属应用 |
| CMDB | 2000020000 | CMDB平台 |



<a name="ironicCom.ResultCode"></a>

### ResultCode
响应码

| Name | Number | Description |
| ---- | ------ | ----------- |
| SUCCESS | 0 | 成功 |
| FAIL | 1 | 失败 |



<a name="ironicCom.SvcCode"></a>

### SvcCode
服务编码

| Name | Number | Description |
| ---- | ------ | ----------- |
| SVCCODEDEFAULT | 0 |  |
| CREATENODES | 1010010001 | 创建裸金属节点 |
| QRYNODEINFO | 1010010002 | 查询裸金属状态 |
| INSTALLNODESYS | 1010010003 | 安装裸金属系统 |
| OPERNODEPOWER | 1010010004 | 操作裸金属实例电源 |


 

 

 



<a name="ironic/ironicServer.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ironic/ironicServer.proto



<a name="ironicServer.CreateNodesRootReq"></a>

### CreateNodesRootReq
--------------- 创建裸金属节点---------------//


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootReq | [CreateNodesRootReq.ContractRootReq](#ironicServer.CreateNodesRootReq.ContractRootReq) |  |  |






<a name="ironicServer.CreateNodesRootReq.ContractRootReq"></a>

### CreateNodesRootReq.ContractRootReq
请求信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContReq](#ironicCom.TcpContReq) |  |  |
| svcCont | [CreateNodesRootReq.SvcContReq](#ironicServer.CreateNodesRootReq.SvcContReq) |  |  |






<a name="ironicServer.CreateNodesRootReq.CreateNodeInfoReq"></a>

### CreateNodesRootReq.CreateNodeInfoReq
创建节点请求对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeInfoReqs | [CreateNodesRootReq.NodeInfoReq](#ironicServer.CreateNodesRootReq.NodeInfoReq) | repeated | 节点 |
| pageInfo | [ironicCom.PageInfo](#ironicCom.PageInfo) |  | 分页信息 |






<a name="ironicServer.CreateNodesRootReq.NodeInfoReq"></a>

### CreateNodesRootReq.NodeInfoReq
节点信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serverAddr | [string](#string) |  | 裸金属区域组件服务地址 |
| regionName | [string](#string) |  | 区域名称 |
| initId | [int64](#int64) |  | 初始化id |
| ip | [string](#string) |  | ipmi的IP地址 |
| userName | [string](#string) |  | ipmi用户名 |
| password | [string](#string) |  | Ipmi密码 |
| macAddr | [string](#string) |  | 裸金属主机mac地址 |
| port | [int64](#int64) |  | Ipmi端口 |






<a name="ironicServer.CreateNodesRootReq.SvcContReq"></a>

### CreateNodesRootReq.SvcContReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requestObject | [CreateNodesRootReq.CreateNodeInfoReq](#ironicServer.CreateNodesRootReq.CreateNodeInfoReq) |  | 节点信息列表 |






<a name="ironicServer.CreateNodesRootRes"></a>

### CreateNodesRootRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootRes | [CreateNodesRootRes.ContractRootRes](#ironicServer.CreateNodesRootRes.ContractRootRes) |  |  |






<a name="ironicServer.CreateNodesRootRes.ContractRootRes"></a>

### CreateNodesRootRes.ContractRootRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContRes](#ironicCom.TcpContRes) |  |  |
| svcCont | [CreateNodesRootRes.SvcContRes](#ironicServer.CreateNodesRootRes.SvcContRes) |  |  |






<a name="ironicServer.CreateNodesRootRes.CreateNodeInfoRsp"></a>

### CreateNodesRootRes.CreateNodeInfoRsp
创建节点请求对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeInfoRsps | [CreateNodesRootRes.NodeInfoRsp](#ironicServer.CreateNodesRootRes.NodeInfoRsp) | repeated | 节点 |
| pageInfo | [ironicCom.PageInfo](#ironicCom.PageInfo) |  | 分页信息 |






<a name="ironicServer.CreateNodesRootRes.NodeInfoRsp"></a>

### CreateNodesRootRes.NodeInfoRsp
节点信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |
| initId | [int64](#int64) |  | 初始化id |
| createResultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 创建节点响应码 |
| createResultMsg | [string](#string) |  | 创建节点响应消息 |






<a name="ironicServer.CreateNodesRootRes.SvcContRes"></a>

### CreateNodesRootRes.SvcContRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 响应码 |
| resultMsg | [string](#string) |  | 响应消息描述 |
| resultObject | [CreateNodesRootRes.CreateNodeInfoRsp](#ironicServer.CreateNodesRootRes.CreateNodeInfoRsp) |  | 响应对象 |






<a name="ironicServer.InstallNodeSysRootReq"></a>

### InstallNodeSysRootReq
--------------- 安装裸金属系统 ---------------//


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootReq | [InstallNodeSysRootReq.ContractRootReq](#ironicServer.InstallNodeSysRootReq.ContractRootReq) |  |  |






<a name="ironicServer.InstallNodeSysRootReq.ContractRootReq"></a>

### InstallNodeSysRootReq.ContractRootReq
请求信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContReq](#ironicCom.TcpContReq) |  |  |
| svcCont | [InstallNodeSysRootReq.SvcContReq](#ironicServer.InstallNodeSysRootReq.SvcContReq) |  |  |






<a name="ironicServer.InstallNodeSysRootReq.InstallNodeSysReq"></a>

### InstallNodeSysRootReq.InstallNodeSysReq
安装裸金属系统请求对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeInstallInfos | [InstallNodeSysRootReq.NodeInstallInfo](#ironicServer.InstallNodeSysRootReq.NodeInstallInfo) | repeated | 安装裸金属系统信息 |






<a name="ironicServer.InstallNodeSysRootReq.NodeInstallInfo"></a>

### InstallNodeSysRootReq.NodeInstallInfo
安装裸金属系统信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |
| imageAddr | [string](#string) |  | 安装镜像地址(一般为http url) |
| imageCheckSum | [string](#string) |  | 镜像MD5校验码 |






<a name="ironicServer.InstallNodeSysRootReq.SvcContReq"></a>

### InstallNodeSysRootReq.SvcContReq
安装裸金属系统请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requestObject | [InstallNodeSysRootReq.InstallNodeSysReq](#ironicServer.InstallNodeSysRootReq.InstallNodeSysReq) |  |  |






<a name="ironicServer.InstallNodeSysRootRes"></a>

### InstallNodeSysRootRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootRes | [InstallNodeSysRootRes.ContractRootRes](#ironicServer.InstallNodeSysRootRes.ContractRootRes) |  |  |






<a name="ironicServer.InstallNodeSysRootRes.ContractRootRes"></a>

### InstallNodeSysRootRes.ContractRootRes
响应信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContRes](#ironicCom.TcpContRes) |  |  |
| svcCont | [InstallNodeSysRootRes.SvcContRes](#ironicServer.InstallNodeSysRootRes.SvcContRes) |  |  |






<a name="ironicServer.InstallNodeSysRootRes.InstallNodeInfoRsp"></a>

### InstallNodeSysRootRes.InstallNodeInfoRsp
创建节点请求对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeInfoRsps | [InstallNodeSysRootRes.NodeInfoRsp](#ironicServer.InstallNodeSysRootRes.NodeInfoRsp) | repeated | 节点 |
| pageInfo | [ironicCom.PageInfo](#ironicCom.PageInfo) |  | 分页信息 |






<a name="ironicServer.InstallNodeSysRootRes.NodeInfoRsp"></a>

### InstallNodeSysRootRes.NodeInfoRsp
节点信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |
| initId | [int64](#int64) |  | 初始化id |
| installResultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 创建节点响应码 |
| installResultMsg | [string](#string) |  | 创建节点响应消息 |






<a name="ironicServer.InstallNodeSysRootRes.SvcContRes"></a>

### InstallNodeSysRootRes.SvcContRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 响应码 |
| resultMsg | [string](#string) |  | 响应消息描述 |
| resultObject | [InstallNodeSysRootRes.InstallNodeInfoRsp](#ironicServer.InstallNodeSysRootRes.InstallNodeInfoRsp) |  | 响应对象 |






<a name="ironicServer.OperNodePowerRootReq"></a>

### OperNodePowerRootReq
--------------- 操作裸金属实例电源 ---------------//


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootReq | [OperNodePowerRootReq.ContractRootReq](#ironicServer.OperNodePowerRootReq.ContractRootReq) |  |  |






<a name="ironicServer.OperNodePowerRootReq.ContractRootReq"></a>

### OperNodePowerRootReq.ContractRootReq
请求信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContReq](#ironicCom.TcpContReq) |  |  |
| svcCont | [OperNodePowerRootReq.SvcContReq](#ironicServer.OperNodePowerRootReq.SvcContReq) |  |  |






<a name="ironicServer.OperNodePowerRootReq.NodePowerOper"></a>

### OperNodePowerRootReq.NodePowerOper



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |






<a name="ironicServer.OperNodePowerRootReq.OperNodePowerReq"></a>

### OperNodePowerRootReq.OperNodePowerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodePowerOpers | [OperNodePowerRootReq.NodePowerOper](#ironicServer.OperNodePowerRootReq.NodePowerOper) | repeated | 安装裸金属系统信息 |






<a name="ironicServer.OperNodePowerRootReq.SvcContReq"></a>

### OperNodePowerRootReq.SvcContReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requestObject | [OperNodePowerRootReq.OperNodePowerReq](#ironicServer.OperNodePowerRootReq.OperNodePowerReq) |  |  |






<a name="ironicServer.OperNodePowerRootRes"></a>

### OperNodePowerRootRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootRes | [OperNodePowerRootRes.ContractRootRes](#ironicServer.OperNodePowerRootRes.ContractRootRes) |  |  |






<a name="ironicServer.OperNodePowerRootRes.ContractRootRes"></a>

### OperNodePowerRootRes.ContractRootRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContRes](#ironicCom.TcpContRes) |  |  |
| svcCont | [OperNodePowerRootRes.SvcContRes](#ironicServer.OperNodePowerRootRes.SvcContRes) |  |  |






<a name="ironicServer.OperNodePowerRootRes.NodePowerRsp"></a>

### OperNodePowerRootRes.NodePowerRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |
| powerResultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 创建节点响应码 |
| powerResultMsg | [string](#string) |  | 创建节点响应消息 |






<a name="ironicServer.OperNodePowerRootRes.OperNodePowerRsp"></a>

### OperNodePowerRootRes.OperNodePowerRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodePowerRsps | [OperNodePowerRootRes.NodePowerRsp](#ironicServer.OperNodePowerRootRes.NodePowerRsp) | repeated | 电源操作应答 |






<a name="ironicServer.OperNodePowerRootRes.SvcContRes"></a>

### OperNodePowerRootRes.SvcContRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 响应码 |
| resultMsg | [string](#string) |  | 响应消息描述 |
| resultObject | [OperNodePowerRootRes.OperNodePowerRsp](#ironicServer.OperNodePowerRootRes.OperNodePowerRsp) |  | 响应对象 |






<a name="ironicServer.QryNodeInfoRootReq"></a>

### QryNodeInfoRootReq
--------------- 裸金属节点状态查询---------------//


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootReq | [QryNodeInfoRootReq.ContractRootReq](#ironicServer.QryNodeInfoRootReq.ContractRootReq) |  |  |






<a name="ironicServer.QryNodeInfoRootReq.ContractRootReq"></a>

### QryNodeInfoRootReq.ContractRootReq
请求信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContReq](#ironicCom.TcpContReq) |  |  |
| svcCont | [QryNodeInfoRootReq.SvcContReq](#ironicServer.QryNodeInfoRootReq.SvcContReq) |  |  |






<a name="ironicServer.QryNodeInfoRootReq.QryNodeInfoReq"></a>

### QryNodeInfoRootReq.QryNodeInfoReq
查询节点信息请求对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeIds | [string](#string) | repeated | 裸金属节点id |






<a name="ironicServer.QryNodeInfoRootReq.SvcContReq"></a>

### QryNodeInfoRootReq.SvcContReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requestObject | [QryNodeInfoRootReq.QryNodeInfoReq](#ironicServer.QryNodeInfoRootReq.QryNodeInfoReq) |  | 节点id列表 |






<a name="ironicServer.QryNodeInfoRootRes"></a>

### QryNodeInfoRootRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootRes | [QryNodeInfoRootRes.ContractRootRes](#ironicServer.QryNodeInfoRootRes.ContractRootRes) |  |  |






<a name="ironicServer.QryNodeInfoRootRes.ContractRootRes"></a>

### QryNodeInfoRootRes.ContractRootRes
响应信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContRes](#ironicCom.TcpContRes) |  |  |
| svcCont | [QryNodeInfoRootRes.SvcContRes](#ironicServer.QryNodeInfoRootRes.SvcContRes) |  |  |






<a name="ironicServer.QryNodeInfoRootRes.NodeInfo"></a>

### QryNodeInfoRootRes.NodeInfo
节点信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |
| status | [QryNodeInfoRootRes.Status](#ironicServer.QryNodeInfoRootRes.Status) |  | 当前状态 |
| ip | [string](#string) |  | ipmi的IP地址 |
| userName | [string](#string) |  | ipmi用户名 |
| password | [string](#string) |  | Ipmi密码 |
| macAddr | [string](#string) |  | 裸金属主机mac地址 |
| imagerAddr | [string](#string) |  | 镜像地址 |
| serverAddr | [string](#string) |  | 所属区域裸金属组件服务地址 |
| qryResultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 查询节点响应码 |
| qryResultMsg | [string](#string) |  | 查询节点响应消息 |
| powerStatus | [QryNodeInfoRootRes.PowerStatus](#ironicServer.QryNodeInfoRootRes.PowerStatus) |  | 电源状态 |






<a name="ironicServer.QryNodeInfoRootRes.QryNodeInfoRsp"></a>

### QryNodeInfoRootRes.QryNodeInfoRsp
查询节点应答对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeInfos | [QryNodeInfoRootRes.NodeInfo](#ironicServer.QryNodeInfoRootRes.NodeInfo) | repeated | 节点 |






<a name="ironicServer.QryNodeInfoRootRes.SvcContRes"></a>

### QryNodeInfoRootRes.SvcContRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 响应码 |
| resultMsg | [string](#string) |  | 响应消息描述 |
| resultObject | [QryNodeInfoRootRes.QryNodeInfoRsp](#ironicServer.QryNodeInfoRootRes.QryNodeInfoRsp) |  | 响应对象 |





 


<a name="ironicServer.QryNodeInfoRootRes.PowerStatus"></a>

### QryNodeInfoRootRes.PowerStatus
电源状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | 未知 |
| POWERON | 1 | 开机 |
| POWEROFF | 2 | 关机 |



<a name="ironicServer.QryNodeInfoRootRes.Status"></a>

### QryNodeInfoRootRes.Status
裸金属节点状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 | 未知 |
| CREATED | 1300 | 已创建 :刚创建完节点后的状态，节点创建、验证成功，可进行装系统 |
| DEPLOYING | 1200 | 部署中 :正在执行装机任务 |
| DEPLOYFAILED | 1100 | 部署失败 : 执行装机异常 |
| ACTIVE | 1000 | 活动 ： 已经安装完系统活动中 |


 

 


<a name="ironicServer.IronicServer"></a>

### IronicServer


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| qryNodeInfo | [QryNodeInfoRootReq](#ironicServer.QryNodeInfoRootReq) | [QryNodeInfoRootRes](#ironicServer.QryNodeInfoRootRes) | 裸金属节点状态查询 |
| createNodes | [CreateNodesRootReq](#ironicServer.CreateNodesRootReq) | [CreateNodesRootRes](#ironicServer.CreateNodesRootRes) | 创建裸金属节点 |
| installNodeSys | [InstallNodeSysRootReq](#ironicServer.InstallNodeSysRootReq) | [InstallNodeSysRootRes](#ironicServer.InstallNodeSysRootRes) | 安装裸金属系统 |
| operNodePower | [OperNodePowerRootReq](#ironicServer.OperNodePowerRootReq) | [OperNodePowerRootRes](#ironicServer.OperNodePowerRootRes) | 操作裸金属实例电源 |

 



<a name="playbook/playbook.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## playbook/playbook.proto



<a name="playbook.CreateRequest"></a>

### CreateRequest
创建项目请求内容体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | playbook项目名称 |
| url | [string](#string) |  | playbook存储在对象存储的url |
| url_type | [UrlType](#playbook.UrlType) |  | 枚举类型 1-文件下载 2-git下载 |
| entrypoint | [string](#string) | repeated | playbook入口yml文件，有多个 |
| description | [string](#string) |  | playbook项目描述 |
| md5 | [string](#string) |  | 文件md5 |
| size | [int64](#int64) |  | 文件大小 |






<a name="playbook.CreateResponse"></a>

### CreateResponse
创建项目请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playbook_id | [int32](#int32) |  | playbook项目ID |
| created | [string](#string) |  | playbook项目创建时间 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="playbook.FilterRequest"></a>

### FilterRequest
筛选项目请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| keyword | [string](#string) |  | 用于模糊筛选的playbook项目名称的关键字 |






<a name="playbook.FilterResponse"></a>

### FilterResponse
筛选项目请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projects | [ProjectObject](#playbook.ProjectObject) | repeated | 返回的筛选到的多个项目实例 |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="playbook.GetRequest"></a>

### GetRequest
根据项目ID获取记录


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playbook_id | [int32](#int32) |  | playbook项目ID |






<a name="playbook.GetResponse"></a>

### GetResponse
获取项目请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project | [ProjectObject](#playbook.ProjectObject) |  | 返回的项目实例 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="playbook.PlaybookEntrypointObject"></a>

### PlaybookEntrypointObject
playbook入口文件实例


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playbook_entrypoint_id | [int32](#int32) |  | 入口文件ID |
| name | [string](#string) |  | 入口文件名称 |






<a name="playbook.ProjectObject"></a>

### ProjectObject
项目实例


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playbook_id | [int32](#int32) |  | playbook项目ID |
| name | [string](#string) |  | playbook项目名称 |
| description | [string](#string) |  | playbook项目描述 |
| url | [string](#string) |  | playbook存储在对象存储的url |
| url_type | [UrlType](#playbook.UrlType) |  | 枚举类型 1-文件下载 2-git下载 |
| version | [string](#string) |  | playbook文件的版本号 |
| entrypoint | [PlaybookEntrypointObject](#playbook.PlaybookEntrypointObject) | repeated | playbook入口yml文件，有多个 |
| md5 | [string](#string) |  | 文件md5 |
| size | [int64](#int64) |  | 文件大小 |
| created | [string](#string) |  | 创建时间 |
| updated | [string](#string) |  | 更新时间 |
| playbook_file_id | [int32](#int32) |  | playbook项目文件ID |
| creator | [common.User](#common.User) |  | 创建人 |






<a name="playbook.SaveVersionRequest"></a>

### SaveVersionRequest
版本列表请求内容体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playbook_id | [int32](#int32) |  | playbook项目ID |
| url | [string](#string) |  | playbook存储在对象存储的url |
| url_type | [UrlType](#playbook.UrlType) |  | 枚举类型 1-文件下载 2-git下载 |
| version | [string](#string) |  | playbook文件的版本号 |
| md5 | [string](#string) |  | 文件md5 |
| size | [int64](#int64) |  | 文件大小 |
| entrypoint | [string](#string) | repeated | playbook入口yml文件，有多个 |






<a name="playbook.SaveVersionResponse"></a>

### SaveVersionResponse
版本列表请求响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version_id | [int32](#int32) |  | 版本信息 |
| created | [string](#string) |  | playbook项目创建时间 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="playbook.UpdateRequest"></a>

### UpdateRequest
更新项目请求内容体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playbook_id | [int32](#int32) |  | playbook项目ID |
| name | [string](#string) |  | playbook项目名称 |
| description | [string](#string) |  | playbook项目描述 |
| url | [string](#string) |  | playbook存储在对象存储的url |
| url_type | [UrlType](#playbook.UrlType) |  | 枚举类型 1-文件下载 2-git下载 |
| version | [string](#string) |  | playbook文件的版本号 |
| entrypoint | [string](#string) | repeated | playbook入口yml文件，有多个 |
| md5 | [string](#string) |  | 文件md5 |
| size | [int64](#int64) |  | 文件大小 |






<a name="playbook.UpdateResponse"></a>

### UpdateResponse
更新项目请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [string](#string) |  | playbook项目更新时间 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |





 


<a name="playbook.UrlType"></a>

### UrlType
枚举URL类型

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 | 未定义 |
| FILE | 1 | 文件下载 |
| GIT | 2 | git下载 |


 

 


<a name="playbook.Playbook"></a>

### Playbook


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#playbook.CreateRequest) | [CreateResponse](#playbook.CreateResponse) | 创建playbook项目 |
| Filter | [FilterRequest](#playbook.FilterRequest) | [FilterResponse](#playbook.FilterResponse) | 筛选playbook项目 |
| Get | [GetRequest](#playbook.GetRequest) | [GetResponse](#playbook.GetResponse) | 获取playbook项目 |
| Update | [UpdateRequest](#playbook.UpdateRequest) | [UpdateResponse](#playbook.UpdateResponse) | 更新playbook项目 |
| SaveVersion | [SaveVersionRequest](#playbook.SaveVersionRequest) | [SaveVersionResponse](#playbook.SaveVersionResponse) | 保存 playbook 版本 |

 



<a name="schedule/schedule.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## schedule/schedule.proto



<a name="schedule.CreateRequest"></a>

### CreateRequest
创建定时任务请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedules | [CreateScheduleObject](#schedule.CreateScheduleObject) | repeated | 创建定时任务的请求内容体 |
| cmdb_search_request | [cmdb.ChooseHostRequest](#cmdb.ChooseHostRequest) | repeated | cmdb的搜索条件 |
| extra_var | [string](#string) |  | 额外变量JSON String 例如： {&#34;key&#34;:&#34;testKey&#34;,&#34;value&#34;:&#34;testVal&#34;,&#34;description&#34;:&#34;测试描述&#34;} |






<a name="schedule.CreateResponse"></a>

### CreateResponse
创建定时任务请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedules | [CreateScheduleResponseObject](#schedule.CreateScheduleResponseObject) | repeated | 创建定时任务的请求返回内容体 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回请求状态 |






<a name="schedule.CreateScheduleObject"></a>

### CreateScheduleObject
创建定时任务的请求内容体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_id | [int32](#int32) |  | 模板ID |
| name | [string](#string) |  | 定时任务名 |
| start_time | [string](#string) |  | 定时任务开始时间 |
| end_time | [string](#string) |  | 定时任务结束时间 |
| cron_expression | [string](#string) |  | 定时任务表达式 |
| task_type | [TaskType](#schedule.TaskType) |  | 任务类型 |






<a name="schedule.CreateScheduleResponseObject"></a>

### CreateScheduleResponseObject
创建定时任务的请求返回内容体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule_id | [int32](#int32) |  | 定时任务ID |
| created | [string](#string) |  | 定时任务创建时间 |






<a name="schedule.FilterRequest"></a>

### FilterRequest
筛选定时任务请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| schedule_name | [string](#string) |  | 用于筛选的定时任务名字的关键字 |
| task_type | [TaskType](#schedule.TaskType) |  | 用于筛选的定时任务的任务类型 |
| task_id | [int32](#int32) |  | task_type对应的id |






<a name="schedule.FilterResponse"></a>

### FilterResponse
筛选定时任务请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedules | [ScheduleObject](#schedule.ScheduleObject) | repeated | 筛选到的多个定时任务实例 |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="schedule.GetRequest"></a>

### GetRequest
获取定时任务请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule_id | [int32](#int32) |  | 定时任务ID |






<a name="schedule.GetResponse"></a>

### GetResponse
获取定时任务请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule | [ScheduleObject](#schedule.ScheduleObject) |  | 定时任务实例 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="schedule.ScheduleObject"></a>

### ScheduleObject
定时任务实例


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule_id | [int32](#int32) |  | 定时任务ID |
| template | [template.TemplateObject](#template.TemplateObject) |  | 模板实例 |
| name | [string](#string) |  | 定时任务名 |
| start_time | [string](#string) |  | 开始日期 |
| end_time | [string](#string) |  | 开始时间 |
| cron_expression | [string](#string) |  | 定时任务表达式 |
| status | [ScheduleStatus](#schedule.ScheduleStatus) |  | 任务状态 |
| task_type | [TaskType](#schedule.TaskType) |  | 任务类型 |
| execute_count | [int32](#int32) |  | 执行总数 |
| next_time | [string](#string) |  | 下次执行时间 |






<a name="schedule.SwitchStatusRequest"></a>

### SwitchStatusRequest
切换定时任务状态请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule_id | [int32](#int32) |  | 定时任务ID |
| schedule_status | [ScheduleStatus](#schedule.ScheduleStatus) |  | 定时任务状态 |






<a name="schedule.SwitchStatusResponse"></a>

### SwitchStatusResponse
切换定时任务状态请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="schedule.UpdateRequest"></a>

### UpdateRequest
更新定时任务请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule_id | [int32](#int32) |  | 定时任务ID |
| name | [string](#string) |  | 定时任务名字 |
| start_time | [string](#string) |  | 定时任务开启时间 |
| end_time | [string](#string) |  | 定时任务结束时间 |
| cron_expression | [string](#string) |  | 定时任务表达式 |






<a name="schedule.UpdateResponse"></a>

### UpdateResponse
更新请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [string](#string) |  | 定时任务更新时间 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |





 


<a name="schedule.ScheduleStatus"></a>

### ScheduleStatus
任务状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| OFF | 0 |  |
| ON | 1 |  |



<a name="schedule.TaskType"></a>

### TaskType
任务类型 0-undefined 1-作业模板；2-容器部署 3-裸金属管理

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 | 0-undefined |
| PLAYBOOK | 1 | 1-作业模板 |
| CONTAINER | 2 | 2-容器部署 |
| BARE_METAL | 3 | 3-裸金属管理 |
| SERVER_COMPARE | 4 | 4-交维表导入后巡检任务 |


 

 


<a name="schedule.Schedule"></a>

### Schedule
定时任务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#schedule.CreateRequest) | [CreateResponse](#schedule.CreateResponse) | 创建定时任务 |
| Filter | [FilterRequest](#schedule.FilterRequest) | [FilterResponse](#schedule.FilterResponse) | 筛选定时任务 |
| Get | [GetRequest](#schedule.GetRequest) | [GetResponse](#schedule.GetResponse) | 获取定时任务 |
| Update | [UpdateRequest](#schedule.UpdateRequest) | [UpdateResponse](#schedule.UpdateResponse) | 更新定时任务 |
| SwitchStatus | [SwitchStatusRequest](#schedule.SwitchStatusRequest) | [SwitchStatusResponse](#schedule.SwitchStatusResponse) | 切换状态 |

 



<a name="subtask/subtask.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## subtask/subtask.proto



<a name="subtask.CompleteRequest"></a>

### CompleteRequest
完成子任务实例请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 子任务实例ID |
| start_time | [string](#string) |  | 执行开始时间 |
| end_time | [string](#string) |  | 执行结束时间 |
| result | [SubTaskResult](#subtask.SubTaskResult) |  | 子任务执行结果 |
| execute_count | [int32](#int32) |  | 总共执行多少主机 |
| fail_count | [int32](#int32) |  | 执行失败多少台主机 |
| success_count | [int32](#int32) |  | 执行成功多少台主机 |






<a name="subtask.CompleteResponse"></a>

### CompleteResponse
完成子任务实例请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="subtask.CreateRequest"></a>

### CreateRequest
创建子任务实例请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [int64](#int64) |  |  |
| cmdb_search_request | [cmdb.ChooseHostRequest](#cmdb.ChooseHostRequest) |  | cmdb的搜索条件 |
| template_id | [int32](#int32) |  | 模板ID |






<a name="subtask.CreateResponse"></a>

### CreateResponse
创建子任务实例请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) | repeated | 子任务实例ID |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="subtask.CreateServerCompareRequest"></a>

### CreateServerCompareRequest
交维表导入后巡检子任务请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [int64](#int64) |  | 任务实例ID |
| cmdb_search_request | [cmdb.ChooseServerCompareRequest](#cmdb.ChooseServerCompareRequest) |  | cmdb的搜索条件 |






<a name="subtask.CreateServerCompareResponse"></a>

### CreateServerCompareResponse
交维表导入后巡检子任务请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 子任务实例ID |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="subtask.CreateServerRequest"></a>

### CreateServerRequest
裸金属子任务实例请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [int64](#int64) |  | 任务实例ID |
| task_type | [ServerTaskType](#subtask.ServerTaskType) |  | 裸金属任务类型 |
| power_state | [cmdb.ServerPowerState](#cmdb.ServerPowerState) |  | 裸金属电源状态 |
| cmdb_search_request | [cmdb.ChooseServerRequest](#cmdb.ChooseServerRequest) |  | cmdb的搜索条件 |
| image_file_url | [string](#string) |  | 安装镜像文件的URL |
| image_file_md5 | [string](#string) |  | 安装镜像文件的MD5 |






<a name="subtask.CreateServerResponse"></a>

### CreateServerResponse
裸金属子任务实例返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 子任务实例ID |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |





 


<a name="subtask.ServerTaskType"></a>

### ServerTaskType
裸金属任务状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 | 未定义任务类型 |
| CHECK | 1 | 裸金属状态查询任务 |
| POWER | 2 | 裸金属电源开关任务 |
| CREATE | 3 | 裸金属初始化创建uuid任务 |
| INSTALL | 4 | 裸金属安装任务 |



<a name="subtask.SubTaskResult"></a>

### SubTaskResult
子任务执行结果

| Name | Number | Description |
| ---- | ------ | ----------- |
| FAILURE | 0 | 失败 |
| SUCCESS | 1 | 成功 |


 

 


<a name="subtask.SubTask"></a>

### SubTask
子任务实例

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#subtask.CreateRequest) | [CreateResponse](#subtask.CreateResponse) | 创建作业子任务 |
| Complete | [CompleteRequest](#subtask.CompleteRequest) | [CompleteResponse](#subtask.CompleteResponse) | 完成作业子任务 |
| CreateServer | [CreateServerRequest](#subtask.CreateServerRequest) | [CreateServerResponse](#subtask.CreateServerResponse) | 创建裸金属子任务 |
| CreateServerCompare | [CreateServerCompareRequest](#subtask.CreateServerCompareRequest) | [CreateServerCompareResponse](#subtask.CreateServerCompareResponse) | 交维表导入后巡检子任务 |

 



<a name="task/queue.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## task/queue.proto



<a name="task.BareMetalCreateTask"></a>

### BareMetalCreateTask
裸金属创建任务


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 子任务实例ID |
| cmdb_search_request | [cmdb.ChooseServerRequest](#cmdb.ChooseServerRequest) |  | cmdb的搜索条件 |






<a name="task.BareMetalInstallTask"></a>

### BareMetalInstallTask
裸金属安装任务


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 子任务实例ID |
| cmdb_search_request | [cmdb.ChooseServerRequest](#cmdb.ChooseServerRequest) |  | cmdb的搜索条件 |
| image_file_url | [string](#string) |  | 安装镜像文件的URL |
| image_file_md5 | [string](#string) |  | 安装镜像文件的MD5 |
| extra_var | [string](#string) |  | 额外变量 |






<a name="task.BareMetalPowerTask"></a>

### BareMetalPowerTask
裸金属电源开关任务


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 子任务实例ID |
| state | [cmdb.ServerPowerState](#cmdb.ServerPowerState) |  | 电源开关操作 |
| cmdb_search_request | [cmdb.ChooseServerRequest](#cmdb.ChooseServerRequest) |  | cmdb的搜索条件 |






<a name="task.BareMetalSearchTask"></a>

### BareMetalSearchTask
裸金属状态查询任务


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 子任务实例ID |
| cmdb_search_request | [cmdb.ChooseServerRequest](#cmdb.ChooseServerRequest) |  | cmdb的搜索条件 |






<a name="task.ServerCompareTask"></a>

### ServerCompareTask
交维表导入后巡检任务


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 任务实例ID |
| cmdb_search_request | [cmdb.ChooseServerCompareRequest](#cmdb.ChooseServerCompareRequest) |  | 查询条件 |
| playbook_file_url | [string](#string) |  | playbook 文件的URL |
| playbook_file_md5 | [string](#string) |  | playbook 文件的MD5 |
| playbook_yml_name | [string](#string) |  | playbook 入口文件 |






<a name="task.TemplateExecuteTask"></a>

### TemplateExecuteTask
作业模板执行任务


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 子任务实例ID |
| cmdb_search_request | [cmdb.ChooseHostRequest](#cmdb.ChooseHostRequest) |  | cmdb的搜索条件 |
| playbook_file_url | [string](#string) |  | playbook 文件的URL |
| playbook_file_md5 | [string](#string) |  | playbook 文件的MD5 |
| playbook_yml_name | [string](#string) |  | playbook 入口文件 |
| extra_var | [string](#string) |  | 额外变量 |





 

 

 

 



<a name="task/task.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## task/task.proto



<a name="task.CheckServerStateRequest"></a>

### CheckServerStateRequest
调用裸金属应用检查裸金属状态请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cmdb_search_request | [cmdb.ChooseServerRequest](#cmdb.ChooseServerRequest) | repeated | 机房及对应物理机 |






<a name="task.CheckServerStateResponse"></a>

### CheckServerStateResponse
调用裸金属应用检查裸金属状态返回体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="task.CreateRequest"></a>

### CreateRequest
创建任务实例请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_id | [int32](#int32) |  | 模板ID |
| cmdb_search_request | [cmdb.ChooseHostRequest](#cmdb.ChooseHostRequest) | repeated | cmdb的搜索条件 |
| task_type | [schedule.TaskType](#schedule.TaskType) |  | 任务类型 |
| extra_var | [string](#string) |  | 额外变量JSON String 例如： {&#34;key&#34;:&#34;testKey&#34;,&#34;value&#34;:&#34;testVal&#34;,&#34;description&#34;:&#34;测试描述&#34;} |
| schedule_id | [int32](#int32) |  | 定时任务ID |
| schedule_type | [ScheduleType](#task.ScheduleType) |  | 定时任务类型 |
| name | [string](#string) |  | 作业名称 |






<a name="task.CreateResponse"></a>

### CreateResponse
创建任务实例请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [int64](#int64) |  | 任务实例ID |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="task.CreateServerCompareRequest"></a>

### CreateServerCompareRequest
交维表导入后巡检任务请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int32](#int32) | repeated | 机房ID列表 |






<a name="task.CreateServerCompareResponse"></a>

### CreateServerCompareResponse
交维表导入后巡检任务返回体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="task.CreateServerRequest"></a>

### CreateServerRequest
调用裸金属创建操作请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cmdb_search_request | [cmdb.ChooseServerRequest](#cmdb.ChooseServerRequest) | repeated | 机房及对应物理机 |






<a name="task.CreateServerResponse"></a>

### CreateServerResponse
调用裸金属创建操作返回体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="task.FilterRequest"></a>

### FilterRequest
筛选获取作业任务请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| task_name | [string](#string) |  | 用于筛选任务名字的关键字 |
| task_type | [schedule.TaskType](#schedule.TaskType) |  | 用于筛选定时任务的任务类型 0-undefined 1-作业模板；2-容器部署；3-裸金属管理 |






<a name="task.FilterResponse"></a>

### FilterResponse
筛选作业任务请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tasks | [TaskObject](#task.TaskObject) | repeated | 筛选到的多个作业任务实例 |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="task.GetLogRequest"></a>

### GetLogRequest
获取作业任务详细执行过程请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [int64](#int64) |  | 作业任务ID |






<a name="task.GetLogResponse"></a>

### GetLogResponse
获取作业任务详细执行过程请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="task.GetRequest"></a>

### GetRequest
获取作业任务请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [int64](#int64) |  | 作业任务ID |






<a name="task.GetResponse"></a>

### GetResponse
获取作业任务请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [TaskObject](#task.TaskObject) |  | 获取的作业任务实例 |
| playbook_version | [string](#string) |  | task对于的作业版本 |
| template_name | [string](#string) |  | 作业对应的模板名字 |
| host_module | [string](#string) |  | task对于的主机模块 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="task.InstallServerRequest"></a>

### InstallServerRequest
调用裸金属安装操作请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cmdb_search_request | [cmdb.ChooseServerRequest](#cmdb.ChooseServerRequest) | repeated | 机房及对应物理机 |
| image_file_url | [string](#string) |  | 安装镜像文件的URL |
| image_file_md5 | [string](#string) |  | 安装镜像文件的MD5 |
| cmdb_host_search_request | [cmdb.ChooseHostRequest](#cmdb.ChooseHostRequest) |  | 主机查询条件 |






<a name="task.InstallServerResponse"></a>

### InstallServerResponse
调用裸金属安装操作返回体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="task.RetryRequest"></a>

### RetryRequest
调用重新执行作业任务操作请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [int64](#int64) |  | 作业任务ID |






<a name="task.RetryResponse"></a>

### RetryResponse
调用重新执行作业任务操作返回体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="task.ServerPowerControlRequest"></a>

### ServerPowerControlRequest
调用裸金属应用执行开关机操作请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [cmdb.ServerPowerState](#cmdb.ServerPowerState) |  | 电源开关操作 |
| cmdb_search_request | [cmdb.ChooseServerRequest](#cmdb.ChooseServerRequest) | repeated | 机房及对应物理机 |






<a name="task.ServerPowerControlResponse"></a>

### ServerPowerControlResponse
调用裸金属应用执行开关机操作返回体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






<a name="task.TaskObject"></a>

### TaskObject
作业任务实例


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [int64](#int64) |  | 作业任务ID |
| template_id | [int32](#int32) |  | 模板ID |
| schedule_id | [int32](#int32) |  | 定时任务ID |
| schedule_type | [ScheduleType](#task.ScheduleType) |  | 定时任务类型 |
| task_type | [schedule.TaskType](#schedule.TaskType) |  | 任务类型 |
| cmdb_search_request | [cmdb.ChooseHostRequest](#cmdb.ChooseHostRequest) | repeated | cmdb的搜索条件 |
| extra_var | [string](#string) |  | 额外变量JSON String 例如： {&#34;key&#34;:&#34;testKey&#34;,&#34;value&#34;:&#34;testVal&#34;,&#34;description&#34;:&#34;测试描述&#34;} |
| executor | [int32](#int32) |  | 执行人ID |
| execute_count | [int32](#int32) |  | 总共执行多少主机 |
| fail_count | [int32](#int32) |  | 执行失败多少台主机 |
| success_count | [int32](#int32) |  | 执行成功多少台主机 |
| start_time | [string](#string) |  | 执行开始时间 |
| end_time | [string](#string) |  | 执行结束时间 |
| name | [string](#string) |  | 作业名称 |





 


<a name="task.ScheduleType"></a>

### ScheduleType
定时任务类型 0-undefined 1-单次任务 2-定时任务

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 | 0-undefined |
| SINGLE | 1 | 1-单次任务 |
| CRONTAB | 2 | 2-定时任务 |


 

 


<a name="task.Task"></a>

### Task
任务实例

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#task.CreateRequest) | [CreateResponse](#task.CreateResponse) | 创建作业任务(执行作业模板) |
| Filter | [FilterRequest](#task.FilterRequest) | [FilterResponse](#task.FilterResponse) | 筛选作业任务 |
| Get | [GetRequest](#task.GetRequest) | [GetResponse](#task.GetResponse) | 获取作业任务 |
| GetLog | [GetLogRequest](#task.GetLogRequest) | [GetLogResponse](#task.GetLogResponse) | 获取作业任务详细执行过程

option (google.api.http) = { get: &#34;/v1/task/{task_id}/logs&#34; }; |
| CreateServerCompare | [CreateServerCompareRequest](#task.CreateServerCompareRequest) | [CreateServerCompareResponse](#task.CreateServerCompareResponse) | 交维表导入后巡检任务 |
| CheckServerState | [CheckServerStateRequest](#task.CheckServerStateRequest) | [CheckServerStateResponse](#task.CheckServerStateResponse) | 调用裸金属应用检查裸金属状态 |
| ServerPowerControl | [ServerPowerControlRequest](#task.ServerPowerControlRequest) | [ServerPowerControlResponse](#task.ServerPowerControlResponse) | 调用裸金属应用执行开关机操作 |
| CreateServer | [CreateServerRequest](#task.CreateServerRequest) | [CreateServerResponse](#task.CreateServerResponse) | 调用裸金属创建操作 |
| InstallServer | [InstallServerRequest](#task.InstallServerRequest) | [InstallServerResponse](#task.InstallServerResponse) | 调用裸金属安装操作 |
| Retry | [RetryRequest](#task.RetryRequest) | [RetryResponse](#task.RetryResponse) | 重新执行作业任务 |

 



<a name="template/template.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## template/template.proto



<a name="template.CreateRequest"></a>

### CreateRequest
创建模板请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 模板名 |
| playbook_id | [int32](#int32) |  | playbook 项目id |
| playbook_file_id | [int32](#int32) |  | 选择的playbook项目对应的playbook_file |
| playbook_entrypoint_id | [int32](#int32) |  | 选择的playbook入口yml文件 |
| description | [string](#string) |  | 模板描述 |
| extra_var | [string](#string) |  | 额外变量JSON String 例如： {&#34;key&#34;:&#34;testKey&#34;,&#34;value&#34;:&#34;testVal&#34;,&#34;description&#34;:&#34;测试描述&#34;} |






<a name="template.CreateResponse"></a>

### CreateResponse
创建模板请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_id | [int32](#int32) |  | 模板ID |
| created | [string](#string) |  | 模板创建的时间 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 请求返回状态 |






<a name="template.DeleteRequest"></a>

### DeleteRequest
删除模板请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_id | [int32](#int32) |  | 模板ID |






<a name="template.DeleteResponse"></a>

### DeleteResponse
删除请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 请求返回的状态 |






<a name="template.FilterRequest"></a>

### FilterRequest
筛选模板请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| name | [string](#string) |  | 模板名的一部分，模糊查询时的关键字 |
| template_id | [string](#string) |  | 模板ID |






<a name="template.FilterResponse"></a>

### FilterResponse
筛选模板请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| templates | [TemplateObject](#template.TemplateObject) | repeated | 请求返回的模板 |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 请求返回状态 |






<a name="template.GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_id | [int32](#int32) |  | 模板ID |






<a name="template.GetResponse"></a>

### GetResponse
获取模板请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template | [TemplateObject](#template.TemplateObject) |  | 请求返回的模板 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 请求返回状态 |






<a name="template.TemplateObject"></a>

### TemplateObject
模板实例


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_id | [int32](#int32) |  | 模板ID |
| name | [string](#string) |  | 模板名 |
| description | [string](#string) |  | 模板描述 |
| playbook | [playbook.ProjectObject](#playbook.ProjectObject) |  | playbook 项目 |
| playbook_file_id | [int32](#int32) |  | 选择的playbook项目对应的playbook_file |
| playbook_entrypoint_id | [int32](#int32) |  | 选择的playbook入口yml文件 |
| extra_var | [string](#string) |  | 额外变量 |
| creator | [common.User](#common.User) |  | 创建人 |
| created | [string](#string) |  | 创建时间 |
| updated | [string](#string) |  | 更新时间 |
| state | [TemplateState](#template.TemplateState) |  | 模板状态开关标识 |






<a name="template.UpdateRequest"></a>

### UpdateRequest
更新模板请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_id | [int32](#int32) |  | 模板ID |
| name | [string](#string) |  | 模板名 |
| description | [string](#string) |  | 模板描述 |
| extra_var | [string](#string) |  | 额外变量JSON String 例如： {&#34;key&#34;:&#34;testKey&#34;,&#34;value&#34;:&#34;testVal&#34;,&#34;description&#34;:&#34;测试描述&#34;} |






<a name="template.UpdateResponse"></a>

### UpdateResponse
更新请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [string](#string) |  | 模板更新的时间 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 请求返回的状态 |






<a name="template.UpdateStateRequest"></a>

### UpdateStateRequest
更新模板状态开关请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_id | [int32](#int32) |  | 模板ID |
| state | [TemplateState](#template.TemplateState) |  | 模板状态开关标识 |






<a name="template.UpdateStateResponse"></a>

### UpdateStateResponse
更新模板状态请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 请求返回的状态 |





 


<a name="template.TemplateState"></a>

### TemplateState
模板状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| OFF | 0 | 没运行 |
| ON | 1 | 运行 |


 

 


<a name="template.Template"></a>

### Template
作业模板

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#template.CreateRequest) | [CreateResponse](#template.CreateResponse) | 创建作业模板 |
| Filter | [FilterRequest](#template.FilterRequest) | [FilterResponse](#template.FilterResponse) | 筛选作业模板 |
| Get | [GetRequest](#template.GetRequest) | [GetResponse](#template.GetResponse) | 获取作业模板 |
| Update | [UpdateRequest](#template.UpdateRequest) | [UpdateResponse](#template.UpdateResponse) | 更新作业模板 |
| UpdateState | [UpdateStateRequest](#template.UpdateStateRequest) | [UpdateStateResponse](#template.UpdateStateResponse) | 更新作业模板状态开关 |
| Delete | [DeleteRequest](#template.DeleteRequest) | [DeleteResponse](#template.DeleteResponse) | 删除作业模板

option (google.api.http) = { post: &#34;/v1/template/delete/{template_id}&#34; }; |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

