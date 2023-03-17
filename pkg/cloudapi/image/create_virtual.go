package image

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create virtual image
type CreateVirtualRequest struct {
	// Name of the virtual image to create
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// ID of real image to link this virtual image to upon creation
	// Required: true
	TargetID uint64 `url:"targetId" json:"targetId" validate:"required"`
}

// CreateVirtual creates virtual image
func (i Image) CreateVirtual(ctx context.Context, req CreateVirtualRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/image/createVirtual"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
