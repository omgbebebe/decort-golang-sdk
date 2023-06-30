package account

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListResourceConsumption show data list amount of consumed and reserved resources (cpu, ram, disk) by specific accounts
func (a Account) ListResourceConsumption(ctx context.Context) (*ListResources, error) {
	url := "/cloudbroker/account/listResourceConsumption"

	info := ListResources{}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
