package account

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for give list account audits
type AuditsRequest struct {
	// ID of the account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`
}

// Audits gets audit records for the specified account object
func (a Account) Audits(ctx context.Context, req AuditsRequest) (ListAudits, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/account/audits"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListAudits{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
