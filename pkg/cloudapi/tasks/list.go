package tasks

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct to get list of tasks
type ListRequest struct {
	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of user API tasks with status PROCESSING as a ListTasks struct
func (t Tasks) List(ctx context.Context, req ListRequest) (*ListTasks, error) {
	res, err := t.ListRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	list := ListTasks{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// ListRaw gets list of user API tasks with status PROCESSING as an array of bytes
func (t Tasks) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudapi/tasks/list"

	res, err := t.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
