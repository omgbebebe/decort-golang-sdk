package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for move compute new resource group
type MoveToRGRequest struct {
	// ID of the compute instance to move
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// ID of the target resource group
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// New name for the compute upon successful move,
	// if name change required.
	// Pass empty string if no name change necessary
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Should the compute be restarted upon successful move
	// Required: false
	Autostart bool `url:"autostart,omitempty" json:"autostart,omitempty"`

	// By default moving compute in a running state is not allowed.
	// Set this flag to True to force stop running compute instance prior to move.
	// Required: false
	ForceStop bool `url:"forceStop,omitempty" json:"forceStop,omitempty"`
}

// MoveToRG moves compute instance to new resource group
func (c Compute) MoveToRG(ctx context.Context, req MoveToRGRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/moveToRg"

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
