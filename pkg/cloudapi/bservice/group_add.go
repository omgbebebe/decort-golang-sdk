package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create new compute group within BasicService
type GroupAddRequest struct {
	// ID of the Basic Service to add a group to
	// Required: true
    ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// Name of the Compute Group to add
	// Required: true
    Name string `url:"name" json:"name" validate:"required"`

	// Computes number. Defines how many computes must be there in the group
	// Required: true
    Count uint64 `url:"count" json:"count" validate:"required"`

	// Compute CPU number. All computes in the group have the same CPU count
	// Required: true
    CPU uint64 `url:"cpu" json:"cpu" validate:"required"`

	// Compute RAM volume in MB. All computes in the group have the same RAM volume
	// Required: true
    RAM uint64 `url:"ram" json:"ram" validate:"required"`

	// Compute boot disk size in GB
	// Required: true
    Disk uint64 `url:"disk" json:"disk" validate:"required"`

	// OS image ID to create computes from
	// Required: true
    ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`

	// Compute driver
	// should be one of:
	//	- KVM_X86
	//	- KVM_PPC
	// Required: true
    Driver string `url:"driver" json:"driver" validate:"driver"`

	// Storage endpoint provider ID
	// Required: false
	SEPID uint64 `url:"sepId,omitempty" json:"sepId,omitempty"`

	// Pool to use if sepId is set, can be also empty if needed to be chosen by system
	// Required: false
	SEPPool string `url:"sepPool,omitempty" json:"sepPool,omitempty"`

	// Group role tag. Can be empty string, does not have to be unique
	// Required: false
	Role string `url:"role,omitempty" json:"role,omitempty"`

	// List of ViNSes to connect computes to
	// Required: false
	VINSes []uint64 `url:"vinses,omitempty" json:"vinses,omitempty"`

	// List of external networks to connect computes to
	// Required: false
	ExtNets []uint64 `url:"extnets,omitempty" json:"extnets,omitempty"`

	// Time of Compute Group readiness
	// Required: false
	TimeoutStart uint64 `url:"timeoutStart,omitempty" json:"timeoutStart,omitempty"`
}

// GroupAdd creates new Compute Group within BasicService.
// Compute Group is NOT started automatically,
// so you need to explicitly start it
func (b BService) GroupAdd(ctx context.Context, req GroupAddRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/groupAdd"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
