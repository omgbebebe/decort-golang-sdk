package rg

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get detailed information about resource group
type GetRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Get gets current configuration of the resource group as a RecordResourceGroup struct
func (r RG) Get(ctx context.Context, req GetRequest) (*RecordResourceGroup, error) {
	res, err := r.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordResourceGroup{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets current configuration of the resource group as an array of bytes
func (r RG) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/rg/get"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
