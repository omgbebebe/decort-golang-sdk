package rg

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create resource group
type CreateRequest struct {
	// Account, which will own this resource group
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Grid ID
	// Required: true
	GID uint64 `url:"gid" json:"gid" validate:"required"`

	// Name of this resource group. Must be unique within the account
	// Required: true
	Name string `url:"name" json:"name" validate:"required,min=2"`

	// Max size of memory in MB
	// Required: false
	MaxMemoryCapacity int64 `url:"maxMemoryCapacity,omitempty" json:"maxMemoryCapacity,omitempty"`

	// Max size of aggregated virtual disks in GB
	// Required: false
	MaxVDiskCapacity int64 `url:"maxVDiskCapacity,omitempty" json:"maxVDiskCapacity,omitempty"`

	// Max number of CPU cores
	// Required: false
	MaxCPUCapacity int64 `url:"maxCPUCapacity,omitempty" json:"maxCPUCapacity,omitempty"`

	// Max sent/received network transfer peering
	// Required: false
	MaxNetworkPeerTransfer int64 `url:"maxNetworkPeerTransfer,omitempty" json:"maxNetworkPeerTransfer,omitempty"`

	// Max number of assigned public IPs
	// Required: false
	MaxNumPublicIP int64 `url:"maxNumPublicIP,omitempty" json:"maxNumPublicIP,omitempty"`

	// Username - owner of this resource group.
	// Leave blank to set current user as owner
	// Required: false
	Owner string `url:"owner,omitempty" json:"owner,omitempty"`

	// Type of the default network for this resource group.
	// virtual machines created in this resource group will be by default connected to this network.
	// Allowed values are:
	//	- PRIVATE
	//	- PUBLIC
	//	- NONE
	// Required: false
	DefNet string `url:"def_net,omitempty" json:"def_net,omitempty" validate:"omitempty,rgDefNet"`

	// Private network IP CIDR if default network PRIVATE
	// Required: false
	IPCIDR string `url:"ipcidr,omitempty" json:"ipcidr,omitempty"`

	// Text description of this resource group
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`

	// External network ID
	// Required: false
	ExtNetID uint64 `url:"extNetId,omitempty" json:"extNetId,omitempty"`

	// External IP address
	// Required: false
	ExtIP string `url:"extIp,omitempty" json:"extIp,omitempty"`

	// Register computes in registration system
	// Required: false
	RegisterComputes bool `url:"registerComputes,omitempty" json:"registerComputes,omitempty"`
}

// Create creates resource group
func (r RG) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/rg/create"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
