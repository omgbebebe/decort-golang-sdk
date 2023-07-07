package lb

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of deleted load balancers
type ListDeletedRequest struct {
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

	// Find by frontend Ip
	// Required: false
	FrontIP string `url:"frontIp,omitempty" json:"frontIp,omitempty"`

	// Find by backend Ip
	// Required: false
	BackIP string `url:"backIp,omitempty" json:"backIp,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: true
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListDeleted gets list of deleted load balancers
func (l LB) ListDeleted(ctx context.Context, req ListDeletedRequest) (*ListLB, error) {
	url := "/cloudapi/lb/listDeleted"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
