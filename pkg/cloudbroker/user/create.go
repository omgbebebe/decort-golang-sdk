package user

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for creating a user.
type CreateRequest struct {
	// ID of user.
	// Required: true
	Username string `url:"username" json:"username" validate:"required"`

	// Email addresses of the user.
	// Required: true
	EmailAddress []string `url:"emailaddress" json:"emailaddress" validate:"required"`

	// Password of user
	// Required: false
	Password string `url:"password,omitempty" json:"password,omitempty"`

	// List of groups this user belongs to.
	// Required: false
	Groups []string `url:"groups,omitempty" json:"groups,omitempty"`

	// List of apiaccess groups this user belongs to.
	// Required: false
	APIAccess []uint64 `url:"apiaccess,omitempty" json:"apiaccess,omitempty"`
}

// Create creates a user.
func (u User) Create(ctx context.Context, req CreateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/user/create"

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
