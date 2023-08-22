package stack

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list stack
type ListRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

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
}

// ListStacks gets list stack
func (i Stack) List(ctx context.Context, req ListRequest) (*ListStacks, error) {
	url := "/cloudbroker/stack/list"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListStacks{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
