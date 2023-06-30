package account

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete group accounts
type DeleteAccountsRequest struct {
	// IDs of accounts
	// Required: true
	AccountsIDs []uint64 `url:"accountIds" json:"accountIds" validate:"min=1"`

	// Reason for deletion
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`

	// Whether to completely destroy accounts or not
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`
}

// DeleteAccounts destroy a group of accounts
func (a Account) DeleteAccounts(ctx context.Context, req DeleteAccountsRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/deleteAccounts"

	_, err = a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
