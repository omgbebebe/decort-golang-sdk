package account

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete account
type DeleteRequest struct {
	// ID of account to delete
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Reason to delete
	// Required: true
	Reason string `url:"reason" json:"reason" validate:"required"`

	// Whether to completely delete the account
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`
}

// Delete completes delete an account from the system Returns true if account is deleted or was already deleted or never existed
func (a Account) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/delete"

	_, err = a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
