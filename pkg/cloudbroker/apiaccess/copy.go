package apiaccess

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for copying apiaccess group.
type CopyRequest struct {
	// ID of the API access group to make copy from
	// Required: true
	APIAccessID uint64 `url:"apiaccessId" json:"apiaccessId" validate:"required"`

	// Name of the target API access group, which will be created on successful copy
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`
}

// Copy creates a copy of the specified apiaccess group with a new name (and a new unique ID).
func (a APIAccess) Copy(ctx context.Context, req CopyRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/copy"

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
