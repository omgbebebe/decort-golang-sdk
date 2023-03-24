package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for remove affinity rule
type AffinityRuleRemoveRequest struct {
	// IDs of the compute instances
	// Required: true
	ComputeIDs []uint64 `url:"computeIds" json:"computeIds" validate:"min=1"`

	// Compute or node, for whom rule applies
	// Required: true
	Topology string `url:"topology" json:"topology" validate:"computeTopology"`

	// The degree of 'strictness' of this rule
	// Should be one of:
	//	- RECOMMENDED
	//	- REQUIRED
	// Required: true
	Policy string `url:"policy" json:"policy" validate:"computePolicy"`

	// The comparison mode is 'value', recorded by the specified 'key'
	// Should be one of:
	//	- EQ
	//	- EN
	//	- ANY
	// Required: true
	Mode string `url:"mode" json:"mode" validate:"computeMode"`

	// Key that are taken into account when analyzing this rule will be identified
	// Required: true
	Key string `url:"key" json:"key" validate:"required"`

	// Value that must match the key to be taken into account when analyzing this rule
	// Required: true
	Value string `url:"value" json:"value" validate:"required"`
}

// AffinityRuleRemove remove affinity rule
func (c Compute) AffinityRuleRemove(ctx context.Context, req AffinityRuleRemoveRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/affinityRuleRemove"

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
