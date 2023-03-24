package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update compute
type UpdateRequest struct {
	// ID of the compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// New name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// New description
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Update updates some properties of the compute
func (c Compute) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/update"

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
