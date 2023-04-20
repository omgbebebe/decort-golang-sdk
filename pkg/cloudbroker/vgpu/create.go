package vgpu

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
)

// Request struct for creating VGPU
type CreateRequest struct {
	// ID of pGPU
	// Required: true
	PGPUID uint64 `url:"pgpuId" json:"pgpuId" validate:"required"`

	// ID of the target resource group.
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Virtual profile id
	// Required: false
	ProfileID uint64 `url:"profileId,omitempty" json:"profileId,omitempty"`

	// Allocate vgpu after creation
	// Required: false
	Allocate bool `url:"allocate,omitempty" json:"allocate,omitempty"`
}

// Create creates VGPU
func (v VGPU) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vgpu/create"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
