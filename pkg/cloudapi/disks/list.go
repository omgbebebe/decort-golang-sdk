package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct to get list of disks
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

	// Find by sep ID
	// Required: false
	SEPID uint64 `url:"sepId,omitempty" json:"sepId,omitempty"`

	// Find by pool name
	// Required: false
	Pool string `url:"pool,omitempty" json:"pool,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of the created disks belonging to an account as a ListDisks struct
func (d Disks) List(ctx context.Context, req ListRequest) (*ListDisks, error) {
	res, err := d.ListRaw(ctx, req)
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

// ListRaw gets list of the created disks belonging to an account as an array of bytes
func (d Disks) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudapi/disks/list"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
