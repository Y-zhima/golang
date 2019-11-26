// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cmdb/cmdb.proto

package cmdb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "git.fogcdn.top/axe/protos/goout/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SearchInstRequest) Validate() error {
	return nil
}
func (this *SearchInstResponse) Validate() error {
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
func (this *UpdateInstRequest) Validate() error {
	return nil
}
func (this *UpdateInstResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
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
	for _, item := range this.Instance {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Instance", err)
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
	for _, item := range this.HostInfoObject {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HostInfoObject", err)
			}
		}
	}
	return nil
}
func (this *ImportCrossTableRequest) Validate() error {
	if this.Url == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Url", fmt.Errorf(`URL不能为空`))
	}
	return nil
}
func (this *ImportCrossTableResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
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
func (this *ChooseServerCompareRequest) Validate() error {
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
	if this.Area != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Area); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Area", err)
		}
	}
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
func (this *CabinetObject) Validate() error {
	return nil
}
func (this *SearchLakeHostRequest) Validate() error {
	return nil
}
func (this *SearchLakeHostResponse) Validate() error {
	for _, item := range this.LakeHost {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LakeHost", err)
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
func (this *LakeHost) Validate() error {
	if this.LakeObject != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LakeObject); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LakeObject", err)
		}
	}
	for _, item := range this.Host {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Host", err)
			}
		}
	}
	return nil
}
func (this *SearchLakeAreaRequest) Validate() error {
	return nil
}
func (this *SearchLakeAreaResponse) Validate() error {
	for _, item := range this.LakeArea {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LakeArea", err)
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
func (this *LakeAreaObject) Validate() error {
	for _, item := range this.Lake {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Lake", err)
			}
		}
	}
	return nil
}
func (this *ImportAssetRequest) Validate() error {
	if this.ImportType == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ImportType", fmt.Errorf(`importType不能为空`))
	}
	if this.Url == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Url", fmt.Errorf(`URL不能为空`))
	}
	if this.Md5 == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Md5", fmt.Errorf(`MD5不能为空`))
	}
	if this.Filename == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Filename", fmt.Errorf(`fileName不能为空`))
	}
	return nil
}
func (this *ImportAssetResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *ImportReviewRequest) Validate() error {
	if !(this.ImportId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ImportId", fmt.Errorf(`导入记录ID不能为空`))
	}
	return nil
}
func (this *ImportReviewResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *ImportDetailRequest) Validate() error {
	if !(this.ImportId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ImportId", fmt.Errorf(`导入记录ID不能为空`))
	}
	if this.Paging != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Paging); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Paging", err)
		}
	}
	return nil
}
func (this *ImportDetailResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
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
func (this *SearchLevelHostRequest) Validate() error {
	if this.Paging != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Paging); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Paging", err)
		}
	}
	return nil
}
func (this *SearchLevelHostResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
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
func (this *LevelHost) Validate() error {
	return nil
}
func (this *UpdateLakeStateRequest) Validate() error {
	return nil
}
func (this *UpdateLakeStateResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *UpdateHostStateRequest) Validate() error {
	return nil
}
func (this *UpdateHostStateResponse) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
