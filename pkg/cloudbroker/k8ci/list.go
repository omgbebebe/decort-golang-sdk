package k8ci

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list information about images
type ListRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Find by worker driver
	// Required: false
	WorkerDriver string `url:"workerDriver,omitempty" json:"workerDriver,omitempty"`

	// Find by master driver
	// Required: false
	MasterDriver string `url:"masterDriver,omitempty" json:"masterDriver,omitempty"`

	// Find by network plugin
	// Required: false
	NetworkPlugins string `url:"netPlugins,omitempty" json:"masterDrnetPluginsiver,omitempty"`

	// List disabled items as well
	// Required: false
	IncludeDisabled bool `url:"includeDisabled,omitempty" json:"includeDisabled,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list all k8ci catalog items available to the current user
func (k K8CI) List(ctx context.Context, req ListRequest) (*ListK8CI, error) {
	url := "/cloudbroker/k8ci/list"

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
