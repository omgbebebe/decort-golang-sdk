package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for detach and delete disk from compute
type DiskDelRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// ID of disk instance
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// False if disk is to be deleted to recycle bin
	// Required: true
	Permanently bool `url:"permanently" json:"permanently" validate:"required"`
}

// DiskDel delete disk and detach from compute
func (c Compute) DiskDel(ctx context.Context, req DiskDelRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/diskDel"

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
