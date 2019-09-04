// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: queue/job.proto

package queue

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "git.fogcdn.top/axe/protos/goout/cmdb"
	_ "git.fogcdn.top/axe/protos/goout/schedule"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DeliverJobObject) Validate() error {
	if this.CmdbSearchRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CmdbSearchRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
		}
	}
	return nil
}
