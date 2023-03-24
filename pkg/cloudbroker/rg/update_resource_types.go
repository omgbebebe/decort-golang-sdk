package rg

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update resource types in account
type UpdateResourceTypesRequest struct {
	// ID of resource group
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Resource types available to create in this resource group
	// Each element in a resource type slice must be one of:
	//	- compute
	//	- vins
	//	- k8s
	//	- openshift
	//	- lb
	//	- flipgroup
	// Required: true
	ResTypes []string `url:"resourceTypes" json:"resourceTypes" validate:"min=1,resTypes"`
}

func (r RG) UpdateResourceTypes(ctx context.Context, req UpdateResourceTypesRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/rg/updateResourceTypes"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
