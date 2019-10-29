# Protocol Documentation
<a name="top"></a>

## Table of Contents

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

