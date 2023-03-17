package account

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list VINS
type ListVINSRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`
}

// ListVINS gets list all ViNSes under specified account, accessible by the user
func (a Account) ListVINS(ctx context.Context, req ListVINSRequest) (ListVINS, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/account/listVins"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListVINS{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
