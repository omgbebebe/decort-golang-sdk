package image

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete image
type DeleteRequest struct {
	// ID of the image to delete
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`

	// Reason for action
	// Required: true
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`

	// Whether to completely delete the image
	// Required: false
	Permanently bool `url:"permanently" json:"permanently"`
}

// Delete deletes image by ID
func (i Image) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/delete"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
