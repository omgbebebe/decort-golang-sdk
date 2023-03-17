package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for detach networ to compute
type NetDetachRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// IP of the network interface
	// Required: false
	IPAddr string `url:"ipAddr,omitempty" json:"ipAddr,omitempty"`

	// MAC of the network interface
	// Required: false
	MAC string `url:"mac,omitempty" json:"mac,omitempty"`
}

// NetDetach detach network to compute
func (c Compute) NetDetach(ctx context.Context, req NetDetachRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/netDetach"

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
