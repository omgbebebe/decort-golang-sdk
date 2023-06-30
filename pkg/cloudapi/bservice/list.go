package bservice

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list BasicService instances
type ListRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// ID of the account to query for BasicService instances
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Find by resource group name
	// Required: false
	RGName string `url:"rgName,omitempty" json:"rgName,omitempty"`

	// ID of the resource group to query for BasicService instances
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Find by tech status
	// Required: false
	TechStatus string `url:"techStatus,omitempty" json:"techStatus,omitempty"`

	// Find by status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Find by account name
	// Required: false
	AccountName string `url:"accountName,omitempty" json:"accountName,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list BasicService instances associated with the specified Resource Group
func (b BService) List(ctx context.Context, req ListRequest) (*ListBasicServices, error) {
	url := "/cloudapi/bservice/list"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListBasicServices{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
