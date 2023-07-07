package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list GPU for compute
type ListVGPURequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Find by GPU id
	// Required: false
	GPUID uint64 `url:"gpuId,omitempty" json:"gpuId,omitempty"`

	// Find by type
	// Required: false
	Type string `url:"type,omitempty" json:"type,omitempty"`

	// Find by status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`

	// Include deleted computes. If using field 'status', then includedeleted will be ignored
	// Required: false
	IncludeDeleted bool `url:"includedeleted,omitempty" json:"includedeleted,omitempty"`
}

// ListVGPU gets list GPU for compute
func (c Compute) ListVGPU(ctx context.Context, req ListVGPURequest) (*ListVGPUs, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/listVGpu"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListVGPUs{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
