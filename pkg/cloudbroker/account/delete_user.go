package account

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for revoke access to account
type DeleteUserRequest struct {
	// ID of the account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// ID or emailaddress of the user to remove
	// Required: true
	UserName string `url:"username" json:"username" validate:"required"`

	// Recursively revoke access rights from owned cloudspaces and vmachines
	// Required: false
	RecursiveDelete bool `url:"recursivedelete" json:"recursivedelete" validate:"required"`
}

// DeleteUser revokes user access from the account
func (a Account) DeleteUser(ctx context.Context, req DeleteUserRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/deleteUser"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
