package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list vGPU
type ListVGPURequest struct {
	// Identifier compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// ListVGPU gets list vGPU
func (c Compute) ListVGPU(ctx context.Context, req ListVGPURequest) ([]interface{}, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/listVgpu"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := []interface{}{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
