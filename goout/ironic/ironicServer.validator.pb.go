// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ironic/ironicServer.proto

package ironic

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *QryNodeInfoRootReq) Validate() error {
	if this.ContractRootReq != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContractRootReq); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContractRootReq", err)
		}
	}
	return nil
}
func (this *QryNodeInfoRootReq_ContractRootReq) Validate() error {
	if this.TcpCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TcpCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TcpCont", err)
		}
	}
	if this.SvcCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SvcCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SvcCont", err)
		}
	}
	return nil
}
func (this *QryNodeInfoRootReq_SvcContReq) Validate() error {
	if this.RequestObject != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RequestObject); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RequestObject", err)
		}
	}
	return nil
}
func (this *QryNodeInfoRootReq_QryNodeInfoReq) Validate() error {
	return nil
}
func (this *QryNodeInfoRootRes) Validate() error {
	if this.ContractRootRes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContractRootRes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContractRootRes", err)
		}
	}
	return nil
}
func (this *QryNodeInfoRootRes_ContractRootRes) Validate() error {
	if this.TcpCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TcpCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TcpCont", err)
		}
	}
	if this.SvcCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SvcCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SvcCont", err)
		}
	}
	return nil
}
func (this *QryNodeInfoRootRes_SvcContRes) Validate() error {
	if this.ResultObject != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResultObject); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResultObject", err)
		}
	}
	return nil
}
func (this *QryNodeInfoRootRes_QryNodeInfoRsp) Validate() error {
	for _, item := range this.NodeInfos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodeInfos", err)
			}
		}
	}
	return nil
}
func (this *QryNodeInfoRootRes_NodeInfo) Validate() error {
	return nil
}
func (this *InstallNodeSysRootReq) Validate() error {
	if this.ContractRootReq != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContractRootReq); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContractRootReq", err)
		}
	}
	return nil
}
func (this *InstallNodeSysRootReq_ContractRootReq) Validate() error {
	if this.TcpCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TcpCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TcpCont", err)
		}
	}
	if this.SvcCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SvcCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SvcCont", err)
		}
	}
	return nil
}
func (this *InstallNodeSysRootReq_SvcContReq) Validate() error {
	if this.RequestObject != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RequestObject); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RequestObject", err)
		}
	}
	return nil
}
func (this *InstallNodeSysRootReq_InstallNodeSysReq) Validate() error {
	for _, item := range this.NodeInstallInfos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodeInstallInfos", err)
			}
		}
	}
	return nil
}
func (this *InstallNodeSysRootReq_NodeInstallInfo) Validate() error {
	return nil
}
func (this *InstallNodeSysRootRes) Validate() error {
	if this.ContractRootRes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContractRootRes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContractRootRes", err)
		}
	}
	return nil
}
func (this *InstallNodeSysRootRes_ContractRootRes) Validate() error {
	if this.TcpCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TcpCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TcpCont", err)
		}
	}
	if this.SvcCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SvcCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SvcCont", err)
		}
	}
	return nil
}
func (this *InstallNodeSysRootRes_SvcContRes) Validate() error {
	return nil
}
func (this *OperNodePowerRootReq) Validate() error {
	if this.ContractRootReq != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContractRootReq); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContractRootReq", err)
		}
	}
	return nil
}
func (this *OperNodePowerRootReq_ContractRootReq) Validate() error {
	if this.TcpCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TcpCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TcpCont", err)
		}
	}
	if this.SvcCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SvcCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SvcCont", err)
		}
	}
	return nil
}
func (this *OperNodePowerRootReq_SvcContReq) Validate() error {
	if this.RequestObject != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RequestObject); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RequestObject", err)
		}
	}
	return nil
}
func (this *OperNodePowerRootReq_OperNodePowerReq) Validate() error {
	for _, item := range this.NodePowerOpers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodePowerOpers", err)
			}
		}
	}
	return nil
}
func (this *OperNodePowerRootReq_NodePowerOper) Validate() error {
	return nil
}
func (this *OperNodePowerRootRes) Validate() error {
	if this.ContractRootRes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContractRootRes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContractRootRes", err)
		}
	}
	return nil
}
func (this *OperNodePowerRootRes_ContractRootRes) Validate() error {
	if this.TcpCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TcpCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TcpCont", err)
		}
	}
	if this.SvcCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SvcCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SvcCont", err)
		}
	}
	return nil
}
func (this *OperNodePowerRootRes_SvcContRes) Validate() error {
	if this.ResultObject != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResultObject); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResultObject", err)
		}
	}
	return nil
}
func (this *OperNodePowerRootRes_OperNodePowerRsp) Validate() error {
	for _, item := range this.NodePowerRsps {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodePowerRsps", err)
			}
		}
	}
	return nil
}
func (this *OperNodePowerRootRes_NodePowerRsp) Validate() error {
	return nil
}
func (this *CreateNodesRootReq) Validate() error {
	if this.ContractRootReq != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContractRootReq); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContractRootReq", err)
		}
	}
	return nil
}
func (this *CreateNodesRootReq_ContractRootReq) Validate() error {
	if this.TcpCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TcpCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TcpCont", err)
		}
	}
	if this.SvcCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SvcCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SvcCont", err)
		}
	}
	return nil
}
func (this *CreateNodesRootReq_SvcContReq) Validate() error {
	if this.RequestObject != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RequestObject); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RequestObject", err)
		}
	}
	return nil
}
func (this *CreateNodesRootReq_CreateNodeInfoReq) Validate() error {
	for _, item := range this.NodeInfoReqs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodeInfoReqs", err)
			}
		}
	}
	if this.PageInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PageInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PageInfo", err)
		}
	}
	return nil
}
func (this *CreateNodesRootReq_NodeInfoReq) Validate() error {
	return nil
}
func (this *CreateNodesRootRes) Validate() error {
	if this.ContractRootRes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContractRootRes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContractRootRes", err)
		}
	}
	return nil
}
func (this *CreateNodesRootRes_ContractRootRes) Validate() error {
	if this.TcpCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TcpCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TcpCont", err)
		}
	}
	if this.SvcCont != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SvcCont); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SvcCont", err)
		}
	}
	return nil
}
func (this *CreateNodesRootRes_SvcContRes) Validate() error {
	if this.ResultObject != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResultObject); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResultObject", err)
		}
	}
	return nil
}
func (this *CreateNodesRootRes_CreateNodeInfoRsp) Validate() error {
	for _, item := range this.NodeInfoRsps {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodeInfoRsps", err)
			}
		}
	}
	if this.PageInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PageInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PageInfo", err)
		}
	}
	return nil
}
func (this *CreateNodesRootRes_NodeInfoRsp) Validate() error {
	return nil
}