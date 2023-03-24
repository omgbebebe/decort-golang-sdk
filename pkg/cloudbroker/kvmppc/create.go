package kvmppc

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create KVM PowerPC VM
type CreateRequest struct {
	// ID of the resource group, which will own this VM
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Name of this VM.
	// Must be unique among all VMs (including those in DELETED state) in target resource group
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Number CPUs to allocate to this VM
	// Required: true
	CPU uint64 `url:"cpu" json:"cpu" validate:"required"`

	// Volume of RAM in MB to allocate to this VM
	// Required: true
	RAM uint64 `url:"ram" json:"ram" validate:"required"`

	// ID of the OS image to base this VM on;
	// Could be boot disk image or CD-ROM image
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`

	// Size of the boot disk in GB
	// Required: false
	BootDisk uint64 `url:"bootDisk,omitempty" json:"bootDisk,omitempty"`

	// ID of SEP to create boot disk on.
	// Uses images SEP ID if not set
	// Required: false
	SEPID uint64 `url:"sepId,omitempty" json:"sepId,omitempty"`

	// Pool to use if SEP ID is set, can be also empty if needed to be chosen by system
	// Required: false
	Pool string `url:"pool,omitempty" json:"pool,omitempty"`

	// Network type
	// Should be one of:
	//	- VINS
	//	- EXTNET
	//	- NONE
	// Required: false
	NetType string `url:"netType,omitempty" json:"netType,omitempty" validate:"omitempty,kvmNetType"`

	// Network ID for connect to,
	// for EXTNET - external network ID,
	// for VINS - VINS ID,
	// when network type is not "NONE"
	// Required: false
	NetID uint64 `url:"netId,omitempty" json:"netId,omitempty"`

	// IP address to assign to this VM when connecting to the specified network
	// Required: false
	IPAddr string `url:"ipAddr,omitempty" json:"ipAddr,omitempty"`

	// Input data for cloud-init facility
	// Required: false
	Userdata string `url:"userdata,omitempty" json:"userdata,omitempty"`

	// Text description of this VM
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Start VM upon success
	// Required: false
	Start bool `url:"start,omitempty" json:"start,omitempty"`

	// Stack ID
	// Required: false
	StackID uint64 `url:"stackId,omitempty" json:"stackId,omitempty"`

	// System name
	// Required: false
	IS string `url:"IS,omitempty" json:"IS,omitempty"`

	// Compute purpose
	// Required: false
	IPAType string `url:"ipaType,omitempty" json:"ipaType,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Create creates KVM PowerPC VM based on specified OS image
func (k KVMPPC) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/kvmppc/create"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
