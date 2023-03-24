package account

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request for enable group accounts
type EnableAccountsRequest struct {
	// IDs od accounts
	// Required: true
	AccountIDs []uint64 `url:"accountIds" json:"accountIds" validate:"min=1"`
}

// EnableAccounts enables accounts
func (a Account) EnableAccounts(ctx context.Context, req EnableAccountsRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/enableAccounts"

	_, err = a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
