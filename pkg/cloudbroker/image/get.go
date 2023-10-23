package image

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get image details
type GetRequest struct {
	// ID of image
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`
}

// Get gets image details by ID as a RecordImage struct
func (i Image) Get(ctx context.Context, req GetRequest) (*RecordImage, error) {
	res, err := i.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordImage{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets image details by ID as an array of bytes
func (i Image) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/get"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
