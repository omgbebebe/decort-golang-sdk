package kvmppc

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for mass create KVM PowerPC
type MassCreateRequest struct {
	// ID of the resource group, which will own this VM
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Name of this VM.
	// Must be unique among all VMs (including those in DELETED state) in target resource group
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Number of VMs
	// Required: true
	Count uint64 `url:"count" json:"count" validate:"required"`

	// Number CPUs to allocate to this VM
	// Required: true
	CPU uint64 `url:"cpu" json:"cpu" validate:"required"`

	// Volume of RAM in MB to allocate to this VM
	// Required: true
	RAM uint64 `url:"ram" json:"ram" validate:"required"`

	// Image ID
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

	// Slice of structs with net interface description.
	// If not specified, compute will be created with default interface from RG.
	// To create compute without interfaces, pass initialized empty slice.
	// Required: false
	Interfaces []Interface `url:"-" json:"interfaces,omitempty" validate:"omitempty,dive"`

	// Input data for cloud-init facility
	// Required: false
	UserData string `url:"userdata,omitempty" json:"userdata,omitempty"`

	// Text description of this VM
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Start after create of not
	// Required: false
	Start bool `url:"start" json:"start"`

	// Reason to action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

type wrapperMassCreateRequest struct {
	MassCreateRequest
	Interfaces []string `url:"interfaces,omitempty"`
}

// MassCreate creates KVM PPC computes based on specified OS image
func (k KVMPPC) MassCreate(ctx context.Context, req MassCreateRequest) ([]uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	var interfaces []string

	if req.Interfaces != nil && len(req.Interfaces) != 0 {
		interfaces = make([]string, 0, len(req.Interfaces))

		for i := range req.Interfaces {
			b, err := json.Marshal(req.Interfaces[i])
			if err != nil {
				return nil, err
			}

			interfaces = append(interfaces, string(b))
		}
	} else if req.Interfaces != nil && len(req.Interfaces) == 0 {
		interfaces = []string{"[]"}
	}

	reqWrapped := wrapperMassCreateRequest{
		MassCreateRequest: req,
		Interfaces:        interfaces,
	}

	url := "/cloudbroker/kvmppc/massCreate"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
	if err != nil {
		return nil, err
	}

	computes := make([]uint64, 0)

	err = json.Unmarshal(res, &computes)
	if err != nil {
		return nil, err
	}

	return computes, nil
}
