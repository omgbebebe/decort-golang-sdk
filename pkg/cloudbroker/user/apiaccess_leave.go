package user

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for leaving user from apiaccess group.
type APIAccessLeaveRequest struct {
	// ID of the user whose membership will be updated.
	// Required: true
	UserID string `url:"userId" json:"userId" validate:"required"`

	// ID of the API access group to leave.
	// Required: true
	APIAccessID uint64 `url:"apiaccessId" json:"apiaccessId" validate:"required"`
}

// APIAccessLeave leaves user from apiaccess group.
func (u User) APIAccessLeave(ctx context.Context, req APIAccessLeaveRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/user/apiaccessLeave"

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
