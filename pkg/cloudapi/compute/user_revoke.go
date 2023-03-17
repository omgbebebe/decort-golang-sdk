package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for revoke user access
type UserRevokeRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Name of the user to remove
	// Required: true
	Username string `url:"userName" json:"userName" validate:"required"`
}

// UserRevoke revokes user access to the compute
func (c Compute) UserRevoke(ctx context.Context, req UserRevokeRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/userRevoke"

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
