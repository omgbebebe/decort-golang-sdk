package rg

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list of affinity groups from resource group
type AffinityGroupsListRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// AffinityGroupsList gets all currently defined affinity groups in this resource group with compute IDs
func (r RG) AffinityGroupsList(ctx context.Context, req AffinityGroupsListRequest) (*ListAffinityGroups, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/rg/affinityGroupsList"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := &ListAffinityGroups{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
