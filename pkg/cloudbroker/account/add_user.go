package account

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for adding permission to access to account for a user
type AddUserRequest struct {
	// ID of account to add to
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Name of the user to be given rights
	// Required: true
	Username string `url:"username" json:"username" validate:"required"`

	// Account permission types:
	//	- 'R' for read only access
	//	- 'RCX' for Write
	//	- 'ARCXDU' for Admin
	// Required: true
	AccessType string `url:"accesstype" json:"accesstype" validate:"accessType"`
}

// AddUser gives a user access rights.
func (a Account) AddUser(ctx context.Context, req AddUserRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/addUser"

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
