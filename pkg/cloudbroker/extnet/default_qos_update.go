package extnet

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update QOS
type DefaultQOSUpdateRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id" validate:"required"`

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

// DefaultQOSUpdate updates default qos values
func (e ExtNet) DefaultQOSUpdate(ctx context.Context, req DefaultQOSUpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/defaultQosUpdate"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
