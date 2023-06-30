package account

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list deleted accounts
type ListDeletedRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page" json:"page"`

	// Page size
	// Required: false
	Size uint64 `url:"size" json:"size"`
}

// ListDeleted gets list all deleted accounts the user has access to
func (a Account) ListDeleted(ctx context.Context, req ListDeletedRequest) (ListDeleted, error) {
	url := "/cloudbroker/account/listDeleted"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListDeleted{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
