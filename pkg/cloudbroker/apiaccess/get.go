package apiaccess

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get apiaccess group.
type GetRequest struct {
	// APIAccess group ID.
	// Required: true
	APIAccessID uint64 `url:"apiaccessId" json:"apiaccessId" validate:"required"`
}

// Get gets apiaccess group as an ItemAPIAccess struct
func (a APIAccess) Get(ctx context.Context, req GetRequest) (*ItemAPIAccess, error) {
	res, err := a.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := ItemAPIAccess{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets apiaccess group as an array of bytes
func (a APIAccess) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/get"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
