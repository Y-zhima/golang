// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: subtask/subtask.proto

package subtask

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "git.fogcdn.top/axe/protos/goout/cmdb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "git.fogcdn.top/axe/protos/goout/common"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateRequest) Validate() error {
	if !(this.TaskId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TaskId", fmt.Errorf(`任务ID不能为空`))
	}
	if nil == this.CmdbSearchRequest {
		return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", fmt.Errorf("message must exist"))
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
func (this *CreateServerCompareRequest) Validate() error {
	if !(this.TaskId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TaskId", fmt.Errorf(`任务ID不能为空`))
	}
	if nil == this.CmdbSearchRequest {
		return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", fmt.Errorf("message must exist"))
	}
	if this.CmdbSearchRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CmdbSearchRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
		}
	}
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
func (this *CreateServerRequest) Validate() error {
	if !(this.TaskId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TaskId", fmt.Errorf(`任务ID不能为空`))
	}
	if nil == this.CmdbSearchRequest {
		return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", fmt.Errorf("message must exist"))
	}
	if this.CmdbSearchRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CmdbSearchRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
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
func (this *CompleteRequest) Validate() error {
	return nil
}
func (this *CompleteResponse) Validate() error {
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
