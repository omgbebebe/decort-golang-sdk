package rg

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get detailed information about resource consumption for ResGroup
type GetResourceConsumptionRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`
}

// GetResourceConsumption gets resource consumption of the resource group
func (r RG) GetResourceConsumption(ctx context.Context, req GetResourceConsumptionRequest) (*ItemResourceConsumption, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/rg/getResourceConsumption"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := ItemResourceConsumption{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
