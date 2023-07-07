package extnet

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list external network
type ListRequest struct {
	// Find by account ID
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by network ip address
	// Required: false
	Network string `url:"network,omitempty" json:"network,omitempty"`

	// Find by vlan ID
	// Required: false
	VLANID uint64 `url:"vlanId,omitempty" json:"vlanId,omitempty"`

	// Find by vnfDevices ID
	// Required: false
	VNFDevID uint64 `url:"vnfDevId,omitempty" json:"vnfDevId,omitempty"`

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

// List gets list all available external networks
func (e ExtNet) List(ctx context.Context, req ListRequest) (*ListExtNet, error) {
	url := "/cloudbroker/extnet/list"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListExtNet{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
