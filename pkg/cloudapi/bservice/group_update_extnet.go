package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update External Network settings
type GroupUpdateExtNetRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// ID of the Compute Group
	// Required: true
	CompGroupID uint64 `url:"compgroupId" json:"compgroupId" validate:"required"`

	// List of Extnets to connect computes
	// Required: false
	ExtNets []uint64 `url:"extnets,omitempty" json:"extnets,omitempty"`
}

// GroupUpdateExtNet updates External Network settings for the group according to the new list
func (b BService) GroupUpdateExtNet(ctx context.Context, req GroupUpdateExtNetRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/groupUpdateExtnet"

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
