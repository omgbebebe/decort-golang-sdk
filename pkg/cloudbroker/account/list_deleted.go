package account

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list deleted accounts
type ListDeletedRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by access control list
	// Required: false
	ACL string `url:"acl,omitempty" json:"acl,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListDeleted gets list all deleted accounts the user has access to
func (a Account) ListDeleted(ctx context.Context, req ListDeletedRequest) (*ListAccounts, error) {
	url := "/cloudbroker/account/listDeleted"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
