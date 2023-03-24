package rg

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list computes from affinity group
type AffinityGroupsGetRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Label affinity group
	// Required: true
	AffinityGroup string `url:"affinityGroup" json:"affinityGroup" validate:"required"`
}

// AffinityGroupsGet gets list computes in the specified affinity group
func (r RG) AffinityGroupsGet(ctx context.Context, req AffinityGroupsGetRequest) ([]uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/rg/affinityGroupsGet"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := []uint64{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
