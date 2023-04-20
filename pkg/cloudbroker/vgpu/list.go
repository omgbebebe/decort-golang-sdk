package vgpu

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for getting list of VGPU
type ListRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list all VGPU
func (v VGPU) List(ctx context.Context, req ListRequest) (ListVGPU, error) {
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

	return list, nil
}
