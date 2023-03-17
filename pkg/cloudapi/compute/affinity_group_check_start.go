package compute

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for check all computes with current affinity label can start
type AffinityGroupCheckStartRequest struct {
	// ID of the resource group
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Affinity group label
	// Required: true
	AffinityLabel string `url:"affinityLabel" json:"affinityLabel" validate:"required"`
}

// AffinityGroupCheckStart check all computes with current affinity label can start
func (c Compute) AffinityGroupCheckStart(ctx context.Context, req AffinityGroupCheckStartRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/affinityGroupCheckStart"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
