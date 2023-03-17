package bservice

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get detailed information about Compute Group
type GroupGetRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
    ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// ID of the Compute Group
	// Required: true
	CompGroupID uint64 `url:"compgroupId" json:"compgroupId" validate:"required"`
}

// GroupGet gets detailed specifications for the Compute Group
func (b BService) GroupGet(ctx context.Context, req GroupGetRequest) (*RecordGroup, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/groupGet"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordGroup{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
