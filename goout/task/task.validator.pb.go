// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task/task.proto

package task

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "git.fogcdn.top/axe/protos/goout/cmdb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "git.fogcdn.top/axe/protos/goout/common"
	_ "git.fogcdn.top/axe/protos/goout/schedule"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CheckServerStateRequest) Validate() error {
	for _, item := range this.CmdbSearchRequest {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
			}
		}
	}
	return nil
}
func (this *CheckServerStateResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *ServerPowerControlRequest) Validate() error {
	for _, item := range this.CmdbSearchRequest {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
			}
		}
	}
	return nil
}
func (this *ServerPowerControlResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *CreateServerRequest) Validate() error {
	for _, item := range this.CmdbSearchRequest {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
			}
		}
	}
	return nil
}
func (this *CreateServerResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *InstallServerRequest) Validate() error {
	for _, item := range this.CmdbSearchRequest {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
			}
		}
	}
	if this.CmdbHostSearchRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CmdbHostSearchRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CmdbHostSearchRequest", err)
		}
	}
	return nil
}
func (this *InstallServerResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *CreateServerCompareRequest) Validate() error {
	return nil
}
func (this *CreateServerCompareResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *TaskObject) Validate() error {
	for _, item := range this.CmdbSearchRequest {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
			}
		}
	}
	if this.Executor != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Executor); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Executor", err)
		}
	}
	return nil
}
func (this *CreateRequest) Validate() error {
	if !(this.TemplateId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TemplateId", fmt.Errorf(`模板ID不能为空`))
	}
	for _, item := range this.CmdbSearchRequest {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
			}
		}
	}
	return nil
}
func (this *CreateResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *GetRequest) Validate() error {
	return nil
}
func (this *GetResponse) Validate() error {
	if this.Task != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Task); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Task", err)
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *GetLogRequest) Validate() error {
	return nil
}
func (this *GetLogResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *GetSubTaskRequest) Validate() error {
	return nil
}
func (this *GetSubTaskResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	for _, item := range this.SubTaskInfo {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SubTaskInfo", err)
			}
		}
	}
	return nil
}
func (this *SubTaskInfo) Validate() error {
	return nil
}
func (this *FilterRequest) Validate() error {
	if this.Paging != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Paging); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Paging", err)
		}
	}
	return nil
}
func (this *FilterResponse) Validate() error {
	for _, item := range this.Tasks {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Tasks", err)
			}
		}
	}
	if this.Paging != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Paging); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Paging", err)
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *RetryRequest) Validate() error {
	return nil
}
func (this *RetryResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
