package k8ci

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	IncludeDisabled bool   `url:"includeDisabled,omitempty"`
	Page            uint64 `url:"page,omitempty"`
	Size            uint64 `url:"size,omitempty"`
}

func (k K8CI) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (K8CIList, error) {
	url := "/cloudapi/k8ci/list"
	k8ciListRaw, err := k.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	k8ciList := K8CIList{}
	if err := json.Unmarshal(k8ciListRaw, &k8ciList); err != nil {
		return nil, err
	}

	return k8ciList, nil
}
