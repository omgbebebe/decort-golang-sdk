package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for attach GPU for compute
type AttachGPURequest struct {
	// Identifier compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Identifier vGPU
	// Required: true
	VGPUID uint64 `url:"vgpuId" json:"vgpuId" validate:"required"`
}

// AttachGPU attach GPU for compute, returns vGPU ID on success
func (c Compute) AttachGPU(ctx context.Context, req AttachGPURequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/attachGpu"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
