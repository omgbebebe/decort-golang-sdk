package image

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get detailed information about image
type GetRequest struct {
	// ID of image to get
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`

	// If set to False returns only images in status CREATED
	// Required: false
	ShowAll bool `url:"show_all,omitempty" json:"show_all,omitempty"`
}

// Get gets image by ID.
// Returns image as a RecordImage struct if user has rights on it
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

// GetRaw gets image by ID.
// Returns image as an array of bytes if user has rights on it
func (i Image) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/image/get"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
