// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task/task.proto

package task

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "git.fogcdn.top/axe/protos/goout/common"
	_ "git.fogcdn.top/axe/protos/goout/schedule"
	_ "git.fogcdn.top/axe/protos/goout/cmdb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *TaskObject) Validate() error {
	if this.CmdbSearchRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CmdbSearchRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
		}
	}
	return nil
}
func (this *CreateRequest) Validate() error {
	if !(this.TemplateId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TemplateId", fmt.Errorf(`模板ID不能为空`))
	}
	if this.CmdbSearchRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CmdbSearchRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
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
