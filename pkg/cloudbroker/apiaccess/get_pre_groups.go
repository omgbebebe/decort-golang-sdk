package apiaccess

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetPreGroups gets list of pre default groups from spec
func (a APIAccess) GetPreGroups(ctx context.Context) (map[string]APIsEndpoints, error) {
	url := "/cloudbroker/apiaccess/getPreGroups"

	info := make(map[string]APIsEndpoints)

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return info, nil
}
