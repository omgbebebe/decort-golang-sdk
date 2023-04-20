package vgpu

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
)

// Request for allocating VGPU
type AllocateRequest struct {
	// Virtual GPU ID
	// Required: true
	VGPUID uint64 `url:"vgpuId" json:"vgpuId" validate:"required"`
}

// Allocate allocates GPU
func (v VGPU) Allocate(ctx context.Context, req AllocateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vgpu/allocate"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
