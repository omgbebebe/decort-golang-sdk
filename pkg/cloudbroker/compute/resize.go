package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for resize compute
type ResizeRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// New CPU count.
	// Pass 0 if no change to CPU count is required
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`

	// New RAM volume in MB.
	// Pass 0 if no change to RAM volume is required
	// Required: false
	CPU uint64 `url:"cpu,omitempty" json:"cpu,omitempty"`

	// Force compute resize
	// Required: false
	RAM uint64 `url:"ram,omitempty" json:"ram,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Resize resize compute instance
func (c Compute) Resize(ctx context.Context, req ResizeRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/resize"

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
