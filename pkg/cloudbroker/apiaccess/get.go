package apiaccess

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for getting apiaccess group.
type GetRequest struct {
	// APIAccess group ID.
	// Required: true
	APIAccessID uint64 `url:"apiaccessId" json:"apiaccessId" validate:"required"`
}

// Get gets apiaccess group.
func (a APIAccess) Get(ctx context.Context, req GetRequest) (*ItemAPIAccess, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/get"

	info := ItemAPIAccess{}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
