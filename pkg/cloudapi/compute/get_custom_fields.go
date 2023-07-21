package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for getting Compute's customFields
type GetCustomFieldsRequest struct {
	// Compute ID
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// GetCustomFields gets Compute's customFields
func (c Compute) GetCustomFields(ctx context.Context, req GetCustomFieldsRequest) (interface{}, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/getCustomFields"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	var info interface{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
