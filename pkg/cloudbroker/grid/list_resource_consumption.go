package grid

import (
	"context"
	"encoding/json"
	"net/http"
)

func (g Grid) ListResourceConsumption(ctx context.Context) (*ListResourceConsumption, error) {
	url := "/cloudbroker/grid/listResourceConsumption"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	result := ListResourceConsumption{}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
