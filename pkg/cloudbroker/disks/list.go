package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list/list_deleted of disks
type ListRequest struct {
	// Find by id
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

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

	// ID of the account the disks belong to
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Type of the disks
	// Required: false
	Type string `url:"type,omitempty" json:"type,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list the created disks belonging to an account
func (d Disks) List(ctx context.Context, req ListRequest) (*ListDisks, error) {
	url := "/cloudbroker/disks/list"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListDisks{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
