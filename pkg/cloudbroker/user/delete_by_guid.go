package user

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for deleting a user using user's GUID.
type DeleteByGUIDRequest struct {
	// GUID of user.
	// Required: true
	GUID string `url:"userguid" json:"userguid" validate:"required"`
}

// DeleteByGUID deletes a user using user's GUID.
// Note: This actor can also be called using username instead of guid to workaround CBGrid
// allowing userguid or username.
func (u User) DeleteByGUID(ctx context.Context, req DeleteByGUIDRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/user/deleteByGuid"

	res, err := u.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
