package extnet

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for include list IPs
type IPsIncludeRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id" validate:"required"`

	// List of IPs for include to external network
	// Required: true
	IPs []string `url:"ips" json:"ips" validate:"min=1"`
}

// IPsInclude include list IPs to external network pool
func (e ExtNet) IPsInclude(ctx context.Context, req IPsIncludeRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/ipsInclude"

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
