package kvmppc

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create KVM PowerPC VM from scratch
type CreateBlankRequest struct {
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

	// Size of the boot disk in GB
	// Required: true
	BootDisk uint64 `url:"bootDisk" json:"bootDisk" validate:"required"`

	// ID of SEP to create boot disk on
	// Uses images SEP ID if not set
	// Required: true
	SEPID uint64 `url:"sepId" json:"sepId" validate:"required"`

	// Pool to use if SEP ID is set, can be also empty if needed to be chosen by system
	// Required: true
	Pool string `url:"pool" json:"pool" validate:"required"`

	// Slice of structs with net interface description
	// Required: false
	Interfaces []Interface `url:"interfaces,omitempty" json:"interfaces,omitempty" validate:"omitempty,min=1,dive"`

	// Text description of this VM
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

// CreateBlank creates KVM PowerPC VM from scratch
func (k KVMPPC) CreateBlank(ctx context.Context, req CreateBlankRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/kvmppc/createBlank"

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
