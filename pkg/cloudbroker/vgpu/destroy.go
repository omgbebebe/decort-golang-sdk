package vgpu

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
)

// Request for destroying VGPU
type DestroyRequest struct {
	// Virtual GPU ID
	// Required: true
	VGPUID uint64 `url:"vgpuId" json:"vgpuId" validate:"required"`

	// Force delete (deallocate and detach from compute)
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`
}

// Destroy destroys VGPU
func (v VGPU) Destroy(ctx context.Context, req DestroyRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vgpu/destroy"

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
