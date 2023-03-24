package image

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update nodes
type UpdateNodesRequest struct {
	// Image ID
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`

	// List of stacks
	// Required: false
	EnabledStacks []uint64 `url:"enabledStacks,omitempty" json:"enabledStacks,omitempty"`
}

// UpdateNodes udates image availability on nodes
func (i Image) UpdateNodes(ctx context.Context, req UpdateNodesRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/updateNodes"

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
