package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list types of disks
type ListTypesRequest struct {
	// Show detailed disk types by seps
	// Required: true
	Detailed bool `url:"detailed" json:"detailed" validate:"required"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListTypes gets list defined disk types
func (d Disks) ListTypes(ctx context.Context, req ListTypesRequest) (*ListTypes, error) {
	url := "/cloudapi/disks/listTypes"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListTypes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
