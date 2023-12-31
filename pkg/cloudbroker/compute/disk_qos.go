package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for change QOS of the disk
type DiskQOSRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// ID of the disk to apply limits
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// Limit IO for a certain disk total and read/write options are not allowed to be combined
	// Required: true
	Limits string `url:"limits" json:"limits" validate:"required"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// DiskQOS change QOS of the disk
func (c Compute) DiskQOS(ctx context.Context, req DiskQOSRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/diskQos"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
