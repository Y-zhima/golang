// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schedule/schedule.proto

package schedule

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "git.fogcdn.top/axe/protos/goout/common"
	_ "git.fogcdn.top/axe/protos/goout/template"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ScheduleObject) Validate() error {
	if this.Template != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Template); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Template", err)
		}
	}
	return nil
}
func (this *CreateRequest) Validate() error {
	if !(this.TemplateId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TemplateId", fmt.Errorf(`模板ID不能为空`))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`定时任务名不能为空`))
	}
	if this.StartTime == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("StartTime", fmt.Errorf(`开始时间不能为空`))
	}
	if this.EndTime == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EndTime", fmt.Errorf(`结束时间不能为空`))
	}
	if this.CronExpression == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CronExpression", fmt.Errorf(`定时任务表达式不能为空`))
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
	if this.Schedule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Schedule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Schedule", err)
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
	for _, item := range this.Schedules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Schedules", err)
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
	if !(this.ScheduleId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleId", fmt.Errorf(`定时任务ID不能为空`))
	}
	if !(this.TemplateId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TemplateId", fmt.Errorf(`模板ID不能为空`))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`定时任务名不能为空`))
	}
	if this.StartTime == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("StartTime", fmt.Errorf(`开始时间不能为空`))
	}
	if this.EndTime == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EndTime", fmt.Errorf(`结束时间不能为空`))
	}
	if this.CronExpression == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CronExpression", fmt.Errorf(`定时任务表达式不能为空`))
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
func (this *SwitchStatusRequest) Validate() error {
	return nil
}
func (this *SwitchStatusResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
