package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)
// Request struct for setting customFields values for the Compute
type SetCustomFieldsRequest struct {
	// ID of the compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Custom fields for Compute. Must be dict.
	// Required: true
	CustomFields string `url:"customFields" json:"customFields" validate:"required"`
}

// SetCustomFields sets customFields values for the Compute
func (c Compute) SetCustomFields(ctx context.Context, req SetCustomFieldsRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/setCustomFields"

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
