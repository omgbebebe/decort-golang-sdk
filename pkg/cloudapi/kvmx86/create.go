package kvmx86

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

type Interface struct {
	// Network type
	// Should be one of:
	//	- VINS
	//	- EXTNET
	NetType string `url:"netType" json:"netType" validate:"required,kvmNetType"`

	// Network ID for connect to,
	// for EXTNET - external network ID,
	// for VINS - VINS ID,
	NetID uint64 `url:"netId" json:"netId" validate:"required"`

	// IP address to assign to this VM when connecting to the specified network
	// Required: false
	IPAddr string `url:"ipAddr,omitempty" json:"ipAddr,omitempty"`
}

// Request struct for create KVM x86 VM
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
	SepID uint64 `url:"sepId,omitempty" json:"sepId,omitempty"`

	// Pool to use if sepId is set, can be also empty if needed to be chosen by system
	// Required: false
	Pool string `url:"pool,omitempty" json:"pool,omitempty"`

	// Slice of structs with net interface description.
	// If not specified, compute will be created with default interface from RG.
	// To create compute without interfaces, pass initialized empty slice .
	// Required: false
	Interfaces []Interface `url:"-" json:"interfaces,omitempty" validate:"omitempty,dive"`

	// Input data for cloud-init facility
	// Required: false
	Userdata string `url:"userdata,omitempty" json:"userdata,omitempty"`

	// Text description of this VM
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Start VM upon success
	// Required: false
	Start bool `url:"start" json:"start"`

	// System name
	// Required: false
	IS string `url:"IS,omitempty" json:"IS,omitempty"`

	// Compute purpose
	// Required: false
	IPAType string `url:"ipaType,omitempty" json:"ipaType,omitempty"`

	// Custom fields for compute. Must be a dict
	// Required: false
	CustomFields string `url:"customFields,omitempty" json:"customFields,omitempty"`

	// Type of compute Stateful (KVM_X86) or Stateless (SVA_KVM_X86)
	// Required: false
	Driver string `url:"driver,omitempty" json:"driver,omitempty" validate:"omitempty,computeDriver"`
}

type wrapperCreateRequest struct {
	CreateRequest
	Interfaces []string `url:"interfaces,omitempty"`
}

// Create creates KVM x86 VM based on specified OS image
func (k KVMX86) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	var interfaces []string

	if req.Interfaces != nil && len(req.Interfaces) != 0 {
		interfaces = make([]string, 0, len(req.Interfaces))

		for i := range req.Interfaces {
			b, err := json.Marshal(req.Interfaces[i])
			if err != nil {
				return 0, err
			}

			interfaces = append(interfaces, string(b))
		}
	} else if req.Interfaces != nil && len(req.Interfaces) == 0 {
		interfaces = []string{"[]"}
	}

	reqWrapped := wrapperCreateRequest{
		CreateRequest: req,
		Interfaces:    interfaces,
	}

	url := "/cloudapi/kvmx86/create"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
