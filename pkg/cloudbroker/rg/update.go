package rg

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update resource group
type UpdateRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// New name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// New description
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

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

	// Register computes in registration system
	// Required: false
	RegisterComputes bool `url:"registerComputes,omitempty" json:"registerComputes,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`

	// List of strings with pools i.e.: ["sep1_poolName1", "sep2_poolName2", etc]
	// Required: false
	UniqPools []string `url:"uniqPools,omitempty" json:"uniqPools,omitempty"`
}

// Update updates resource group
func (r RG) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/rg/update"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
