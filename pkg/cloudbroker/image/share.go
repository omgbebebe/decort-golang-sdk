package image

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for share image
type ShareRequest struct {
	// ID of the image to share
	// Required: true
	ImageId uint64 `url:"imageId" json:"imageId" validate:"required"`

	// List of account IDs
	// Required: true
	AccountIDs []uint64 `url:"accounts" json:"accounts" validate:"min=1"`
}

// Share shares image with accounts
func (i Image) Share(ctx context.Context, req ShareRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/share"

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
