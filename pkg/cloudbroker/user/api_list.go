package user

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for getting API list.
type APIListRequest struct {
	// ID of the user.
	// Required: true
	UserID string `url:"userId" json:"userId" validate:"required"`
}

// APIList gets a list of all API functions that a given user has
// access to according to their apiaccess group membership.
func (u User) APIList(ctx context.Context, req APIListRequest) (*APIsEndpoints, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/user/apiList"

	info := APIsEndpoints{}

	res, err := u.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
