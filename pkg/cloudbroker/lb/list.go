package lb

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of load balancers
type ListRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by account ID
	// Required: false
	AccountID uint64 `url:"accountID,omitempty" json:"accountID,omitempty"`

	// Find by resource group ID
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Find by tech status
	// Required: false
	TechStatus string `url:"techStatus,omitempty" json:"techStatus,omitempty"`

	// Find by status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Find by frontend Ip
	// Required: false
	FrontIP string `url:"frontIp,omitempty" json:"frontIp,omitempty"`

	// Find by backend Ip
	// Required: false
	BackIP string `url:"backIp,omitempty" json:"backIp,omitempty"`

	// Included deleted load balancers
	// Required: false
	IncludeDeleted bool `url:"includedeleted,omitempty" json:"includedeleted,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list all load balancers
func (lb LB) List(ctx context.Context, req ListRequest) (*ListLB, error) {
	url := "/cloudbroker/lb/list"

	res, err := lb.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListLB{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil

}
