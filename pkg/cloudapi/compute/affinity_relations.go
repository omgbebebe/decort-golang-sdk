package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get dict of computes
type AffinityRelationsRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// AffinityRelations gets dict of computes divided by affinity and anti affinity rules
func (c Compute) AffinityRelations(ctx context.Context, req AffinityRelationsRequest) (*RecordAffinityRelations, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/affinityRelations"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordAffinityRelations{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
