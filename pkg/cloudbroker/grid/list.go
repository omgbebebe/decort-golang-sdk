package grid

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct to get list of locations
type ListRequest struct {
	// Find by id grid
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name grid
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of all locations as a ListGrids struct
func (g Grid) List(ctx context.Context, req ListRequest) (*ListGrids, error) {
	res, err := g.ListRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	list := ListGrids{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// ListRaw gets list of all locations as an array of bytes
func (g Grid) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudbroker/grid/list"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
