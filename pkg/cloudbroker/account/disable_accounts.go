package account

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for disable group accounts
type DisableAccountsRequest struct {
	// IDs of accounts
	// Required: true
	AccountIDs []uint64 `url:"accountIds" json:"accountIds" validate:"min=1"`
}

// DisableAccounts disables accounts
func (a Account) DisableAccounts(ctx context.Context, req DisableAccountsRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/disableAccounts"

	_, err = a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
