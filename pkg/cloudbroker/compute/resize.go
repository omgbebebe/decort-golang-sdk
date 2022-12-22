package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for resize compute
type ResizeRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`

	// New CPU count.
	// Pass 0 if no change to CPU count is required
	// Required: false
	Force bool `url:"force,omitempty"`

	// New RAM volume in MB.
	// Pass 0 if no change to RAM volume is required
	// Required: false
	CPU uint64 `url:"cpu,omitempty"`

	// Force compute resize
	// Required: false
	RAM uint64 `url:"ram,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}

func (crq ResizeRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// Resize resize compute instance
func (c Compute) Resize(ctx context.Context, req ResizeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
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
