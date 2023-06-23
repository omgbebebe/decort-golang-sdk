package apiaccess

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetFull gets full current endpoints dictionary
func (a APIAccess) GetFull(ctx context.Context) (*APIsEndpoints, error) {
	url := "/cloudbroker/apiaccess/getFull"

	info := APIsEndpoints{}

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
