package rg

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list of computes
type ListComputesRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Find by compute id
	// Required: false
	ComputeID uint64 `url:"computeId,omitempty" json:"computeId,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// ID an account
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Find by tech status
	// Required: false
	TechStatus string `url:"techStatus,omitempty" json:"techStatus,omitempty"`

	// Find by status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Find by ip address
	// Required: false
	IPAddress string `url:"ipAddress,omitempty" json:"ipAddress,omitempty"`

	// Find by external network name
	// Required: false
	ExtNetName string `url:"extNetName,omitempty" json:"extNetName,omitempty"`

	// Find by external network id
	// Required: false
	ExtNetID uint64 `url:"extNetId,omitempty" json:"extNetId,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListComputes gets list of all compute instances under specified resource group, accessible by the user
func (r RG) ListComputes(ctx context.Context, req ListComputesRequest) (*ListComputes, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/rg/listComputes"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListComputes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
