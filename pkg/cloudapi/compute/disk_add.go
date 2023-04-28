package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create and attach disk to compute
type DiskAddRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Name for disk
	// Required: true
	DiskName string `url:"diskName" json:"diskName" validate:"required"`

	// Disk size in GB
	// Required: true
	Size uint64 `url:"size" json:"size" validate:"required"`

	// Storage endpoint provider ID
	// By default the same with boot disk
	// Required: false
	SepID uint64 `url:"sepId,omitempty" json:"sepId,omitempty"`

	// Type of the disk
	// Should be one of:
	//	- D
	//	- B
	// Required: false
	DiskType string `url:"diskType,omitempty" json:"diskType,omitempty" validate:"omitempty,computeDiskType"`

	// Pool name
	// By default will be chosen automatically
	// Required: false
	Pool string `url:"pool,omitempty" json:"pool,omitempty"`

	// Optional description
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Specify image id for create disk from template
	// Required: false
	ImageID uint64 `url:"imageId,omitempty" json:"imageId,omitempty"`
}

// DiskAdd creates new disk and attach to compute
func (c Compute) DiskAdd(ctx context.Context, req DiskAddRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/diskAdd"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
