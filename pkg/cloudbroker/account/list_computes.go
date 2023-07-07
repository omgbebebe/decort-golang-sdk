package account

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for a get list compute instances
type ListComputesRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Find by compute id
	// Required: false
	ComputeID uint64 `url:"computeId,omitempty" json:"computeId,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by resource group name
	// Required: false
	RGName string `url:"rgName,omitempty" json:"rgName,omitempty"`

	// Find by resource group id
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Find by tech status
	// Required: false
	TechStatus string `url:"techStatus,omitempty" json:"techStatus,omitempty"`

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

// ListComputes gets list all compute instances under specified account, accessible by the user
func (a Account) ListComputes(ctx context.Context, req ListComputesRequest) (*ListComputes, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/listComputes"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
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
