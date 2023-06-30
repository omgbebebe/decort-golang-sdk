package bservice

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of deleted BasicService instances
type ListDeletedRequest struct {
	// ID of the account to query for BasicService instances
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// ID of the resource group to query for BasicService instances
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListDeleted gets list deleted BasicService instances associated with the specified Resource Group
func (b BService) ListDeleted(ctx context.Context, req ListDeletedRequest) (*ListBasicServices, error) {
	url := "/cloudapi/bservice/listDeleted"

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
