package compute

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for restore compute
type RestoreRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// Restore restore compute from recycle bin
func (c Compute) Restore(ctx context.Context, req RestoreRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/restore"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
