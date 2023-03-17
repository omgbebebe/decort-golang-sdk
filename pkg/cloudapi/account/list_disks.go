package account

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list deleted disks
type ListDisksRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`
}

// ListDisks gets list all currently unattached disks under specified account
func (a Account) ListDisks(ctx context.Context, req ListDisksRequest) (ListDisks, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/account/listDisks"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListDisks{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
