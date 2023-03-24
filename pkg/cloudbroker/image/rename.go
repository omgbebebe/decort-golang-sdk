package image

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for rename image
type RenameRequest struct {
	// ID of the virtual image to rename
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`

	// New name
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`
}

// Rename renames image by ID
func (i Image) Rename(ctx context.Context, req RenameRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/rename"

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
