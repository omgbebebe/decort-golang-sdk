package image

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get image details
type GetRequest struct {
	// ID of image
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`
}

// Get get image details by ID
func (i Image) Get(ctx context.Context, req GetRequest) (*RecordImage, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/get"

	info := RecordImage{}

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
