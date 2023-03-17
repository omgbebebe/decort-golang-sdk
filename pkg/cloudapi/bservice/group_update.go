package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update existing Compute group
type GroupUpdateRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
    ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// ID of the Compute Group
	// Required: true
    CompGroupID uint64 `url:"compgroupId" json:"compgroupId" validate:"required"`

	// Specify non-empty string to update Compute Group name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Specify non-empty string to update group role
	// Required: false
	Role string `url:"role,omitempty" json:"role,omitempty"`

	// Specify positive value to set new compute CPU count
	// Required: false
	CPU uint64 `url:"cpu,omitempty" json:"cpu,omitempty"`

	// Specify positive value to set new compute RAM volume in MB
	// Required: false
	RAM uint64 `url:"ram,omitempty" json:"ram,omitempty"`

	// Specify new compute boot disk size in GB
	// Required: false
	Disk uint64 `url:"disk,omitempty" json:"disk,omitempty"`

	// Force resize Compute Group
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`
}

// GroupUpdate updates existing Compute group within Basic Service and apply new settings to its computes as necessary
func (b BService) GroupUpdate(ctx context.Context, req GroupUpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/groupUpdate"

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
