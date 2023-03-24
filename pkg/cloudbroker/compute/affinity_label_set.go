package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for set affinity label for compute
type AffinityLabelSetRequest struct {
	// IDs of the compute instances
	ComputeIDs []uint64 `url:"computeIds" json:"computeIds" validate:"min=1"`

	// Affinity group label
	// Required: true
	AffinityLabel string `url:"affinityLabel" json:"affinityLabel" validate:"required"`
}

// AffinityLabelSet set affinity label for compute
func (c Compute) AffinityLabelSet(ctx context.Context, req AffinityLabelSetRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/affinityLabelSet"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
