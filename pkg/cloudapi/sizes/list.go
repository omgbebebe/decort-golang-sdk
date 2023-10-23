package sizes

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct for list of the available flavors
type ListRequest struct {
	// ID of the cloudspace
	// Required: false
	CloudspaceID uint64 `url:"cloudspaceId,omitempty" json:"cloudspaceId,omitempty"`

	// Location code for the sizes
	// Required: false
	Location string `url:"location,omitempty" json:"location,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of the available flavors as a ListSizes struct, filtering can be based on the user which is doing the request
func (s Sizes) List(ctx context.Context, req ListRequest) (*ListSizes, error) {
	res, err := s.ListRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	list := &ListSizes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// ListRaw gets list of the available flavors as an array of bytes
func (s Sizes) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudapi/sizes/list"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
