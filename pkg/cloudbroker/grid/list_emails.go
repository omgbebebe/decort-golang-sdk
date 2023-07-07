package grid

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for getting list of email addresses of users
type ListEmailsRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListEmails returns list of email addresses of users
func (g Grid) ListEmails(ctx context.Context, req ListEmailsRequest) (*ListEmails, error) {

	url := "/cloudbroker/grid/listEmails"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListEmails{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
