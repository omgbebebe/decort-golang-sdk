package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for detach VGPU for compute
type DetachGPURequest struct {
	// Identifier compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Identifier virtual GPU
	// Required: false
	VGPUID int64 `url:"vgpuId,omitempty" json:"vgpuId,omitempty"`
}

// DetachGPU detach VGPU for compute.
// If param VGPU ID is equivalent -1, then detach all VGPU for compute
func (c Compute) DetachGPU(ctx context.Context, req DetachGPURequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/detachGpu"

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
