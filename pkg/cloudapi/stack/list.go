package stack

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct to get list of stacks
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

// List gets list of stacks as a ListStacks struct
func (i Stack) List(ctx context.Context, req ListRequest) (*ListStacks, error) {
	res, err := i.ListRaw(ctx, req)
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

// ListRaw gets list of stacks as an array of bytes
func (i Stack) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudapi/stack/list"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
