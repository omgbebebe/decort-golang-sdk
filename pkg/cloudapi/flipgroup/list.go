package flipgroup

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct to get list of FLIPGroup available to the current user
type ListRequest struct {
	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by vinsId
	// Required: false
	VINSID uint64 `url:"vinsId,omitempty" json:"vinsId,omitempty"`

	// Find by VINS name
	// Required: false
	VINSName string `url:"vinsName,omitempty" json:"vinsName,omitempty"`

	// Find by extnetId
	// Required: false
	ExtNetID uint64 `url:"extnetId,omitempty" json:"extnetId,omitempty"`

	// Find by IP
	// Required: false
	ByIP string `url:"byIp,omitempty" json:"byIp,omitempty"`

	// Find by resource group ID
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Find by id
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of FLIPGroup managed cluster instances available to the current user as a ListFLIPGroups struct
func (f FLIPGroup) List(ctx context.Context, req ListRequest) (*ListFLIPGroups, error) {
	res, err := f.ListRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	list := ListFLIPGroups{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// ListRaw gets list of FLIPGroup managed cluster instances available to the current user as an array of bytes
func (f FLIPGroup) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudapi/flipgroup/list"

	res, err := f.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
