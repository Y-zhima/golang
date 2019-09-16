// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cmdb/cmdb.proto

package cmdb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "git.fogcdn.top/axe/protos/goout/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *TopologyObject) Validate() error {
	for _, item := range this.Child {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Child", err)
			}
		}
	}
	return nil
}
func (this *InstanceTopologyRequest) Validate() error {
	return nil
}
func (this *InstanceTopologyResponse) Validate() error {
	if this.Instance != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Instance); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Instance", err)
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *HostObject) Validate() error {
	return nil
}
func (this *ModuleObject) Validate() error {
	return nil
}
func (this *SetObject) Validate() error {
	return nil
}
func (this *ZoneObject) Validate() error {
	return nil
}
func (this *BizObject) Validate() error {
	return nil
}
func (this *CommonObject) Validate() error {
	return nil
}
func (this *CreateAssociationRequest) Validate() error {
	return nil
}
func (this *HostInfoObject) Validate() error {
	if this.Host != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Host); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Host", err)
		}
	}
	for _, item := range this.Module {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Module", err)
			}
		}
	}
	for _, item := range this.Set {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Set", err)
			}
		}
	}
	for _, item := range this.Biz {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Biz", err)
			}
		}
	}
	return nil
}
func (this *SearchHostRequest) Validate() error {
	if this.Paging != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Paging); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Paging", err)
		}
	}
	if this.Host != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Host); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Host", err)
		}
	}
	if this.Module != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Module); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Module", err)
		}
	}
	if this.Set != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Set); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Set", err)
		}
	}
	if this.Biz != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Biz); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Biz", err)
		}
	}
	if this.Zone != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Zone); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Zone", err)
		}
	}
	return nil
}
func (this *SearchHostResponse) Validate() error {
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
	for _, item := range this.Info {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
			}
		}
	}
	return nil
}
func (this *ImportHostRequest) Validate() error {
	if this.Url == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Url", fmt.Errorf(`URL不能为空`))
	}
	return nil
}
func (this *ImportHostResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *ImportServerRequest) Validate() error {
	if this.Url == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Url", fmt.Errorf(`URL不能为空`))
	}
	return nil
}
func (this *ImportServerResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *ChooseHostRequest) Validate() error {
	return nil
}
func (this *Server) Validate() error {
	return nil
}
func (this *ChooseServerRequest) Validate() error {
	for _, item := range this.Server {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Server", err)
			}
		}
	}
	return nil
}
func (this *ImportSwitchRequest) Validate() error {
	if this.Url == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Url", fmt.Errorf(`URL不能为空`))
	}
	return nil
}
func (this *ImportSwitchResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *ImportLakeRequest) Validate() error {
	if this.Url == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Url", fmt.Errorf(`URL不能为空`))
	}
	return nil
}
func (this *ImportLakeResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *RoomObject) Validate() error {
	for _, item := range this.Child {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Child", err)
			}
		}
	}
	return nil
}
func (this *RoomTopologyRequest) Validate() error {
	return nil
}
func (this *RoomTopologyResponse) Validate() error {
	for _, item := range this.Rooms {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rooms", err)
			}
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *ServerObject) Validate() error {
	return nil
}
func (this *ServerListRequest) Validate() error {
	if this.Paging != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Paging); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Paging", err)
		}
	}
	return nil
}
func (this *ServerListResponse) Validate() error {
	for _, item := range this.Servers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Servers", err)
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
func (this *SearchMoudleRequest) Validate() error {
	if this.CmdbSearchRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CmdbSearchRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CmdbSearchRequest", err)
		}
	}
	return nil
}
func (this *SearchMoudleResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	for _, item := range this.Moudule {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Moudule", err)
			}
		}
	}
	return nil
}
func (this *ImportHistoryObject) Validate() error {
	if this.Operator != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Operator); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Operator", err)
		}
	}
	if this.StartTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartTime", err)
		}
	}
	if this.EndTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EndTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EndTime", err)
		}
	}
	return nil
}
func (this *ImportHistoryRequest) Validate() error {
	if this.Paging != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Paging); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Paging", err)
		}
	}
	return nil
}
func (this *ImportHistoryResponse) Validate() error {
	for _, item := range this.History {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("History", err)
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
func (this *AreaObject) Validate() error {
	return nil
}
func (this *IspObject) Validate() error {
	return nil
}
func (this *VipObject) Validate() error {
	return nil
}
func (this *ServerRoomObject) Validate() error {
	return nil
}
func (this *LakeObject) Validate() error {
	for _, item := range this.Vip {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Vip", err)
			}
		}
	}
	if this.ConstructUpdatetime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConstructUpdatetime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConstructUpdatetime", err)
		}
	}
	for _, item := range this.Room {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Room", err)
			}
		}
	}
	return nil
}
func (this *AssociationObject) Validate() error {
	return nil
}
func (this *SearchLakeRequest) Validate() error {
	return nil
}
func (this *SearchLakeResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	for _, item := range this.Lake {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Lake", err)
			}
		}
	}
	return nil
}
