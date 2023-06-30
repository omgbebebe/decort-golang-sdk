package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list unattached disk
type ListUnattachedRequest struct {
	// Find by id
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by account name
	// Required: false
	AccountName string `url:"accountName,omitempty" json:"accountName,omitempty"`

	// Find by max size disk
	// Required: false
	DiskMaxSize int64 `url:"diskMaxSize,omitempty" json:"diskMaxSize,omitempty"`

	// Find by status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Find by shared, true or false
	// Required: false
	Shared bool `url:"shared,omitempty" json:"shared,omitempty"`

	// Type of the disks
	// Required: false
	Type string `url:"type,omitempty" json:"type,omitempty"`

	// ID of the account
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListUnattached gets list of unattached disks
func (d Disks) ListUnattached(ctx context.Context, req ListUnattachedRequest) (*ListDisksUnattached, error) {
	url := "/cloudapi/disks/listUnattached"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListDisksUnattached{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
