package vins

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get information about VINS
type GetRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`
}

// Get gets information about VINS by ID as a RecordVINS struct
func (v VINS) Get(ctx context.Context, req GetRequest) (*RecordVINS, error) {
	res, err := v.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordVINS{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets information about VINS by ID as an array of bytes
func (v VINS) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/vins/get"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
