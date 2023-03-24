package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update binding
type FrontendBindUpdateRequest struct {
	// ID of the load balancer instance to FrontendBindUpdate
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`

	// Name of the frontend to update
	// Required: true
	FrontendName string `url:"frontendName" json:"frontendName" validate:"required"`

	// Name of the binding to update
	// Required: true
	BindingName string `url:"bindingName" json:"bindingName" validate:"required"`

	// If specified must be within the IP range of either Ext Net or ViNS,
	// where this load balancer is connected - new IP address to use for this binding.
	// If omitted, current IP address is retained
	// Required: true
	BindingAddress string `url:"bindingAddress" json:"bindingAddress" validate:"required"`

	// New port number to use for this binding.
	// If omitted, current port number is retained
	// Required: true
	BindingPort uint64 `url:"bindingPort" json:"bindingPort" validate:"required"`
}

// FrontendBindUpdate updates binding for the specified load balancer frontend
func (lb LB) FrontendBindUpdate(ctx context.Context, req FrontendBindUpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/lb/frontendBindingUpdate"

	res, err := lb.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
