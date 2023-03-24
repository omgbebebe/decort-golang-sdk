package rg

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for grant access to resource group
type AccessGrantRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// User or group name to grant access
	// Required: true
	User string `url:"user" json:"user" validate:"required"`

	// Access rights to set, one of:
	//	- "R"
	//	-  "RCX"
	//	- "ARCXDU"
	// Required: true
	Right string `url:"right" json:"right" validate:"accessType"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// AccessGrant grants user or group access to the resource group as specified
func (r RG) AccessGrant(ctx context.Context, req AccessGrantRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/rg/accessGrant"

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
