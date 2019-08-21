// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: template/template.proto

package template

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "git.fogcdn.top/axe/protos/goout/common"
	_ "git.fogcdn.top/axe/protos/goout/playbook"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *TemplateObject) Validate() error {
	if this.Playbook != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Playbook); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Playbook", err)
		}
	}
	if this.Creator != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Creator); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Creator", err)
		}
	}
	return nil
}
func (this *CreateRequest) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`模板名称不能为空`))
	}
	if !(this.PlaybookId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PlaybookId", fmt.Errorf(`项目ID不能为空`))
	}
	if !(this.PlaybookFileId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PlaybookFileId", fmt.Errorf(`项目文件ID不能为空`))
	}
	if !(this.PlaybookEntrypointId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PlaybookEntrypointId", fmt.Errorf(`项目入口文件ID不能为空`))
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
	if this.Template != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Template); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Template", err)
		}
	}
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
	for _, item := range this.Templates {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Templates", err)
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
func (this *UpdateRequest) Validate() error {
	if !(this.TemplateId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TemplateId", fmt.Errorf(`模板ID不能为空`))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`模板名称不能为空`))
	}
	return nil
}
func (this *UpdateResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *DeleteRequest) Validate() error {
	return nil
}
func (this *DeleteResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
