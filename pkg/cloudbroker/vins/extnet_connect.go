package vins

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for connect external network
type ExtNetConnectRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// External network ID
	// Required: false
	NetID uint64 `url:"netId,omitempty" json:"netId,omitempty"`

	// Directly set IP address
	// Required: false
	IP string `url:"ip,omitempty" json:"ip,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// ExtNetConnect connect VINS to external network
func (v VINS) ExtNetConnect(ctx context.Context, req ExtNetConnectRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/extNetConnect"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
