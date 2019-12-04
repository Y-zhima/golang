// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: image/common.proto

package image

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PageInfo) Validate() error {
	return nil
}
func (this *TimeScope) Validate() error {
	if this.BeginDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BeginDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BeginDate", err)
		}
	}
	if this.EndDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EndDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EndDate", err)
		}
	}
	return nil
}
func (this *ResponseStatus) Validate() error {
	return nil
}
