package vins

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create NAT rules
type NATRuleAddRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// Internal IP address to apply this rule to
	// Required: true
	IntIP string `url:"intIp" json:"intIp" validate:"required"`

	// Internal IP port number to use for this rule
	// Required: true
	IntPort uint `url:"intPort" json:"intPort" validate:"required"`

	// External IP start port to use for this rule
	// Required: true
	ExtPortStart uint `url:"extPortStart" json:"extPortStart" validate:"required"`

	// External IP end port to use for this rule
	// Required: false
	ExtPortEnd uint `url:"extPortEnd,omitempty" json:"extPortEnd,omitempty"`

	// IP protocol type
	// Should be one of:
	//	- "tcp"
	//	- "udp"
	// Required: false
	Proto string `url:"proto,omitempty" json:"proto,omitempty" validate:"omitempty,proto"`
}

// NATRuleAdd create NAT (port forwarding) rule on VINS
func (v VINS) NATRuleAdd(ctx context.Context, req NATRuleAddRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/vins/natRuleAdd"

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
