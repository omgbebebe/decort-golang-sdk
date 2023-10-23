package sep

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get SEP parameters
type GetRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id" validate:"required"`
}

// Get gets SEP parameters as a RecordSEP struct
func (s SEP) Get(ctx context.Context, req GetRequest) (*RecordSEP, error) {
	res, err := s.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordSEP{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets SEP parameters as an array of bytes
func (s SEP) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/sep/get"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
