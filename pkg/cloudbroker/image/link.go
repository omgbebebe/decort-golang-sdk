package image

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for link virtual image to another image
type LinkRequest struct {
	// ID of the virtual image
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`

	// ID of real image to link this virtual image to
	// Required: true
	TargetID uint64 `url:"targetId" json:"targetId" validate:"required"`
}

// Link links virtual image to another image in the platform
func (i Image) Link(ctx context.Context, req LinkRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/link"

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
