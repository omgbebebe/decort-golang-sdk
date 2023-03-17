package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for clear anti affinity rules
type AntiAffinityRulesClearRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// AntiAffinityRulesClear clear anti affinity rules
func (c Compute) AntiAffinityRulesClear(ctx context.Context, req AntiAffinityRulesClearRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/antiAffinityRulesClear"

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
