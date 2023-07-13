package vins

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for search VINSes
type SearchRequest struct {
	// ID of the account to search for the ViNSes
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// ID of the resource group to limit search to the specified RG level only
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Name of the ViNS to search for
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// If False, then VINSes having one of the statuses are not listed for
	// Required: false
	ShowAll bool `url:"show_all,omitempty" json:"show_all,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Search  search VINSes
func (v VINS) Search(ctx context.Context, req SearchRequest) (SearchVINS, error) {
	url := "/cloudbroker/vins/search"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := SearchVINS{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
