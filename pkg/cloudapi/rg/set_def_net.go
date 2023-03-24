package rg

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for set default network
type SetDefNetRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Network type
	// Should be one of:
	//	- "PUBLIC"
	//	- "PRIVATE"
	// Required: true
	NetType string `url:"netType" json:"netType" validate:"rgNetType"`

	// Network ID
	// Required: false
	NetID uint64 `url:"netId,omitempty" json:"netId,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// SetDefNet sets default network for attach associated virtual machines
func (r RG) SetDefNet(ctx context.Context, req SetDefNetRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/rg/setDefNet"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
