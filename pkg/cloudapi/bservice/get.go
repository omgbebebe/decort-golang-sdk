package bservice

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get detailed information about service
type GetRequest struct {
	// ID of the service to query information
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`
}

// Get gets detailed specifications for the BasicService.
func (b BService) Get(ctx context.Context, req GetRequest) (*RecordBasicService, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/get"

	bsRaw, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordBasicService{}

	err = json.Unmarshal(bsRaw, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
