package lb

import (
	"context"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for frontend bind
type FrontendBindRequest struct {
	// ID of the load balancer instance to FrontendBind
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`

	// Name of the frontend to update
	// Required: true
	FrontendName string `url:"frontendName" json:"frontendName" validate:"required"`

	// Name of the binding to update
	// Required: true
	BindingName string `url:"bindingName" json:"bindingName"`

	// If specified must be within the IP range of either Ext Net or ViNS,
	// where this load balancer is connected - new IP address to use for this binding.
	// If omitted, current IP address is retained
	// Required: false
	BindingAddress string `url:"bindingAddress,omitempty" json:"bindingAddress,omitempty"`

	// New port number to use for this binding.
	// If omitted, current port number is retained
	// Required: false
	BindingPort uint64 `url:"bindingPort,omitempty" json:"bindingPort,omitempty"`
}

// FrontendBind bind frontend from specified load balancer instance
func (l LB) FrontendBind(ctx context.Context, req FrontendBindRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/lb/frontendBind"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
