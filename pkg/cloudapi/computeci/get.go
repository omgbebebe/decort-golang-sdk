package computeci

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for information about computeci
type GetRequest struct {
	// ID of the Compute CI
	// Required: true
	ComputeCIID uint64 `url:"computeciId" json:"computeciId" validate:"required"`
}

// Get gets information about computeci by ID
func (c ComputeCI) Get(ctx context.Context, req GetRequest) (*ItemComputeCI, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validatonError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validatonError)
		}
	}

	url := "/cloudapi/computeci/get"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := ItemComputeCI{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
