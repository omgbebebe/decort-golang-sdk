package locations

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct to get list of locations
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

// List gets list of all locations as a ListLocations struct
func (l Locations) List(ctx context.Context, req ListRequest) (*ListLocations, error) {
	res, err := l.ListRaw(ctx, req)
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

// ListRaw gets list of all locations as an array of bytes
func (l Locations) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudapi/locations/list"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
