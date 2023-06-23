package apiaccess

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for creating apiaccess group.
type CreateRequest struct {
	// Name of this apiaccess group.
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Description of this apiaccess group.
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

// Create creates apiaccess group.
func (a APIAccess) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/create"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
