package apiaccess

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Reqeust struct for setting default apiaccess group.
type SetDefaultRequest struct {
	// APIAccess group ID
	// Required: true
	APIAccessID uint64 `url:"apiaccessId" json:"apiaccessId" validate:"required"`
}

// SetDefault sets current apiaccess group default.
func (a APIAccess) SetDefault(ctx context.Context, req SetDefaultRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/setDefault"

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
