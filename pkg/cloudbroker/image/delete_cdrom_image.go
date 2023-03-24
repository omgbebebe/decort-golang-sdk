package image

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete CD-ROM image
type DeleteCDROMImageRequest struct {
	// ID of the CD-ROM image to delete
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`

	// Whether to completely delete the CD-ROM image, needs to be unused
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`
}

// DeleteCDROMImage delete a CD-ROM image
func (i Image) DeleteCDROMImage(ctx context.Context, req DeleteCDROMImageRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/deleteCDROMImage"

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
