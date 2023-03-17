package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for stop the specified Compute Group
type GroupStopRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
    ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// ID of the Compute Group to stop
	// Required: true
    CompGroupID uint64 `url:"compgroupId" json:"compgroupId" validate:"required"`

	// Force stop Compute Group
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`
}

// GroupStop stops the specified Compute Group within BasicService
func (b BService) GroupStop(ctx context.Context, req GroupStopRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/groupStop"

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
