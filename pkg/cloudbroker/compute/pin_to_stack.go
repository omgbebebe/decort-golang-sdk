package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for pin comptute to stack
type PinToStackRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Stack ID to pin to
	// Required: true
	TargetStackID uint64 `url:"targetStackId" json:"targetStackId" validate:"required"`

	// Try to migrate or not if compute in running states
	// Required: false
	Force bool `url:"force" json:"force"`
}

// PinToStack pin compute to current stack
func (c Compute) PinToStack(ctx context.Context, req PinToStackRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/pinToStack"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
