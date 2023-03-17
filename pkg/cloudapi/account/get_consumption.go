package account

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for download the resources tracking files for an account
type GetConsumptionRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Epoch represents the start time
	// Required: true
	Start uint64 `url:"start" json:"start" validate:"required"`

	// Epoch represents the end time
	// Required: true
	End uint64 `url:"end" json:"end" validate:"required"`
}

// GetConsumption downloads the resources tracking files for an account within a given period
func (a Account) GetConsumption(ctx context.Context, req GetConsumptionRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/account/getConsumption"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil

}

// GetConsumptionGet downloads the resources tracking files for an account within a given period
func (a Account) GetConsumptionGet(ctx context.Context, req GetConsumptionRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/account/getConsumption"

	res, err := a.client.DecortApiCall(ctx, http.MethodGet, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
