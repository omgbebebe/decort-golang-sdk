package rg

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListResourceConsumption gets resource consumptions of the resource groups
func (r RG) ListResourceConsumption(ctx context.Context) (*ListResourceConsumption, error) {
	url := "/cloudbroker/rg/listResourceConsumption"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	list := ListResourceConsumption{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
