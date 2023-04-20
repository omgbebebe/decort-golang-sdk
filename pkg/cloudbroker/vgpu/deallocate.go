package vgpu

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
)

// Request for deallocating VGPU
type DeallocateRequest struct {
	// Virtual GPU ID
	// Required: true
	VGPUID uint64 `url:"vgpuId" json:"vgpuId" validate:"required"`

	// Force delete (detach from compute)
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`
}

// Deallocate releases GPU resources
func (v VGPU) Deallocate(ctx context.Context, req DeallocateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vgpu/deallocate"

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
