package locations

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of locations
type ListRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`

	// Find by flag
	// Required: false
	Flag string `url:"flag,omitempty" json:"flag,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by code location
	// Required: false
	LocationCode string `url:"locationCode,omitempty" json:"locationCode,omitempty"`
}

// List gets list all locations
func (l Locations) List(ctx context.Context, req ListRequest) (*ListLocations, error) {
	url := "/cloudapi/locations/list"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListLocations{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
