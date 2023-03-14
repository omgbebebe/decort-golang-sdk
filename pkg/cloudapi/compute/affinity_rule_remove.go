package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"repos.digitalenergy.online/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for remove affinity rule
type AffinityRuleRemoveRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// Compute or node, for whom rule applies
	// Required: true
	Topology string `url:"topology" json:"topology"`

	// The degree of 'strictness' of this rule
	// Should be one of:
	//	- RECOMMENDED
	//	- REQUIRED
	// Required: true
	Policy string `url:"policy" json:"policy"`

	// The comparison mode is 'value', recorded by the specified 'key'
	// Should be one of:
	//	- EQ
	//	- EN
	//	- ANY
	// Required: true
	Mode string `url:"mode" json:"mode"`

	// Key that are taken into account when analyzing this rule will be identified
	// Required: true
	Key string `url:"key" json:"key"`

	// Value that must match the key to be taken into account when analyzing this rule
	// Required: true
	Value string `url:"value" json:"value"`
}

func (crq AffinityRuleRemoveRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.Topology == "" {
		return errors.New("validation-error: field Topology can not be empty")
	}
	validator := validators.StringInSlice(crq.Topology, []string{"compute", "node"})
	if !validator {
		return errors.New("validation-error: field Topology can be only compute or node")
	}
	if crq.Policy == "" {
		return errors.New("validation-error: field Policy can not be empty")
	}
	validator = validators.StringInSlice(crq.Policy, []string{"RECOMMENDED", "REQUIRED"})
	if !validator {
		return errors.New("validation-error: field Policy can be only RECOMMENDED or REQUIRED")
	}
	if crq.Mode == "" {
		return errors.New("validation-error: field Mode can not be empty")
	}
	validator = validators.StringInSlice(crq.Mode, []string{"EQ", "NE", "ANY"})
	if !validator {
		return errors.New("validation-error: field Mode can be only EQ, NE or ANY")
	}
	if crq.Key == "" {
		return errors.New("validation-error: field Key can not be empty")
	}
	if crq.Value == "" {
		return errors.New("validation-error: field Value can not be empty")
	}

	return nil
}

// AffinityRuleRemove remove affinity rule
func (c Compute) AffinityRuleRemove(ctx context.Context, req AffinityRuleRemoveRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/affinityRuleRemove"

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
