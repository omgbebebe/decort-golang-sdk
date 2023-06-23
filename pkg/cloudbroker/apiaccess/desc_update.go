package apiaccess

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for updating apiaccess group description.
type DescUpdateRequest struct {
	// APIAccess group ID.
	// Required: true
	APIAccessID uint64 `url:"apiaccessId" json:"apiaccessId" validate:"required"`

	// New description to set for the apiaccess group.
	// Required: true
	Description string `url:"desc" json:"desc" validate:"required"`
}

// DescUpdate sets a new text description of the apiaccess group.
func (a APIAccess) DescUpdate(ctx context.Context, req DescUpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/descUpdate"

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
