package group

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// ListRequest struct to get list of group instances.
type ListRequest struct {
	// Find by id.
	// Requires: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by list users.
	// Required: false
	User string `url:"user,omitempty" json:"user,omitempty"`

	// Page number.
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size, maximum - 100.
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`

	// Find by active True or False.
	// Required: true
	Active bool `url:"active" json:"active" validate:"required"`
}

// List gets list of group instances as a ListGroups struct
func (g Group) List(ctx context.Context, req ListRequest) (*ListGroups, error) {
	res, err := g.ListRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := ListGroups{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// ListRaw gets list of group instances as an array of bytes
func (g Group) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/group/list"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
