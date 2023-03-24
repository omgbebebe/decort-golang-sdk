package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for migrate compute
type MigrateRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Particular Stack ID to migrate this compute to
	// Required: false
	TargetStackID uint64 `url:"targetStackId,omitempty" json:"targetStackId,omitempty"`

	// If live migration fails, destroy compute
	// on source node and recreate on the target
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Migrate migrates compute to another stack
func (c Compute) Migrate(ctx context.Context, req MigrateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/migrate"

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
