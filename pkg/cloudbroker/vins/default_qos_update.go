package vins

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update QOS
type DefaultQOSUpdateRequest struct {
	// ID of VINS
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// Internal traffic, kbit
	// Required: false
	IngressRate uint64 `url:"ingress_rate,omitempty" json:"ingress_rate,omitempty"`

	// Internal traffic burst, kbit
	// Required: false
	IngressBirst uint64 `url:"ingress_birst,omitempty" json:"ingress_birst,omitempty"`

	// External traffic rate, kbit
	// Required: false
	EgressRate uint64 `url:"egress_rate,omitempty" json:"egress_rate,omitempty"`
}

// DefaultQOSUpdate update default QOS values
func (v VINS) DefaultQOSUpdate(ctx context.Context, req DefaultQOSUpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/defaultQosUpdate"

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
