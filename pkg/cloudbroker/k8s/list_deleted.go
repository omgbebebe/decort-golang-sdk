package k8s

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list deleted kubernetes cluster
type ListDeletedRequest struct {
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

	// Find by techStatus
	// Required: false
	TechStatus string `url:"techStatus,omitempty" json:"techStatus,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListDeleted gets all deleted kubernetes clusters
func (k K8S) ListDeleted(ctx context.Context, req ListDeletedRequest) (*ListK8S, error) {

	url := "/cloudbroker/k8s/listDeleted"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListK8S{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
