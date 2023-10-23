package user

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get user details.
type GetRequest struct {
	// ID of the user.
	// Required: true
	UserID string `url:"userId" json:"userId" validate:"required"`
}

// Get gets user details as an ItemUser struct.
func (u User) Get(ctx context.Context, req GetRequest) (*ItemUser, error) {
	res, err := u.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	item := ItemUser{}

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// GetRaw gets user details as an array of bytes
func (u User) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/user/get"

	res, err := u.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
