# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [auth/auth.proto](#auth/auth.proto)
    - [CallbackRequest](#auth.CallbackRequest)
    - [CallbackResponse](#auth.CallbackResponse)
    - [CheckRequest](#auth.CheckRequest)
    - [CheckResponse](#auth.CheckResponse)
    - [GetActionListRequest](#auth.GetActionListRequest)
    - [GetActionListResponse](#auth.GetActionListResponse)
    - [GetResourceListRequest](#auth.GetResourceListRequest)
    - [GetResourceListResponse](#auth.GetResourceListResponse)
    - [GetUserInfoRequest](#auth.GetUserInfoRequest)
    - [GetUserInfoResponse](#auth.GetUserInfoResponse)
    - [LoginRequest](#auth.LoginRequest)
    - [LoginResponse](#auth.LoginResponse)
    - [LogoutRequest](#auth.LogoutRequest)
    - [LogoutResponse](#auth.LogoutResponse)
    - [ResourceActionData](#auth.ResourceActionData)
    - [ResourceActionItem](#auth.ResourceActionItem)
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
    - [HostCloudInfo](#cmdb.HostCloudInfo)
    - [HostInfoObject](#cmdb.HostInfoObject)
    - [HostObject](#cmdb.HostObject)
    - [ImportAssetRequest](#cmdb.ImportAssetRequest)
    - [ImportAssetResponse](#cmdb.ImportAssetResponse)
    - [ImportCrossTableRequest](#cmdb.ImportCrossTableRequest)
    - [ImportCrossTableResponse](#cmdb.ImportCrossTableResponse)
    - [ImportDetailRequest](#cmdb.ImportDetailRequest)
    - [ImportDetailResponse](#cmdb.ImportDetailResponse)
    - [ImportDetailResponse.InstInfoEntry](#cmdb.ImportDetailResponse.InstInfoEntry)
    - [ImportHistoryObject](#cmdb.ImportHistoryObject)
    - [ImportHistoryRequest](#cmdb.ImportHistoryRequest)
    - [ImportHistoryResponse](#cmdb.ImportHistoryResponse)
    - [ImportHostRequest](#cmdb.ImportHostRequest)
    - [ImportHostResponse](#cmdb.ImportHostResponse)
    - [ImportLakeRequest](#cmdb.ImportLakeRequest)
    - [ImportLakeResponse](#cmdb.ImportLakeResponse)
    - [ImportReviewRequest](#cmdb.ImportReviewRequest)
    - [ImportReviewResponse](#cmdb.ImportReviewResponse)
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
    - [LevelHost](#cmdb.LevelHost)
    - [ModuleHost](#cmdb.ModuleHost)
    - [ModuleObject](#cmdb.ModuleObject)
    - [RoomObject](#cmdb.RoomObject)
    - [RoomTopologyRequest](#cmdb.RoomTopologyRequest)
    - [RoomTopologyResponse](#cmdb.RoomTopologyResponse)
    - [SearchHostRequest](#cmdb.SearchHostRequest)
    - [SearchHostResponse](#cmdb.SearchHostResponse)
    - [SearchInstRequest](#cmdb.SearchInstRequest)
    - [SearchInstResponse](#cmdb.SearchInstResponse)
    - [SearchLakeAreaRequest](#cmdb.SearchLakeAreaRequest)
    - [SearchLakeAreaResponse](#cmdb.SearchLakeAreaResponse)
    - [SearchLakeHostRequest](#cmdb.SearchLakeHostRequest)
    - [SearchLakeHostResponse](#cmdb.SearchLakeHostResponse)
    - [SearchLakeRequest](#cmdb.SearchLakeRequest)
    - [SearchLakeResponse](#cmdb.SearchLakeResponse)
    - [SearchLevelHostRequest](#cmdb.SearchLevelHostRequest)
    - [SearchLevelHostResponse](#cmdb.SearchLevelHostResponse)
    - [SearchLevelHostResponse.ResultEntry](#cmdb.SearchLevelHostResponse.ResultEntry)
    - [SearchModuleHostRequest](#cmdb.SearchModuleHostRequest)
    - [SearchModuleHostResponse](#cmdb.SearchModuleHostResponse)
    - [SearchModuleListRequest](#cmdb.SearchModuleListRequest)
    - [SearchModuleListResponse](#cmdb.SearchModuleListResponse)
    - [SearchModuleRequest](#cmdb.SearchModuleRequest)
    - [SearchModuleResponse](#cmdb.SearchModuleResponse)
    - [Server](#cmdb.Server)
    - [ServerListRequest](#cmdb.ServerListRequest)
    - [ServerListResponse](#cmdb.ServerListResponse)
    - [ServerObject](#cmdb.ServerObject)
    - [ServerRoomObject](#cmdb.ServerRoomObject)
    - [SetObject](#cmdb.SetObject)
    - [TopologyObject](#cmdb.TopologyObject)
    - [UpdateHostStateRequest](#cmdb.UpdateHostStateRequest)
    - [UpdateHostStateResponse](#cmdb.UpdateHostStateResponse)
    - [UpdateInstRequest](#cmdb.UpdateInstRequest)
    - [UpdateInstResponse](#cmdb.UpdateInstResponse)
    - [UpdateLakeStateRequest](#cmdb.UpdateLakeStateRequest)
    - [UpdateLakeStateResponse](#cmdb.UpdateLakeStateResponse)
    - [UpdateVipStateRequest](#cmdb.UpdateVipStateRequest)
    - [UpdateVipStateResponse](#cmdb.UpdateVipStateResponse)
    - [VipObject](#cmdb.VipObject)
    - [ZoneObject](#cmdb.ZoneObject)
  
    - [AreaLevel](#cmdb.AreaLevel)
    - [ImportStatus](#cmdb.ImportStatus)
    - [ImportType](#cmdb.ImportType)
    - [InstType](#cmdb.InstType)
    - [LakeNodeLevel](#cmdb.LakeNodeLevel)
    - [OnlineState](#cmdb.OnlineState)
    - [ServerInstallState](#cmdb.ServerInstallState)
    - [ServerPowerState](#cmdb.ServerPowerState)
  
  
    - [Cmdb](#cmdb.Cmdb)
  

- [common/api.proto](#common/api.proto)
    - [Condition](#common.Condition)
    - [MongoCondition](#common.MongoCondition)
    - [Paging](#common.Paging)
    - [ResponseStatus](#common.ResponseStatus)
  
    - [StatusCode](#common.StatusCode)
  
  
  

- [common/user.proto](#common/user.proto)
    - [User](#common.User)
  
  
  
  

- [file/file.proto](#file/file.proto)
    - [DownloadTemplateRequest](#file.DownloadTemplateRequest)
    - [DownloadTemplateResponse](#file.DownloadTemplateResponse)
    - [UploadPlaybookRequest](#file.UploadPlaybookRequest)
    - [UploadPlaybookResponse](#file.UploadPlaybookResponse)
    - [UploadRequest](#file.UploadRequest)
    - [UploadResponse](#file.UploadResponse)
  
    - [TemplateType](#file.TemplateType)
  
  
    - [File](#file.File)
  

- [greeter/greeter.proto](#greeter/greeter.proto)
    - [HelloReply](#greeter.HelloReply)
    - [HelloRequest](#greeter.HelloRequest)
  
  
  
    - [Greeter](#greeter.Greeter)
  

- [image/common.proto](#image/common.proto)
    - [Paging](#image.Paging)
    - [ResponseStatus](#image.ResponseStatus)
    - [TimeScope](#image.TimeScope)
  
    - [StatusCode](#image.StatusCode)
  
  
  

- [image/image.proto](#image/image.proto)
    - [ClusterInstance](#image.ClusterInstance)
    - [HealthRequest](#image.HealthRequest)
    - [HealthResponse](#image.HealthResponse)
    - [ImageAttr](#image.ImageAttr)
    - [ImageObject](#image.ImageObject)
    - [QueryRequest](#image.QueryRequest)
    - [QueryResponse](#image.QueryResponse)
  
    - [AppStatus](#image.AppStatus)
    - [ImageFormat](#image.ImageFormat)
    - [ImageType](#image.ImageType)
  
  
    - [Image](#image.Image)
  

- [ironic/ironicCom.proto](#ironic/ironicCom.proto)
    - [PageInfo](#ironicCom.PageInfo)
    - [ScopeInfo](#ironicCom.ScopeInfo)
    - [TcpContReq](#ironicCom.TcpContReq)
    - [TcpContRsp](#ironicCom.TcpContRsp)
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
    - [CreateNodesRootRsp](#ironicServer.CreateNodesRootRsp)
    - [CreateNodesRootRsp.ContractRootRsp](#ironicServer.CreateNodesRootRsp.ContractRootRsp)
    - [CreateNodesRootRsp.CreateNodeInfoRsp](#ironicServer.CreateNodesRootRsp.CreateNodeInfoRsp)
    - [CreateNodesRootRsp.NodeInfoRsp](#ironicServer.CreateNodesRootRsp.NodeInfoRsp)
    - [CreateNodesRootRsp.SvcContRsp](#ironicServer.CreateNodesRootRsp.SvcContRsp)
    - [InstallNodeSysRootReq](#ironicServer.InstallNodeSysRootReq)
    - [InstallNodeSysRootReq.ContractRootReq](#ironicServer.InstallNodeSysRootReq.ContractRootReq)
    - [InstallNodeSysRootReq.InstallNodeSysReq](#ironicServer.InstallNodeSysRootReq.InstallNodeSysReq)
    - [InstallNodeSysRootReq.NodeInstallInfo](#ironicServer.InstallNodeSysRootReq.NodeInstallInfo)
    - [InstallNodeSysRootReq.SvcContReq](#ironicServer.InstallNodeSysRootReq.SvcContReq)
    - [InstallNodeSysRootRsp](#ironicServer.InstallNodeSysRootRsp)
    - [InstallNodeSysRootRsp.ContractRootRsp](#ironicServer.InstallNodeSysRootRsp.ContractRootRsp)
    - [InstallNodeSysRootRsp.InstallNodeInfoRsp](#ironicServer.InstallNodeSysRootRsp.InstallNodeInfoRsp)
    - [InstallNodeSysRootRsp.NodeInfoRsp](#ironicServer.InstallNodeSysRootRsp.NodeInfoRsp)
    - [InstallNodeSysRootRsp.SvcContRsp](#ironicServer.InstallNodeSysRootRsp.SvcContRsp)
    - [OperNodePowerRootReq](#ironicServer.OperNodePowerRootReq)
    - [OperNodePowerRootReq.ContractRootReq](#ironicServer.OperNodePowerRootReq.ContractRootReq)
    - [OperNodePowerRootReq.NodePowerOper](#ironicServer.OperNodePowerRootReq.NodePowerOper)
    - [OperNodePowerRootReq.OperNodePowerReq](#ironicServer.OperNodePowerRootReq.OperNodePowerReq)
    - [OperNodePowerRootReq.SvcContReq](#ironicServer.OperNodePowerRootReq.SvcContReq)
    - [OperNodePowerRootRsp](#ironicServer.OperNodePowerRootRsp)
    - [OperNodePowerRootRsp.ContractRootRsp](#ironicServer.OperNodePowerRootRsp.ContractRootRsp)
    - [OperNodePowerRootRsp.NodePowerRsp](#ironicServer.OperNodePowerRootRsp.NodePowerRsp)
    - [OperNodePowerRootRsp.OperNodePowerRsp](#ironicServer.OperNodePowerRootRsp.OperNodePowerRsp)
    - [OperNodePowerRootRsp.SvcContRsp](#ironicServer.OperNodePowerRootRsp.SvcContRsp)
    - [QryHeartbeatRootReq](#ironicServer.QryHeartbeatRootReq)
    - [QryHeartbeatRootReq.ContractRootReq](#ironicServer.QryHeartbeatRootReq.ContractRootReq)
    - [QryHeartbeatRootReq.QryHeartbeatReq](#ironicServer.QryHeartbeatRootReq.QryHeartbeatReq)
    - [QryHeartbeatRootReq.SvcContReq](#ironicServer.QryHeartbeatRootReq.SvcContReq)
    - [QryHeartbeatRootRsp](#ironicServer.QryHeartbeatRootRsp)
    - [QryHeartbeatRootRsp.ClusterInstance](#ironicServer.QryHeartbeatRootRsp.ClusterInstance)
    - [QryHeartbeatRootRsp.ContractRootRsp](#ironicServer.QryHeartbeatRootRsp.ContractRootRsp)
    - [QryHeartbeatRootRsp.QryHeartbeatRsp](#ironicServer.QryHeartbeatRootRsp.QryHeartbeatRsp)
    - [QryHeartbeatRootRsp.SvcContRsp](#ironicServer.QryHeartbeatRootRsp.SvcContRsp)
    - [QryNodeInfoRootReq](#ironicServer.QryNodeInfoRootReq)
    - [QryNodeInfoRootReq.ContractRootReq](#ironicServer.QryNodeInfoRootReq.ContractRootReq)
    - [QryNodeInfoRootReq.QryNodeInfoReq](#ironicServer.QryNodeInfoRootReq.QryNodeInfoReq)
    - [QryNodeInfoRootReq.SvcContReq](#ironicServer.QryNodeInfoRootReq.SvcContReq)
    - [QryNodeInfoRootRsp](#ironicServer.QryNodeInfoRootRsp)
    - [QryNodeInfoRootRsp.ContractRootRsp](#ironicServer.QryNodeInfoRootRsp.ContractRootRsp)
    - [QryNodeInfoRootRsp.NodeInfo](#ironicServer.QryNodeInfoRootRsp.NodeInfo)
    - [QryNodeInfoRootRsp.QryNodeInfoRsp](#ironicServer.QryNodeInfoRootRsp.QryNodeInfoRsp)
    - [QryNodeInfoRootRsp.SvcContRsp](#ironicServer.QryNodeInfoRootRsp.SvcContRsp)
    - [ReInstallNodeSysRootReq](#ironicServer.ReInstallNodeSysRootReq)
    - [ReInstallNodeSysRootReq.ContractRootReq](#ironicServer.ReInstallNodeSysRootReq.ContractRootReq)
    - [ReInstallNodeSysRootReq.InstallNodeSysReq](#ironicServer.ReInstallNodeSysRootReq.InstallNodeSysReq)
    - [ReInstallNodeSysRootReq.NodeInstallInfo](#ironicServer.ReInstallNodeSysRootReq.NodeInstallInfo)
    - [ReInstallNodeSysRootReq.SvcContReq](#ironicServer.ReInstallNodeSysRootReq.SvcContReq)
    - [ReInstallNodeSysRootRsp](#ironicServer.ReInstallNodeSysRootRsp)
    - [ReInstallNodeSysRootRsp.ContractRootRsp](#ironicServer.ReInstallNodeSysRootRsp.ContractRootRsp)
    - [ReInstallNodeSysRootRsp.InstallNodeInfoRsp](#ironicServer.ReInstallNodeSysRootRsp.InstallNodeInfoRsp)
    - [ReInstallNodeSysRootRsp.ReNodeInfoRsp](#ironicServer.ReInstallNodeSysRootRsp.ReNodeInfoRsp)
    - [ReInstallNodeSysRootRsp.SvcContRsp](#ironicServer.ReInstallNodeSysRootRsp.SvcContRsp)
  
    - [QryHeartbeatRootRsp.AppStatus](#ironicServer.QryHeartbeatRootRsp.AppStatus)
    - [QryNodeInfoRootRsp.PowerStatus](#ironicServer.QryNodeInfoRootRsp.PowerStatus)
    - [QryNodeInfoRootRsp.Status](#ironicServer.QryNodeInfoRootRsp.Status)
  
  
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
    - [GetLogRequest](#subtask.GetLogRequest)
    - [GetLogResponse](#subtask.GetLogResponse)
  
    - [ServerTaskType](#subtask.ServerTaskType)
    - [SubTaskResult](#subtask.SubTaskResult)
    - [SubTaskState](#subtask.SubTaskState)
  
  
    - [SubTask](#subtask.SubTask)
  

- [task/queue.proto](#task/queue.proto)
    - [BareMetalCreateTask](#task.BareMetalCreateTask)
    - [BareMetalInstallTask](#task.BareMetalInstallTask)
    - [BareMetalPowerTask](#task.BareMetalPowerTask)
    - [BareMetalSearchTask](#task.BareMetalSearchTask)
    - [JobAgentLog](#task.JobAgentLog)
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
    - [GetSubTaskRequest](#task.GetSubTaskRequest)
    - [GetSubTaskResponse](#task.GetSubTaskResponse)
    - [InstallServerRequest](#task.InstallServerRequest)
    - [InstallServerResponse](#task.InstallServerResponse)
    - [RetryRequest](#task.RetryRequest)
    - [RetryResponse](#task.RetryResponse)
    - [ServerPowerControlRequest](#task.ServerPowerControlRequest)
    - [ServerPowerControlResponse](#task.ServerPowerControlResponse)
    - [SubTaskInfo](#task.SubTaskInfo)
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
  
    - [ReleaseCode](#template.ReleaseCode)
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






<a name="auth.GetActionListRequest"></a>

### GetActionListRequest
动作列表请求（命名适配IAM）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serviceId | [string](#string) |  | 服务ID |
| workspaceId | [string](#string) |  | 工作区ID |






<a name="auth.GetActionListResponse"></a>

### GetActionListResponse
动作列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [ResourceActionData](#auth.ResourceActionData) |  | 响应信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 响应状态 |






<a name="auth.GetResourceListRequest"></a>

### GetResourceListRequest
资源列表请求（命名适配IAM）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| serviceId | [string](#string) |  | 服务ID |
| workspaceId | [string](#string) |  | 工作区ID |
| serviceType | [string](#string) |  | 服务类型 |






<a name="auth.GetResourceListResponse"></a>

### GetResourceListResponse
资源列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [ResourceActionData](#auth.ResourceActionData) |  | 响应信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 响应状态 |






<a name="auth.GetUserInfoRequest"></a>

### GetUserInfoRequest
获取用户信息请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [bool](#bool) |  | 控制是否显示API访问授权Token，默认值false，即不显示 |






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






<a name="auth.ResourceActionData"></a>

### ResourceActionData
资源动作数据列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| list | [ResourceActionItem](#auth.ResourceActionItem) | repeated |  |






<a name="auth.ResourceActionItem"></a>

### ResourceActionItem
资源动作条目（命名适配IAM）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID |
| name | [string](#string) |  | 名称 |
| note | [string](#string) |  | 说明 |
| expr | [string](#string) |  | 表达式 |
| parentId | [int64](#int64) |  | 父节点ID |
| resourceType | [string](#string) |  | 资源类型 |






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
| access_token | [string](#string) |  | API访问授权Token |





 


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
| GetResourceList | [GetResourceListRequest](#auth.GetResourceListRequest) | [GetResourceListResponse](#auth.GetResourceListResponse) | 对外提供资源列表的接口 |
| GetActionList | [GetActionListRequest](#auth.GetActionListRequest) | [GetActionListResponse](#auth.GetActionListResponse) | 对外提供动作列表的接口 |

 



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






<a name="cmdb.HostCloudInfo"></a>

### HostCloudInfo
主机域信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 域ID |
| bk_obj_id | [string](#string) |  | 对象类型 |
| bk_obj_icon | [string](#string) |  | 对象图标 |
| bk_inst_id | [int32](#int32) |  | 实体ID |
| bk_obj_name | [string](#string) |  | 对象名称 |
| bk_inst_name | [string](#string) |  | 实体名称 |






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
| bk_host_outerip | [string](#string) |  | 外网IP |
| operator | [string](#string) |  | 主要维护人 |
| bk_bak_operator | [string](#string) |  | 备份维护人 |
| bk_asset_id | [string](#string) |  | 固资编号 |
| bk_sn | [string](#string) |  | 设备SN |
| bk_comment | [string](#string) |  | 备注 |
| bk_service_term | [int32](#int32) |  | 质保年限 |
| bk_sla | [string](#string) |  | SLA级别 |
| bk_state_name | [string](#string) |  | 所在国家 |
| bk_province_name | [string](#string) |  | 所在省份 |
| bk_isp_name | [string](#string) |  | 所属运营商 |
| bk_host_name | [string](#string) |  | 主机名称 |
| bk_os_type | [string](#string) |  | 操作系统类型 |
| bk_os_name | [string](#string) |  | 操作系统名称 |
| bk_os_version | [string](#string) |  | 操作系统版本 |
| bk_os_bit | [string](#string) |  | 操作系统位数 |
| bk_cpu | [int32](#int32) |  | CPU逻辑核心数 |
| bk_cpu_mhz | [int32](#int32) |  | CPU频率 |
| bk_cpu_module | [string](#string) |  | CPU型号 |
| bk_mem | [int32](#int32) |  | 内存容量 |
| bk_disk | [int32](#int32) |  | 磁盘容量 |
| bk_mac | [string](#string) |  | 内网MAC地址 |
| bk_outer_mac | [string](#string) |  | 外网MAC |
| create_time | [string](#string) |  | 录入时间 |
| import_from | [string](#string) |  | 录入方式 |
| cn2ip | [string](#string) |  | CN2IP |
| bond_type | [string](#string) |  | Bond类型 |
| wan_gate | [string](#string) |  | wan_gate |
| cn2_gate | [string](#string) |  | cn2_gate |
| nic_speed | [int32](#int32) |  | 网卡带宽 |
| ssh_port | [int32](#int32) |  | ssh端口 |
| cn2ip_mask | [int32](#int32) |  | cn2ip掩码数 |
| hostip_mask | [int32](#int32) |  | 设备ip掩码数 |
| function_code | [string](#string) |  | 设备功能代码 |
| state | [string](#string) |  | 主机状态 |
| service_bandwidth | [int32](#int32) |  | 服务能力 |
| ipv6 | [string](#string) |  | ipv6地址 |
| module_name | [string](#string) | repeated | 设备角色(模块名) |
| bk_host_id | [int32](#int32) |  | 主机ID |
| bk_cloud_id | [HostCloudInfo](#cmdb.HostCloudInfo) | repeated | 主机域信息 |






<a name="cmdb.ImportAssetRequest"></a>

### ImportAssetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| import_type | [string](#string) |  | 导入资产实体类型 |
| url | [string](#string) |  | xlsx文件下载路径 |
| md5 | [string](#string) |  | xlsx文件md5 |
| filename | [string](#string) |  | 用户上传xlsx文件的文件名 |






<a name="cmdb.ImportAssetResponse"></a>

### ImportAssetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |






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






<a name="cmdb.ImportDetailRequest"></a>

### ImportDetailRequest
查看导入的实体信息列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| import_id | [int32](#int32) |  | 导入记录id |
| bk_obj_id | [string](#string) |  | 指定查看实体类型 |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |






<a name="cmdb.ImportDetailResponse"></a>

### ImportDetailResponse
查看导入的实体信息列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| inst_info | [ImportDetailResponse.InstInfoEntry](#cmdb.ImportDetailResponse.InstInfoEntry) | repeated | 实体信息列表, map&lt;bk_obj_id, 可以转为map的字符串实体信息&gt; |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="cmdb.ImportDetailResponse.InstInfoEntry"></a>

### ImportDetailResponse.InstInfoEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






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
| review_note | [string](#string) |  | 审批描述 |
| filename | [string](#string) |  | 用户上传的文件名 |
| logfile_content | [string](#string) |  | 日志文件内容 |
| start_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 起始时间 |
| end_time | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 结束时间 |






<a name="cmdb.ImportHistoryRequest"></a>

### ImportHistoryRequest
查询导入历史列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| import_type | [ImportType](#cmdb.ImportType) |  | 指定查看类型 |
| import_status | [ImportStatus](#cmdb.ImportStatus) |  | 查看指定状态记录，0-查看全部 |
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






<a name="cmdb.ImportReviewRequest"></a>

### ImportReviewRequest
导入交维表审批结果请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| import_id | [int32](#int32) |  | 导入记录id |
| import_status | [ImportStatus](#cmdb.ImportStatus) |  | 指定审批结果，只处理通过，不通过和撤销 |
| review_note | [string](#string) |  | 操作描述 |






<a name="cmdb.ImportReviewResponse"></a>

### ImportReviewResponse
执行审批结果返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






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


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| with_idlepool | [bool](#bool) |  | int32 level = 1; |






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
| lake_object | [LakeObject](#cmdb.LakeObject) |  |  |
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
| node_code | [string](#string) |  | 节点编码 |






<a name="cmdb.LevelHost"></a>

### LevelHost
level下的主机ip列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| level | [LakeNodeLevel](#cmdb.LakeNodeLevel) |  | LakeNodeLevel 节点层级 |
| values | [string](#string) | repeated | 对应的ip列表 |






<a name="cmdb.ModuleHost"></a>

### ModuleHost
模块下主机ip列表，{业务模块名，[ip]}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| module_name | [string](#string) |  | 业务模块名，biz##set##module |
| value | [string](#string) | repeated | ip列表 |






<a name="cmdb.ModuleObject"></a>

### ModuleObject
模块


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_module_name | [string](#string) |  | 模块名 |
| bk_module_id | [int32](#int32) |  | 模块ID |
| TopModuleName | [string](#string) |  | 业务拓扑模块名 |
| bk_parent_id | [int32](#int32) |  | 父节点id |






<a name="cmdb.RoomObject"></a>

### RoomObject
机房对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) |  | 机房ID |
| bk_inst_name | [string](#string) |  | 机房名称 |
| bk_obj_id | [string](#string) |  | 对象ID |
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
| conditions | [string](#string) |  | 查询条件 |






<a name="cmdb.SearchHostResponse"></a>

### SearchHostResponse
查找主机请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |
| host_info_object | [HostInfoObject](#cmdb.HostInfoObject) | repeated | 主机信息 |






<a name="cmdb.SearchInstRequest"></a>

### SearchInstRequest
实体 host，lake，vip的查询请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| search_type | [InstType](#cmdb.InstType) |  | 用于区分查询类型，此处查询主要包括 host，lake，vip |
| fields | [string](#string) |  | 用于fields查询条件 |
| condition | [string](#string) |  | 用于condition查询条件 |
| page | [string](#string) |  | 用于page查询条件 |






<a name="cmdb.SearchInstResponse"></a>

### SearchInstResponse
实体 host，lake，vip的查询返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| search_inst | [string](#string) |  | 查询实体信息 |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






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






<a name="cmdb.SearchLakeHostRequest"></a>

### SearchLakeHostRequest
查询lake下的主机列表的请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lake_name | [string](#string) | repeated | lake节点的名字列表 |
| node_id | [int32](#int32) | repeated | lake节点的nodeID列表 |
| node_code | [string](#string) | repeated | lake节点的节点编码列表 |
| area_level | [AreaLevel](#cmdb.AreaLevel) |  | 要查询的区域的地区层级 |






<a name="cmdb.SearchLakeHostResponse"></a>

### SearchLakeHostResponse
查询lake下的主机列表的返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lake_host | [LakeHost](#cmdb.LakeHost) | repeated |  |
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
| node_code | [string](#string) | repeated | 节点编码(一个或多个，可为空) |






<a name="cmdb.SearchLakeResponse"></a>

### SearchLakeResponse
查找Lake请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 状态码 |
| lake | [LakeObject](#cmdb.LakeObject) | repeated | 主机信息 |






<a name="cmdb.SearchLevelHostRequest"></a>

### SearchLevelHostRequest
按level查询lake节点下主机列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| level | [LakeNodeLevel](#cmdb.LakeNodeLevel) |  | 指定level：0-全部 1-全国中心 2-区域中心 3-省边缘 4-地市边缘 5-区县边缘 |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| ipv6 | [bool](#bool) |  | 默认返回主机的ipv4地址，可以选择返回ipv6地址 |






<a name="cmdb.SearchLevelHostResponse"></a>

### SearchLevelHostResponse
按level查询lake节点下主机列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [SearchLevelHostResponse.ResultEntry](#cmdb.SearchLevelHostResponse.ResultEntry) | repeated | LakeNodeLevel 下的主机ip列表 |
| paging | [common.Paging](#common.Paging) |  | 分页信息 |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="cmdb.SearchLevelHostResponse.ResultEntry"></a>

### SearchLevelHostResponse.ResultEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [LevelHost](#cmdb.LevelHost) |  |  |






<a name="cmdb.SearchModuleHostRequest"></a>

### SearchModuleHostRequest
查询指定模块名下的全部主机请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| module_name | [string](#string) |  |  |






<a name="cmdb.SearchModuleHostResponse"></a>

### SearchModuleHostResponse
查询指定模块名下的全部主机请求返回，[{业务模块名，[ip]}]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  |  |
| data | [ModuleHost](#cmdb.ModuleHost) | repeated |  |






<a name="cmdb.SearchModuleListRequest"></a>

### SearchModuleListRequest
查询全部模块请求






<a name="cmdb.SearchModuleListResponse"></a>

### SearchModuleListResponse
查询全部模块返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  |  |
| data | [string](#string) | repeated |  |






<a name="cmdb.SearchModuleRequest"></a>

### SearchModuleRequest
查询模块请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cmdb_search_request | [ChooseHostRequest](#cmdb.ChooseHostRequest) |  |  |






<a name="cmdb.SearchModuleResponse"></a>

### SearchModuleResponse
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
| bk_inst_id | [int32](#int32) |  | 实体主键id |






<a name="cmdb.ServerListRequest"></a>

### ServerListRequest
裸金属列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_inst_id | [int32](#int32) | repeated | roomId或areaCode |
| bk_obj_id | [string](#string) |  | 对象类型 |
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
| bk_host_innerip | [string](#string) |  | 主机内网IP |
| room_id | [int32](#int32) |  | 机房id |






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
| up_bandwidth | [int32](#int32) |  | 上行带宽 |
| min_bandwidth | [int32](#int32) |  | 保底带宽 |
| maintainer | [string](#string) |  | 维护人 |
| mobile | [string](#string) |  | 联系方式 |






<a name="cmdb.SetObject"></a>

### SetObject
集群


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_set_name | [string](#string) |  | 集群名 |
| bk_set_id | [int32](#int32) |  | 集群ID |
| TopModuleName | [string](#string) |  | 业务拓扑模块名 |
| bk_parent_id | [int32](#int32) |  | 父节点id |






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






<a name="cmdb.UpdateHostStateRequest"></a>

### UpdateHostStateRequest
修改主机上下线状态的请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ipv4 | [string](#string) | repeated | ipv4,可传多个 |
| ipv6 | [string](#string) | repeated | ipv6,可传多个 |
| state | [OnlineState](#cmdb.OnlineState) |  | 主机上下线状态 |






<a name="cmdb.UpdateHostStateResponse"></a>

### UpdateHostStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="cmdb.UpdateInstRequest"></a>

### UpdateInstRequest
实体 host，lake，vip的更新请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| update_type | [InstType](#cmdb.InstType) |  | 用于区分更新类型，此处更新主要包括 host，lake，vip |
| bk_inst_id | [int32](#int32) |  | 实例id |
| update_inst | [string](#string) |  | 更新实体信息 |






<a name="cmdb.UpdateInstResponse"></a>

### UpdateInstResponse
实体 host，lake，vip的更新返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="cmdb.UpdateLakeStateRequest"></a>

### UpdateLakeStateRequest
修改lake节点上下线状态的请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node_code | [string](#string) | repeated | lake节点的节点编码列表 |
| construct_state | [OnlineState](#cmdb.OnlineState) |  | lake节点的节点建设状态 |






<a name="cmdb.UpdateLakeStateResponse"></a>

### UpdateLakeStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






<a name="cmdb.UpdateVipStateRequest"></a>

### UpdateVipStateRequest
修改VIP可用状态的请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip | [string](#string) |  | ip地址 |
| state | [OnlineState](#cmdb.OnlineState) |  | 状态 |
| ipv6 | [bool](#bool) |  | 默认接收ipv4地址，可以选择ipv6地址 |






<a name="cmdb.UpdateVipStateResponse"></a>

### UpdateVipStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |






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
| ipv4_state | [string](#string) |  | ipv4的状态 |
| ipv6_state | [string](#string) |  | ipv6的状态 |






<a name="cmdb.ZoneObject"></a>

### ZoneObject
区域


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_set_name | [string](#string) |  | 区域名 |
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
导入状态: 0-正在导入 1-导入完成 2-导入失败 3-待审核 4-审核通过 5-审核不通过 6-撤销 7-回滚失败

| Name | Number | Description |
| ---- | ------ | ----------- |
| IMPORTING | 0 | 0-正在导入 |
| COMPLETED | 1 | 1-导入完成 |
| FAILED | 2 | 2-导入失败 |
| REVIEWING | 3 | 3-待审核 |
| ACCEPT | 4 | 4-审核通过 |
| REJECT | 5 | 5-审核不通过 |
| CANCEL | 6 | 6-撤销 |
| ROLL_FAIL | 7 | 7-回滚失败 |



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
| VIP | 6 | 6-VIP |



<a name="cmdb.InstType"></a>

### InstType
实体类型:0-undefined 1-HOST 2-LAKE 3-vip  4-ROOM

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED_INST | 0 | 0-undefined |
| HOST_INST | 1 | 1-主机 |
| LAKE_INST | 2 | 2-LAKE |
| VIP_INST | 3 | 3-VIP |
| ROOM_INST | 4 | 4-ROOM |



<a name="cmdb.LakeNodeLevel"></a>

### LakeNodeLevel
节点层次

| Name | Number | Description |
| ---- | ------ | ----------- |
| ALL | 0 |  |
| NATION_CENTER | 1 | 0-全国中心 |
| REGION_CENTER | 2 | 1-区域中心 |
| PROVINCE_EDGE | 3 | 2-省边缘 |
| CITY_EDGE | 4 | 3-地市边缘 |
| COUNTY_EDGE | 5 | 4-区县边缘 |



<a name="cmdb.OnlineState"></a>

### OnlineState
上下线状态:0-未定义 1-下线 2-上线

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINEDSTATE | 0 | 0-未定义 |
| OFFLINE | 1 | 1-下线 |
| ONLINE | 2 | 2-上线 |



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
| SearchModule | [SearchModuleRequest](#cmdb.SearchModuleRequest) | [SearchModuleResponse](#cmdb.SearchModuleResponse) | 查询模块 |
| ImportHistory | [ImportHistoryRequest](#cmdb.ImportHistoryRequest) | [ImportHistoryResponse](#cmdb.ImportHistoryResponse) | 查询导入历史记录列表 |
| SearchInst | [SearchInstRequest](#cmdb.SearchInstRequest) | [SearchInstResponse](#cmdb.SearchInstResponse) | 实体的查询,包括 host，lake，vip，room列表 |
| UpdateInst | [UpdateInstRequest](#cmdb.UpdateInstRequest) | [UpdateInstResponse](#cmdb.UpdateInstResponse) | 实体的更新,包括 host，lake，vip列表 |
| SearchLake | [SearchLakeRequest](#cmdb.SearchLakeRequest) | [SearchLakeResponse](#cmdb.SearchLakeResponse) | 查询Lake节点 |
| SearchLakeHost | [SearchLakeHostRequest](#cmdb.SearchLakeHostRequest) | [SearchLakeHostResponse](#cmdb.SearchLakeHostResponse) | 查询Lake节点下的主机列表 |
| SearchLakeArea | [SearchLakeAreaRequest](#cmdb.SearchLakeAreaRequest) | [SearchLakeAreaResponse](#cmdb.SearchLakeAreaResponse) | 通过IP获取节点的地区和位置信息 |
| ImportAsset | [ImportAssetRequest](#cmdb.ImportAssetRequest) | [ImportAssetResponse](#cmdb.ImportAssetResponse) | 导入实体资产 |
| ImportReview | [ImportReviewRequest](#cmdb.ImportReviewRequest) | [ImportReviewResponse](#cmdb.ImportReviewResponse) | 导入交维表审批结果 |
| ImportDetail | [ImportDetailRequest](#cmdb.ImportDetailRequest) | [ImportDetailResponse](#cmdb.ImportDetailResponse) | 查看导入的实体信息列表 |
| SearchLevelHost | [SearchLevelHostRequest](#cmdb.SearchLevelHostRequest) | [SearchLevelHostResponse](#cmdb.SearchLevelHostResponse) | 按level查询lake节点下主机列表 |
| UpdateLakeState | [UpdateLakeStateRequest](#cmdb.UpdateLakeStateRequest) | [UpdateLakeStateResponse](#cmdb.UpdateLakeStateResponse) | 修改lake节点上下线状态 |
| UpdateHostState | [UpdateHostStateRequest](#cmdb.UpdateHostStateRequest) | [UpdateHostStateResponse](#cmdb.UpdateHostStateResponse) | 修改主机上下线状态 |
| UpdateVipState | [UpdateVipStateRequest](#cmdb.UpdateVipStateRequest) | [UpdateVipStateResponse](#cmdb.UpdateVipStateResponse) | VIP的状态更新 |
| SearchModuleList | [SearchModuleListRequest](#cmdb.SearchModuleListRequest) | [SearchModuleListResponse](#cmdb.SearchModuleListResponse) | 查询全部模块 |
| SearchModuleHost | [SearchModuleHostRequest](#cmdb.SearchModuleHostRequest) | [SearchModuleHostResponse](#cmdb.SearchModuleHostResponse) | 查询全部模块 |

 



<a name="common/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## common/api.proto



<a name="common.Condition"></a>

### Condition
指定CMDB类型条件


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bk_obj_id | [string](#string) |  | 实体类型 |
| fields | [string](#string) | repeated | 指定查找返回数据，如bk_inst_name,id，默认全部 |
| condition | [MongoCondition](#common.MongoCondition) | repeated | 实体内具体的筛选条件 |






<a name="common.MongoCondition"></a>

### MongoCondition
MongoDB型条件，如{&#34;name&#34;, &#34;$regex&#34;, &#34;小明&#34;}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [string](#string) |  | 指定筛选字段 |
| operator | [string](#string) |  | 筛选操作符，如$eq，$in，$ne等，默认$regex |
| value | [google.protobuf.Any](#google.protobuf.Any) |  | 筛选值 |






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



<a name="file.DownloadTemplateRequest"></a>

### DownloadTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_type | [TemplateType](#file.TemplateType) |  | 导入类型 |
| asset_type | [string](#string) |  | 非关系链的导入模板的类型 |






<a name="file.DownloadTemplateResponse"></a>

### DownloadTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |
| content | [bytes](#bytes) |  | 返回的excel文件 |






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
| bucket | [string](#string) |  | 指定上传bucket_name（可选） |
| key | [string](#string) |  | 指定上传存储路径（可选） |






<a name="file.UploadResponse"></a>

### UploadResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | 通用文件对象存储的url |
| md5 | [string](#string) |  | 解析出来的通用文件md5 |
| filesize | [string](#string) |  | 解析出来的通用文件大小 |
| filename | [string](#string) |  | 解析出来的通用文件name |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |





 


<a name="file.TemplateType"></a>

### TemplateType
下载模板类型

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 | 0-undefined,指除了关系链模板之外的导入模板，也可不传 |
| HOSTCHAIN | 1 | 1-导入主机业务拓扑 |
| LAKECHAIN | 2 | 2-导入lake节点关系链 |
| CROSSTABLE | 3 | 3-导入交维表 |


 

 


<a name="file.File"></a>

### File
文件服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UploadPlaybook | [UploadPlaybookRequest](#file.UploadPlaybookRequest) | [UploadPlaybookResponse](#file.UploadPlaybookResponse) | 上传playbook压缩包并且解析入口yml文件 |
| Upload | [UploadRequest](#file.UploadRequest) | [UploadResponse](#file.UploadResponse) | 上传csv等通用文件 |
| DownloadTemplate | [DownloadTemplateRequest](#file.DownloadTemplateRequest) | [DownloadTemplateResponse](#file.DownloadTemplateResponse) | 获取导入模板 |

 



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

 



<a name="image/common.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## image/common.proto



<a name="image.Paging"></a>

### Paging
分页信息(每个服务通用的控制信息，当客户端发起请求时，信息填写如下)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_page | [int32](#int32) |  | 总页数 |
| page | [int32](#int32) |  | 当前页数 |
| per_page | [int32](#int32) |  | 每页显示的记录条数 |
| total_record | [int32](#int32) |  | 总记录数 |






<a name="image.ResponseStatus"></a>

### ResponseStatus
请求返回状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [StatusCode](#image.StatusCode) |  | 状态码 |
| message | [string](#string) |  | 信息 |






<a name="image.TimeScope"></a>

### TimeScope
时间范围信息(用于指定时间范围。如按受理开始时间、受理结束时间查询订单列表。按某个具体时间查询，不使用此对象。)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| begin_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 起始时间 |
| end_date | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 结束时间 |





 


<a name="image.StatusCode"></a>

### StatusCode
请求返回状态码

| Name | Number | Description |
| ---- | ------ | ----------- |
| SUCCESS | 0 | 成功 |
| INVALID_ARGUMENT | 400 | 参数错误 |
| ACCESS_DENIED | 403 | 访问拒绝 |
| INTERNAL_ERROR | 500 | 内部错误 |


 

 

 



<a name="image/image.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## image/image.proto



<a name="image.ClusterInstance"></a>

### ClusterInstance
节点信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip_address | [string](#string) |  | 应用ip |
| cluster_name | [string](#string) |  | 应用名，集群名 |
| app_status | [AppStatus](#image.AppStatus) |  | 运行状态 |
| message | [string](#string) |  | 信息 |






<a name="image.HealthRequest"></a>

### HealthRequest
--------------- 健康监测---------------//

不需要请求体






<a name="image.HealthResponse"></a>

### HealthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [ResponseStatus](#image.ResponseStatus) |  | 响应码 |
| cluster_instances | [ClusterInstance](#image.ClusterInstance) | repeated | 响应对象 |






<a name="image.ImageAttr"></a>

### ImageAttr



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attr_name | [string](#string) |  | 属性、标签名称：例如docker |
| attr_add_type | [string](#string) |  | 属性、标签的附加信息类型，例如可以说明是version表示版本号 |
| attr_add_value | [string](#string) |  | 属性、标签附加的值,例docker的版本CE或者EE |






<a name="image.ImageObject"></a>

### ImageObject



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| image_id | [int64](#int64) |  | 镜像标识 |
| image_name | [string](#string) |  | 镜像名称 |
| image_type | [ImageType](#image.ImageType) |  | 系统类型，使用枚举，比如常用的CentOS\Ubuntu\openSUSE等 |
| version | [string](#string) |  | 系统大版本号，比如CentOS的7.4,7.5,7.6等 |
| sub_version | [string](#string) |  | 系统的小版本号，例如CentOS7.4下的1708 |
| format | [ImageFormat](#image.ImageFormat) |  | 镜像的格式，使用枚举，例如iso，qcow2等 |
| check_sum | [string](#string) |  | 镜像md5完整性校验码 |
| uri | [string](#string) |  | 镜像存储在ceph或者其他后端存储的路径或者访问地址 |
| remark | [string](#string) |  | 镜像附加信息 |
| image_attrs | [ImageAttr](#image.ImageAttr) | repeated | 镜像附加属性 |






<a name="image.QueryRequest"></a>

### QueryRequest
--------------- 查询镜像列表---------------//


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| image_name | [string](#string) |  | 镜像名称 |
| image_type | [ImageType](#image.ImageType) |  | 系统类型，使用枚举，比如常用的CentOS\Ubuntu\openSUSE等 |
| version | [string](#string) |  | 系统大版本号，比如CentOS的7.4,7.5,7.6等 |
| sub_version | [string](#string) |  | 系统的小版本号，例如CentOS7.4下的1708 |
| format | [ImageFormat](#image.ImageFormat) |  | 镜像的格式，使用枚举，例如iso，qcow2等 |
| image_attrs | [ImageAttr](#image.ImageAttr) | repeated | 镜像附加属性 |
| paging | [Paging](#image.Paging) |  | 若没有入参，则默认查询阈值为第一页10条记录 |






<a name="image.QueryResponse"></a>

### QueryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [ResponseStatus](#image.ResponseStatus) |  | 响应码 |
| images | [ImageObject](#image.ImageObject) | repeated | 响应对象 |
| paging | [Paging](#image.Paging) |  |  |





 


<a name="image.AppStatus"></a>

### AppStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| APP_STATUS_UNDEFINED | 0 | 未知 |
| APP_STATUS_STOPPED | 1 | 活动中 |
| APP_STATUS_ACTIVE | 2 | 未知 |
| APP_STATUS_UNKNOWN | 3 | 已停止 |



<a name="image.ImageFormat"></a>

### ImageFormat
镜像格式

| Name | Number | Description |
| ---- | ------ | ----------- |
| IMAGE_FORMAT_UNDEFINED | 0 |  |
| IMAGE_FORMAT_ISO | 1 |  |
| IMAGE_FORMAT_QCOW2 | 2 |  |



<a name="image.ImageType"></a>

### ImageType
系统类型

| Name | Number | Description |
| ---- | ------ | ----------- |
| IMAGE_TYPE_UNDEFINED | 0 |  |
| IMAGE_TYPE_CENTOS | 1 |  |
| IMAGE_TYPE_UBUNTU | 2 |  |
| IMAGE_TYPE_OPENSUSE | 3 |  |


 

 


<a name="image.Image"></a>

### Image


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Query | [QueryRequest](#image.QueryRequest) | [QueryResponse](#image.QueryResponse) | 查询镜像列表 |
| Health | [HealthRequest](#image.HealthRequest) | [HealthResponse](#image.HealthResponse) | 健康监测 |

 



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






<a name="ironicCom.TcpContRsp"></a>

### TcpContRsp
控制对象(每个服务通用的控制信息，当服务端反馈应答时，信息填写如下)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transactionID | [string](#string) |  | 交易流水号 |
| rspTime | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | 应答时间 |
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
| UNDEFINED_0 | 0 | 无定义 |
| IRONIC_APP | 1000010000 | 裸金属应用 |
| CMDB | 2000020000 | CMDB平台 |



<a name="ironicCom.ResultCode"></a>

### ResultCode
响应码

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED_3 | 0 | 没有定义 |
| FAIL | 1 | 失败 |
| SUCCESS | 2 | 成功 |



<a name="ironicCom.SvcCode"></a>

### SvcCode
服务编码

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED_1 | 0 | 无定义 |
| CREATE_NODES | 1010010001 | 创建裸金属节点 |
| QRY_NODE_INFO | 1010010002 | 查询裸金属状态 |
| INSTALL_NODE_SYS | 1010010003 | 安装裸金属系统 |
| OPER_NODE_POWER | 1010010004 | 操作裸金属实例电源 |
| RE_INSTALL_NODE_SYS | 1010010005 | 重新安装裸金属系统 |
| QRY_HEART_BEAT | 1010010006 | 健康监测 |


 

 

 



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






<a name="ironicServer.CreateNodesRootRsp"></a>

### CreateNodesRootRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootRsp | [CreateNodesRootRsp.ContractRootRsp](#ironicServer.CreateNodesRootRsp.ContractRootRsp) |  |  |






<a name="ironicServer.CreateNodesRootRsp.ContractRootRsp"></a>

### CreateNodesRootRsp.ContractRootRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContRsp](#ironicCom.TcpContRsp) |  |  |
| svcCont | [CreateNodesRootRsp.SvcContRsp](#ironicServer.CreateNodesRootRsp.SvcContRsp) |  |  |






<a name="ironicServer.CreateNodesRootRsp.CreateNodeInfoRsp"></a>

### CreateNodesRootRsp.CreateNodeInfoRsp
创建节点请求对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeInfoRsps | [CreateNodesRootRsp.NodeInfoRsp](#ironicServer.CreateNodesRootRsp.NodeInfoRsp) | repeated | 节点 |
| pageInfo | [ironicCom.PageInfo](#ironicCom.PageInfo) |  | 分页信息 |






<a name="ironicServer.CreateNodesRootRsp.NodeInfoRsp"></a>

### CreateNodesRootRsp.NodeInfoRsp
节点信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |
| initId | [int64](#int64) |  | 初始化id |
| createResultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 创建节点响应码 |
| createResultMsg | [string](#string) |  | 创建节点响应消息 |






<a name="ironicServer.CreateNodesRootRsp.SvcContRsp"></a>

### CreateNodesRootRsp.SvcContRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 响应码 |
| resultMsg | [string](#string) |  | 响应消息描述 |
| resultObject | [CreateNodesRootRsp.CreateNodeInfoRsp](#ironicServer.CreateNodesRootRsp.CreateNodeInfoRsp) |  | 响应对象 |






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






<a name="ironicServer.InstallNodeSysRootRsp"></a>

### InstallNodeSysRootRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootRsp | [InstallNodeSysRootRsp.ContractRootRsp](#ironicServer.InstallNodeSysRootRsp.ContractRootRsp) |  |  |






<a name="ironicServer.InstallNodeSysRootRsp.ContractRootRsp"></a>

### InstallNodeSysRootRsp.ContractRootRsp
响应信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContRsp](#ironicCom.TcpContRsp) |  |  |
| svcCont | [InstallNodeSysRootRsp.SvcContRsp](#ironicServer.InstallNodeSysRootRsp.SvcContRsp) |  |  |






<a name="ironicServer.InstallNodeSysRootRsp.InstallNodeInfoRsp"></a>

### InstallNodeSysRootRsp.InstallNodeInfoRsp
创建节点请求对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeInfoRsps | [InstallNodeSysRootRsp.NodeInfoRsp](#ironicServer.InstallNodeSysRootRsp.NodeInfoRsp) | repeated | 节点 |
| pageInfo | [ironicCom.PageInfo](#ironicCom.PageInfo) |  | 分页信息 |






<a name="ironicServer.InstallNodeSysRootRsp.NodeInfoRsp"></a>

### InstallNodeSysRootRsp.NodeInfoRsp
节点信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |
| initId | [int64](#int64) |  | 初始化id |
| installResultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 创建节点响应码 |
| installResultMsg | [string](#string) |  | 创建节点响应消息 |






<a name="ironicServer.InstallNodeSysRootRsp.SvcContRsp"></a>

### InstallNodeSysRootRsp.SvcContRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 响应码 |
| resultMsg | [string](#string) |  | 响应消息描述 |
| resultObject | [InstallNodeSysRootRsp.InstallNodeInfoRsp](#ironicServer.InstallNodeSysRootRsp.InstallNodeInfoRsp) |  | 响应对象 |






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






<a name="ironicServer.OperNodePowerRootRsp"></a>

### OperNodePowerRootRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootRsp | [OperNodePowerRootRsp.ContractRootRsp](#ironicServer.OperNodePowerRootRsp.ContractRootRsp) |  |  |






<a name="ironicServer.OperNodePowerRootRsp.ContractRootRsp"></a>

### OperNodePowerRootRsp.ContractRootRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContRsp](#ironicCom.TcpContRsp) |  |  |
| svcCont | [OperNodePowerRootRsp.SvcContRsp](#ironicServer.OperNodePowerRootRsp.SvcContRsp) |  |  |






<a name="ironicServer.OperNodePowerRootRsp.NodePowerRsp"></a>

### OperNodePowerRootRsp.NodePowerRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |
| powerResultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 创建节点响应码 |
| powerResultMsg | [string](#string) |  | 创建节点响应消息 |






<a name="ironicServer.OperNodePowerRootRsp.OperNodePowerRsp"></a>

### OperNodePowerRootRsp.OperNodePowerRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodePowerRsps | [OperNodePowerRootRsp.NodePowerRsp](#ironicServer.OperNodePowerRootRsp.NodePowerRsp) | repeated | 电源操作应答 |






<a name="ironicServer.OperNodePowerRootRsp.SvcContRsp"></a>

### OperNodePowerRootRsp.SvcContRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 响应码 |
| resultMsg | [string](#string) |  | 响应消息描述 |
| resultObject | [OperNodePowerRootRsp.OperNodePowerRsp](#ironicServer.OperNodePowerRootRsp.OperNodePowerRsp) |  | 响应对象 |






<a name="ironicServer.QryHeartbeatRootReq"></a>

### QryHeartbeatRootReq
--------------- 重新安装裸金属节点结束-----------//
--------------- 健康监测---------------//


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootReq | [QryHeartbeatRootReq.ContractRootReq](#ironicServer.QryHeartbeatRootReq.ContractRootReq) |  |  |






<a name="ironicServer.QryHeartbeatRootReq.ContractRootReq"></a>

### QryHeartbeatRootReq.ContractRootReq
请求信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContReq](#ironicCom.TcpContReq) |  |  |
| svcCont | [QryHeartbeatRootReq.SvcContReq](#ironicServer.QryHeartbeatRootReq.SvcContReq) |  |  |






<a name="ironicServer.QryHeartbeatRootReq.QryHeartbeatReq"></a>

### QryHeartbeatRootReq.QryHeartbeatReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| req | [string](#string) |  | 安装裸金属系统信息 |






<a name="ironicServer.QryHeartbeatRootReq.SvcContReq"></a>

### QryHeartbeatRootReq.SvcContReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requestObject | [QryHeartbeatRootReq.QryHeartbeatReq](#ironicServer.QryHeartbeatRootReq.QryHeartbeatReq) |  |  |






<a name="ironicServer.QryHeartbeatRootRsp"></a>

### QryHeartbeatRootRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootRsp | [QryHeartbeatRootRsp.ContractRootRsp](#ironicServer.QryHeartbeatRootRsp.ContractRootRsp) |  |  |






<a name="ironicServer.QryHeartbeatRootRsp.ClusterInstance"></a>

### QryHeartbeatRootRsp.ClusterInstance
节点信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ipAddress | [string](#string) |  | 应用ip |
| clusterName | [string](#string) |  | 应用名，集群名 |
| appStatus | [QryHeartbeatRootRsp.AppStatus](#ironicServer.QryHeartbeatRootRsp.AppStatus) |  | 运行状态 |
| message | [string](#string) |  | 信息 |






<a name="ironicServer.QryHeartbeatRootRsp.ContractRootRsp"></a>

### QryHeartbeatRootRsp.ContractRootRsp
响应信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContRsp](#ironicCom.TcpContRsp) |  |  |
| svcCont | [QryHeartbeatRootRsp.SvcContRsp](#ironicServer.QryHeartbeatRootRsp.SvcContRsp) |  |  |






<a name="ironicServer.QryHeartbeatRootRsp.QryHeartbeatRsp"></a>

### QryHeartbeatRootRsp.QryHeartbeatRsp
健康监测应答对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterInstances | [QryHeartbeatRootRsp.ClusterInstance](#ironicServer.QryHeartbeatRootRsp.ClusterInstance) | repeated | 应用信息 |






<a name="ironicServer.QryHeartbeatRootRsp.SvcContRsp"></a>

### QryHeartbeatRootRsp.SvcContRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 响应码 |
| resultMsg | [string](#string) |  | 响应消息描述 |
| resultObject | [QryHeartbeatRootRsp.QryHeartbeatRsp](#ironicServer.QryHeartbeatRootRsp.QryHeartbeatRsp) |  | 响应对象 |






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






<a name="ironicServer.QryNodeInfoRootRsp"></a>

### QryNodeInfoRootRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootRsp | [QryNodeInfoRootRsp.ContractRootRsp](#ironicServer.QryNodeInfoRootRsp.ContractRootRsp) |  |  |






<a name="ironicServer.QryNodeInfoRootRsp.ContractRootRsp"></a>

### QryNodeInfoRootRsp.ContractRootRsp
响应信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContRsp](#ironicCom.TcpContRsp) |  |  |
| svcCont | [QryNodeInfoRootRsp.SvcContRsp](#ironicServer.QryNodeInfoRootRsp.SvcContRsp) |  |  |






<a name="ironicServer.QryNodeInfoRootRsp.NodeInfo"></a>

### QryNodeInfoRootRsp.NodeInfo
节点信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |
| status | [QryNodeInfoRootRsp.Status](#ironicServer.QryNodeInfoRootRsp.Status) |  | 当前状态 |
| ip | [string](#string) |  | ipmi的IP地址 |
| userName | [string](#string) |  | ipmi用户名 |
| password | [string](#string) |  | Ipmi密码 |
| macAddr | [string](#string) |  | 裸金属主机mac地址 |
| imagerAddr | [string](#string) |  | 镜像地址 |
| serverAddr | [string](#string) |  | 所属区域裸金属组件服务地址 |
| qryResultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 查询节点响应码 |
| qryResultMsg | [string](#string) |  | 查询节点响应消息 |
| powerStatus | [QryNodeInfoRootRsp.PowerStatus](#ironicServer.QryNodeInfoRootRsp.PowerStatus) |  | 电源状态 |






<a name="ironicServer.QryNodeInfoRootRsp.QryNodeInfoRsp"></a>

### QryNodeInfoRootRsp.QryNodeInfoRsp
查询节点应答对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeInfos | [QryNodeInfoRootRsp.NodeInfo](#ironicServer.QryNodeInfoRootRsp.NodeInfo) | repeated | 节点 |






<a name="ironicServer.QryNodeInfoRootRsp.SvcContRsp"></a>

### QryNodeInfoRootRsp.SvcContRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 响应码 |
| resultMsg | [string](#string) |  | 响应消息描述 |
| resultObject | [QryNodeInfoRootRsp.QryNodeInfoRsp](#ironicServer.QryNodeInfoRootRsp.QryNodeInfoRsp) |  | 响应对象 |






<a name="ironicServer.ReInstallNodeSysRootReq"></a>

### ReInstallNodeSysRootReq
--------------- 重新安装裸金属节点开始-----------//


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootReq | [ReInstallNodeSysRootReq.ContractRootReq](#ironicServer.ReInstallNodeSysRootReq.ContractRootReq) |  |  |






<a name="ironicServer.ReInstallNodeSysRootReq.ContractRootReq"></a>

### ReInstallNodeSysRootReq.ContractRootReq
请求信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContReq](#ironicCom.TcpContReq) |  |  |
| svcCont | [ReInstallNodeSysRootReq.SvcContReq](#ironicServer.ReInstallNodeSysRootReq.SvcContReq) |  |  |






<a name="ironicServer.ReInstallNodeSysRootReq.InstallNodeSysReq"></a>

### ReInstallNodeSysRootReq.InstallNodeSysReq
安装裸金属系统请求对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeInstallInfos | [ReInstallNodeSysRootReq.NodeInstallInfo](#ironicServer.ReInstallNodeSysRootReq.NodeInstallInfo) | repeated | 安装裸金属系统信息 |






<a name="ironicServer.ReInstallNodeSysRootReq.NodeInstallInfo"></a>

### ReInstallNodeSysRootReq.NodeInstallInfo
安装裸金属系统信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeId | [string](#string) |  | 裸金属节点id |
| imageAddr | [string](#string) |  | 安装镜像地址(一般为http url) |
| imageCheckSum | [string](#string) |  | 镜像MD5校验码 |






<a name="ironicServer.ReInstallNodeSysRootReq.SvcContReq"></a>

### ReInstallNodeSysRootReq.SvcContReq
安装裸金属系统请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requestObject | [ReInstallNodeSysRootReq.InstallNodeSysReq](#ironicServer.ReInstallNodeSysRootReq.InstallNodeSysReq) |  |  |






<a name="ironicServer.ReInstallNodeSysRootRsp"></a>

### ReInstallNodeSysRootRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractRootRsp | [ReInstallNodeSysRootRsp.ContractRootRsp](#ironicServer.ReInstallNodeSysRootRsp.ContractRootRsp) |  |  |






<a name="ironicServer.ReInstallNodeSysRootRsp.ContractRootRsp"></a>

### ReInstallNodeSysRootRsp.ContractRootRsp
响应信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcpCont | [ironicCom.TcpContRsp](#ironicCom.TcpContRsp) |  |  |
| svcCont | [ReInstallNodeSysRootRsp.SvcContRsp](#ironicServer.ReInstallNodeSysRootRsp.SvcContRsp) |  |  |






<a name="ironicServer.ReInstallNodeSysRootRsp.InstallNodeInfoRsp"></a>

### ReInstallNodeSysRootRsp.InstallNodeInfoRsp
创建节点请求对象


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeInfoRsps | [ReInstallNodeSysRootRsp.ReNodeInfoRsp](#ironicServer.ReInstallNodeSysRootRsp.ReNodeInfoRsp) | repeated | 节点 |
| pageInfo | [ironicCom.PageInfo](#ironicCom.PageInfo) |  | 分页信息 |






<a name="ironicServer.ReInstallNodeSysRootRsp.ReNodeInfoRsp"></a>

### ReInstallNodeSysRootRsp.ReNodeInfoRsp
节点信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| oldNodeId | [string](#string) |  | 旧裸金属节点id |
| newNodeId | [string](#string) |  | 新裸金属节点id |
| initId | [int64](#int64) |  | 初始化id |
| installResultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 创建节点响应码 |
| installResultMsg | [string](#string) |  | 创建节点响应消息 |






<a name="ironicServer.ReInstallNodeSysRootRsp.SvcContRsp"></a>

### ReInstallNodeSysRootRsp.SvcContRsp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultCode | [ironicCom.ResultCode](#ironicCom.ResultCode) |  | 响应码 |
| resultMsg | [string](#string) |  | 响应消息描述 |
| resultObject | [ReInstallNodeSysRootRsp.InstallNodeInfoRsp](#ironicServer.ReInstallNodeSysRootRsp.InstallNodeInfoRsp) |  | 响应对象 |





 


<a name="ironicServer.QryHeartbeatRootRsp.AppStatus"></a>

### QryHeartbeatRootRsp.AppStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 | 没有定义 |
| STOPPED | 1 | 活动中 |
| ACTIVE | 2 | 未知 |
| UNKNOWN | 3 | 已停止 |



<a name="ironicServer.QryNodeInfoRootRsp.PowerStatus"></a>

### QryNodeInfoRootRsp.PowerStatus
电源状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED_1 | 0 | 没有定义 |
| POWERON | 1 | 开机 |
| POWEROFF | 2 | 关机 |
| UNKNOWN | 3 | 未知 |



<a name="ironicServer.QryNodeInfoRootRsp.Status"></a>

### QryNodeInfoRootRsp.Status
裸金属节点状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED_2 | 0 | 没有定义 |
| NULL | 1400 | 未知 |
| CREATED | 1300 | 已创建 :刚创建完节点后的状态，节点创建、验证成功，可进行装系统 |
| DEPLOYING | 1200 | 部署中 :正在执行装机任务 |
| DEPLOYFAILED | 1100 | 部署失败 : 执行装机异常 |
| ACTIVE | 1000 | 活动 ： 已经安装完系统活动中 |


 

 


<a name="ironicServer.IronicServer"></a>

### IronicServer


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| qryNodeInfo | [QryNodeInfoRootReq](#ironicServer.QryNodeInfoRootReq) | [QryNodeInfoRootRsp](#ironicServer.QryNodeInfoRootRsp) | 裸金属节点状态查询 |
| createNodes | [CreateNodesRootReq](#ironicServer.CreateNodesRootReq) | [CreateNodesRootRsp](#ironicServer.CreateNodesRootRsp) | 创建裸金属节点 |
| installNodeSys | [InstallNodeSysRootReq](#ironicServer.InstallNodeSysRootReq) | [InstallNodeSysRootRsp](#ironicServer.InstallNodeSysRootRsp) | 安装裸金属系统 |
| reInstallNodeSys | [ReInstallNodeSysRootReq](#ironicServer.ReInstallNodeSysRootReq) | [ReInstallNodeSysRootRsp](#ironicServer.ReInstallNodeSysRootRsp) | 重装裸金属系统 |
| operNodePower | [OperNodePowerRootReq](#ironicServer.OperNodePowerRootReq) | [OperNodePowerRootRsp](#ironicServer.OperNodePowerRootRsp) | 操作裸金属实例电源 |
| qryHeartbeat | [QryHeartbeatRootReq](#ironicServer.QryHeartbeatRootReq) | [QryHeartbeatRootRsp](#ironicServer.QryHeartbeatRootRsp) | 健康监测 |

 



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
| log | [bytes](#bytes) |  | 执行日志 |
| log_url | [string](#string) |  | 执行日志访问路径 |






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






<a name="subtask.GetLogRequest"></a>

### GetLogRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 作业子任务ID |






<a name="subtask.GetLogResponse"></a>

### GetLogResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |
| sub_task_log | [string](#string) |  | 子任务实例名 |
| sub_task_status | [SubTaskState](#subtask.SubTaskState) |  |  |





 


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



<a name="subtask.SubTaskState"></a>

### SubTaskState
模板状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| COMPLETE | 0 | 执行完成 |
| RUNNING | 1 | 执行中 |


 

 


<a name="subtask.SubTask"></a>

### SubTask
子任务实例

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#subtask.CreateRequest) | [CreateResponse](#subtask.CreateResponse) | 创建作业子任务 |
| Complete | [CompleteRequest](#subtask.CompleteRequest) | [CompleteResponse](#subtask.CompleteResponse) | 完成作业子任务 |
| CreateServer | [CreateServerRequest](#subtask.CreateServerRequest) | [CreateServerResponse](#subtask.CreateServerResponse) | 创建裸金属子任务 |
| CreateServerCompare | [CreateServerCompareRequest](#subtask.CreateServerCompareRequest) | [CreateServerCompareResponse](#subtask.CreateServerCompareResponse) | 交维表导入后巡检子任务 |
| GetLog | [GetLogRequest](#subtask.GetLogRequest) | [GetLogResponse](#subtask.GetLogResponse) | 获取作业任务SubTask具体日志 |

 



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






<a name="task.JobAgentLog"></a>

### JobAgentLog
agent上报kafka日志结构体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 任务实例ID |
| job_agent_log | [string](#string) |  | ansible执行日志 |






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
| extra_var | [string](#string) |  | 额外变量 |






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
| template_id | [int32](#int32) |  | 检查作业模板 |






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






<a name="task.GetSubTaskRequest"></a>

### GetSubTaskRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [int64](#int64) |  | 作业任务ID |






<a name="task.GetSubTaskResponse"></a>

### GetSubTaskResponse
获取作业任务详细执行过程请求返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [common.ResponseStatus](#common.ResponseStatus) |  | 返回的请求状态 |
| sub_task_info | [SubTaskInfo](#task.SubTaskInfo) | repeated | 子任务信息列表 |






<a name="task.InstallServerRequest"></a>

### InstallServerRequest
调用裸金属安装操作请求体


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cmdb_search_request | [cmdb.ChooseServerRequest](#cmdb.ChooseServerRequest) | repeated | 机房及对应物理机 |
| image_file_url | [string](#string) |  | 安装镜像文件的URL |
| image_file_md5 | [string](#string) |  | 安装镜像文件的MD5 |






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






<a name="task.SubTaskInfo"></a>

### SubTaskInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sub_task_id | [int64](#int64) |  | 子任务实例ID |
| sub_task_name | [string](#string) |  | 子任务实例名 |






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
| executor | [common.User](#common.User) |  | 执行人 |
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
| GetSubTask | [GetSubTaskRequest](#task.GetSubTaskRequest) | [GetSubTaskResponse](#task.GetSubTaskResponse) | 获取作业任务日志SubTask列表 |
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
| release_code | [ReleaseCode](#template.ReleaseCode) |  | 模板发布状态码，用于标记某些特定的模板，相同release code默认使用最进更新的template |






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
| release_code | [ReleaseCode](#template.ReleaseCode) |  | 模板发布状态码，用于标记某些特定的模板，相同release code默认获取最新的template(优先匹配template_id) |






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
| release_code | [ReleaseCode](#template.ReleaseCode) |  | 模板发布状态码，用于标记某些特定的模板，相同release code默认使用最进更新的template |






<a name="template.UpdateRequest"></a>

### UpdateRequest
更新模板请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_id | [int32](#int32) |  | 模板ID |
| name | [string](#string) |  | 模板名 |
| description | [string](#string) |  | 模板描述 |
| extra_var | [string](#string) |  | 额外变量JSON String 例如： {&#34;key&#34;:&#34;testKey&#34;,&#34;value&#34;:&#34;testVal&#34;,&#34;description&#34;:&#34;测试描述&#34;} |
| release_code | [ReleaseCode](#template.ReleaseCode) |  | 模板发布状态码，用于标记某些特定的模板，相同release code默认使用最进更新的template |






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





 


<a name="template.ReleaseCode"></a>

### ReleaseCode
系统模板的Release code

| Name | Number | Description |
| ---- | ------ | ----------- |
| RELEASE_CODE_UNDEFINED | 0 |  |
| RELEASE_CODE_SERVER_COMPARE | 1 | 服务器巡检模板 |
| RELEASE_CODE_IRONIC_APP | 2 | 裸金属服务 |



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

