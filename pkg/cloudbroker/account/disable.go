package account

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for disable account
type DisableRequest struct {
	// ID of account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Reason to disable
	// Required: true
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Disable disables an account
func (a Account) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/disable"

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
