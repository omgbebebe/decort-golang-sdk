package compute

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get deleted computes list
type ListDeletedRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by account ID
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Find by resource group name
	// Required: false
	RGName string `url:"rgName,omitempty" json:"rgName,omitempty"`

	// Find by resource group ID
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Find by tech status
	// Required: false
	TechStatus string `url:"techStatus,omitempty" json:"techStatus,omitempty"`

	// Find by IP address
	// Required: false
	IPAddress string `url:"ipAddress,omitempty" json:"ipAddress,omitempty"`

	// Find by external network name
	// Required: false
	ExtNetName string `url:"extNetName,omitempty" json:"extNetName,omitempty"`

	// Find by external network ID
	// Required: false
	ExtNetID uint64 `url:"extNetId,omitempty" json:"extNetId,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListDeleted gets list all deleted computes
func (c Compute) ListDeleted(ctx context.Context, req ListDeletedRequest) (*ListComputes, error) {
	url := "/cloudbroker/compute/listDeleted"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListComputes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
