package vgpu

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for getting list of VGPU
type ListRequest struct {
	// Find by id
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by vgpu status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Find by vgpu type
	// Required: false
	Type string `url:"type,omitempty" json:"type,omitempty"`

	// Find by vgpu mode
	// Required: false
	Mode string `url:"mode,omitempty" json:"mode,omitempty"`

	// Find by id resgroup
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Find by account id
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Find by compute id
	// Required: false
	ComputeID uint64 `url:"computeId,omitempty" json:"computeId,omitempty"`

	// Find by pgpu id
	// Required: false
	PGPUID uint64 `url:"pgpuId,omitempty" json:"pgpuId,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list all VGPU
func (v VGPU) List(ctx context.Context, req ListRequest) (*ListVGPU, error) {
	url := "/cloudbroker/vgpu/list"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListVGPU{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
