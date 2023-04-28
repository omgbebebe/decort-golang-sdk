package compute

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
)

// Request struct for changing link state
type ChangeLinkStateRequest struct {
	// Compute ID
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Interface name or MAC address
	// Required: true
	Interface string `url:"interface" json:"interface" validate:"required"`

	// Interface state
	// Must be either "on" or "off"
	// Required: true
	State string `url:"state" json:"state" validate:"required,interfaceState"`
}

// ChangeLinkState changes the status link virtual of compute
func (c Compute) ChangeLinkState(ctx context.Context, req ChangeLinkStateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/changeLinkState"

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
