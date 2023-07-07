package k8ci

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list information about deleted images
type ListDeletedRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"k8cId,omitempty" json:"k8cId,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by worker driver
	// Required: false
	WorkerDriver string `url:"workerDriver,omitempty" json:"workerDriver,omitempty"`

	// Find by master driver
	// Required: false
	MasterDriver string `url:"masterDriver,omitempty" json:"masterDriver,omitempty"`

	// Find by network plugin
	// Required: false
	NetworkPlugins string `url:"netPlugins,omitempty" json:"netPlugins,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListDeleted gets list all deleted k8ci catalog items available to the current user
func (k K8CI) ListDeleted(ctx context.Context, req ListDeletedRequest) (*ListK8CI, error) {
	url := "/cloudapi/k8ci/listDeleted"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListK8CI{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
