package rg

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of resource groups
type ListRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by account ID
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Find by name account
	// Required: false
	AccountName string `url:"accountName,omitempty" json:"accountName,omitempty"`

	// Find by created after time (unix timestamp)
	// Required: false
	CreatedAfter uint64 `url:"createdAfter,omitempty" json:"createdAfter,omitempty"`

	// Find by created before time (unix timestamp)
	// Required: false
	CreatedBefore uint64 `url:"createdBefore,omitempty" json:"createdBefore,omitempty"`

	// Find by status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Included deleted resource groups
	// Required: false
	IncludeDeleted bool `url:"includedeleted,omitempty" json:"includedeleted,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of all resource groups the user has access to
func (r RG) List(ctx context.Context, req ListRequest) (*ListRG, error) {
	url := "/cloudbroker/rg/list"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListRG{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
