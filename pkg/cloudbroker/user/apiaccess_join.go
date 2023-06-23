package user

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for joining user into apiaccess group.
type APIAccessJoinRequest struct {
	// ID of the user whose membership will be updated.
	// Required: true
	UserID string `url:"userId" json:"userId" validate:"required"`

	// ID of the API access group to join
	// Required: true
	APIAccessID uint64 `url:"apiaccessId" json:"apiaccessId" validate:"required"`
}

// APIAccessJoin joins user into apiaccess group.
func (u User) APIAccessJoin(ctx context.Context, req APIAccessJoinRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/user/apiaccessJoin"

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
