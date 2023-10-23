package account

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct to get list of accounts
type ListRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by access control list
	// Required: false
	ACL string `url:"acl,omitempty" json:"acl,omitempty"`

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

// List gets list of all accounts the user has access to as a ListAccounts struct
func (a Account) List(ctx context.Context, req ListRequest) (*ListAccounts, error) {
	res, err := a.ListRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	list := ListAccounts{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// ListRaw gets list of all accounts the user has access to as an array of bytes
func (a Account) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudbroker/account/list"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
