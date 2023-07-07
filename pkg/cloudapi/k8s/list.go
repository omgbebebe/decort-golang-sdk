package k8s

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list information K8S
type ListRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by IP address
	// Required: false
	IPAddress string `url:"ipAddress,omitempty" json:"ipAddress,omitempty"`

	// Find by resource group ID
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Find by lbId
	// Required: false
	LBID uint64 `url:"lbId,omitempty" json:"lbId,omitempty"`

	// Find by basicServiceId
	// Required: false
	BasicServiceID uint64 `url:"basicServiceId,omitempty" json:"basicServiceId,omitempty"`

	// Find by status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Find by techStatus
	// Required: false
	TechStatus string `url:"techStatus,omitempty" json:"techStatus,omitempty"`

	// Include deleted clusters in result
	// Required: false
	IncludeDeleted bool `url:"includedeleted,omitempty" json:"includedeleted,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list all kubernetes clusters the user has access to
func (k8s K8S) List(ctx context.Context, req ListRequest) (*ListK8SClusters, error) {
	url := "/cloudapi/k8s/list"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListK8SClusters{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
