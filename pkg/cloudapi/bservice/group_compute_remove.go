package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for remove group compute
type GroupComputeRemoveRequest struct {
	// ID of the Basic Service
	// Required: true
    ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// ID of the Compute GROUP
	// Required: true
	CompGroupID uint64 `url:"compgroupId" json:"compgroupId" validate:"required"`

	// ID of the Compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// GroupComputeRemove makes group compute remove of the Basic Service
func (b BService) GroupComputeRemove(ctx context.Context, req GroupComputeRemoveRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/groupComputeRemove"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
