package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for unpin from stack
type UnpinFromStackRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId"`
}

func (crq UnpinFromStackRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}

	return nil
}

// UnpinFromStack unpin compute from current stack
func (c Compute) UnpinFromStack(ctx context.Context, req UnpinFromStackRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/compute/unpinFromStack"

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
