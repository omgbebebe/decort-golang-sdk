package account

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get information about account
type GetRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`
}

// Get gets information about account as a RecordAccount struct
func (a Account) Get(ctx context.Context, req GetRequest) (*RecordAccount, error) {
	res, err := a.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordAccount{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets information about account as an array of bytes
func (a Account) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/get"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
