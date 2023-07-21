package rg

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list of all computes with their relationships
type AffinityGroupComputesRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Affinity group label
	// Required: true
	AffinityGroup string `url:"affinityGroup" json:"affinityGroup" validate:"required"`
}

// AffinityGroupComputes gets list of all computes with their relationships to another computes
func (r RG) AffinityGroupComputes(ctx context.Context, req AffinityGroupComputesRequest) (ListAffinityGroupsComputes, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/rg/affinityGroupComputes"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListAffinityGroupsComputes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
