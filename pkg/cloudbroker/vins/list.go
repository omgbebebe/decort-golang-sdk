package vins

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct to get list of VINSes
type ListRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by account ID
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Find by resource group id
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Find by external network IP
	// Required: false
	ExtIP string `url:"extIp,omitempty" json:"extIp,omitempty"`

	// Include deleted
	// Required: false
	IncludeDeleted bool `url:"includeDeleted,omitempty" json:"includeDeleted,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of VINSes as a ListVINS struct
func (v VINS) List(ctx context.Context, req ListRequest) (*ListVINS, error) {
	res, err := v.ListRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	list := ListVINS{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// ListRaw gets list of VINSes as an array of bytes
func (v VINS) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudbroker/vins/list"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
