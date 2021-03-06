// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth/auth.proto

package auth

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "git.fogcdn.top/axe/protos/goout/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CheckRequest) Validate() error {
	return nil
}
func (this *CheckResponse) Validate() error {
	return nil
}
func (this *LoginRequest) Validate() error {
	return nil
}
func (this *LoginResponse) Validate() error {
	return nil
}
func (this *LogoutRequest) Validate() error {
	return nil
}
func (this *LogoutResponse) Validate() error {
	return nil
}
func (this *CallbackRequest) Validate() error {
	return nil
}
func (this *CallbackResponse) Validate() error {
	return nil
}
func (this *GetUserInfoRequest) Validate() error {
	return nil
}
func (this *UserInfo) Validate() error {
	return nil
}
func (this *GetUserInfoResponse) Validate() error {
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *ResourceActionItem) Validate() error {
	return nil
}
func (this *ResourceActionData) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *GetResourceListRequest) Validate() error {
	return nil
}
func (this *GetResourceListResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *GetActionListRequest) Validate() error {
	return nil
}
func (this *GetActionListResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
