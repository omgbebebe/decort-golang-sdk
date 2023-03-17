package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update VINS settings
type GroupUpdateVINSRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// ID of the Compute Group
	// Required: true
	CompGroupID uint64 `url:"compgroupId" json:"compgroupId" validate:"required"`

	// List of ViNSes to connect computes
	// Required: false
	VINSes []uint64 `url:"vinses,omitempty" json:"vinses,omitempty"`
}

// GroupUpdateVINS update ViNS settings for the group according to the new list
func (b BService) GroupUpdateVINS(ctx context.Context, req GroupUpdateVINSRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/groupUpdateVins"

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
