package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update QOS
type NetQOSRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Network ID
	// Required: true
	NetID uint64 `url:"netId" json:"netId" validate:"required"`

	// Network type
	// Should be one of:
	//	- VINS
	//	- EXTNET
	// Required: true
	NetType string `url:"netType" json:"netType" validate:"computeNetType"`

	// Internal traffic, kbit
	// Required: false
	IngressRate uint64 `url:"ingress_rate,omitempty" json:"ingress_rate,omitempty"`

	// Internal traffic burst, kbit
	// Required: false
	IngressBurst uint64 `url:"ingress_burst,omitempty" json:"ingress_burst,omitempty"`

	// External traffic rate, kbit
	// Required: false
	EgressRate uint64 `url:"egress_rate,omitempty" json:"egress_rate,omitempty"`
}

// NetQOS update compute interfaces QOS
func (c Compute) NetQOS(ctx context.Context, req NetQOSRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/netQos"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
