package disks

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list unattached disk
type ListUnattachedRequest struct {
	// ID of the account
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`
}

// ListUnattached gets list of unattached disks
func (d Disks) ListUnattached(ctx context.Context, req ListUnattachedRequest) (ListDisksUnattached, error) {
	url := "/cloudapi/disks/listUnattached"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListDisksUnattached{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
