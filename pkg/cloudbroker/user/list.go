package user

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// ListRequest struct to get all non deleted user instances.
type ListRequest struct {
	// Find by ID.
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by active. True or False.
	// Required: false
	Active bool `url:"active,omitempty" json:"active,omitempty"`

	// Find by serviceaccount. True or False.
	// Required: false
	ServiceAccount bool `url:"serviceaccount,omitempty" json:"serviceaccount,omitempty"`

	// Page number.
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size, maximum - 100.
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets all non deleted user instances as a ListUsers struct
func (u User) List(ctx context.Context, req ListRequest) (*ListUsers, error) {
	res, err := u.ListRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	list := ListUsers{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// ListRaw gets all non deleted user instances as an array of bytes
func (u User) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/user/list"

	res, err := u.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
