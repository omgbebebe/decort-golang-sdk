package image

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for edit image
type EditRequest struct {
	// ID of the image to edit
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`

	// Name for the image
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Username for the image
	// Required: false
	Username string `url:"username,omitempty" json:"username,omitempty"`

	// Password for the image
	// Required: false
	Password string `url:"password,omitempty" json:"password,omitempty"`

	// Account ID to make the image exclusive
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Does this machine supports hot resize
	// Required: false
	HotResize bool `url:"hotresize,omitempty" json:"hotresize,omitempty"`

	// Does this image boot OS
	// Required: false
	Bootable bool `url:"bootable,omitempty" json:"bootable,omitempty"`
}

// Edit edits an existing image
func (i Image) Edit(ctx context.Context, req EditRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/edit"

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
