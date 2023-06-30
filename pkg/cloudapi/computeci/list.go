package computeci

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of computeci
type ListRequest struct {
	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by computeci ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by drivers
	// Find by computeci ID
	Drivers []string `url:"drivers,omitempty" json:"drivers,omitempty"`

	// If true list deleted instances as well
	// Required: false
	IncludeDeleted bool `url:"includeDeleted,omitempty" json:"includeDeleted,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of computeci instances
func (c ComputeCI) List(ctx context.Context, req ListRequest) (*ListComputeCI, error) {
	url := "/cloudapi/computeci/list"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListComputeCI{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
