package kvmx86

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for create KVM x86 VM
type CreateRequest struct {
	// ID of the resource group, which will own this VM
	// Required: true
	RGID uint64 `url:"rgId"`

	// Name of this VM.
	// Must be unique among all VMs (including those in DELETED state) in target resource group
	// Required: true
	Name string `url:"name"`

	// Number CPUs to allocate to this VM
	// Required: true
	CPU uint64 `url:"cpu"`

	// Volume of RAM in MB to allocate to this VM
	// Required: true
	RAM uint64 `url:"ram"`

	// ID of the OS image to base this VM on;
	// Could be boot disk image or CD-ROM image
	// Required: true
	ImageID uint64 `url:"imageId"`

	// Size of the boot disk in GB
	// Required: false
	BootDisk uint64 `url:"bootDisk,omitempty"`

	// ID of SEP to create boot disk on.
	// Uses images SEP ID if not set
	// Required: false
	SepID uint64 `url:"sepId,omitempty"`

	// Pool to use if sepId is set, can be also empty if needed to be chosen by system
	// Required: false
	Pool string `url:"pool,omitempty"`

	// Network type
	// Should be one of:
	//	- VINS
	//	- EXTNET
	//	- NONE
	// Required: false
	NetType string `url:"netType,omitempty"`

	// Network ID for connect to,
	// for EXTNET - external network ID,
	// for VINS - VINS ID,
	// when network type is not "NONE"
	// Required: false
	NetID uint64 `url:"netId,omitempty"`

	// IP address to assign to this VM when connecting to the specified network
	// Required: false
	IPAddr string `url:"ipAddr,omitempty"`

	// Input data for cloud-init facility
	// Required: false
	Userdata string `url:"userdata,omitempty"`

	// Text description of this VM
	// Required: false
	Description string `url:"desc,omitempty"`

	// Start VM upon success
	// Required: false
	Start bool `url:"start,omitempty"`

	// System name
	// Required: false
	IS string `url:"IS,omitempty"`

	// Compute purpose
	// Required: false
	IPAType string `url:"ipaType,omitempty"`
}

func (krq CreateRequest) validate() error {
	if krq.RGID == 0 {
		return errors.New("validation-error: field RGID can not be empty or equal to 0")
	}
	if krq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}
	if krq.CPU == 0 {
		return errors.New("validation-error: field CPU can not be empty or equal to 0")
	}
	if krq.RAM == 0 {
		return errors.New("validation-error: field RAM can not be empty or equal to 0")
	}
	if krq.ImageID == 0 {
		return errors.New("validation-error: field ImageID can not be empty or equal to 0")
	}

	return nil
}

// Create creates KVM x86 VM based on specified OS image
func (k KVMX86) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudapi/kvmx86/create"

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
