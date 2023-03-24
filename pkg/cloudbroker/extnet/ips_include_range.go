package extnet

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for include range of IPs
type IPsIncludeRangeRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id" validate:"required"`

	// Starting IP
	// Required: true
	IPStart string `url:"ip_start" json:"ip_start" validate:"required"`

	// Ending IP
	// Required: true
	IPEnd string `url:"ip_end" json:"ip_end" validate:"required"`
}

// IPsIncludeRange include range of IPs to external network pool
func (e ExtNet) IPsIncludeRange(ctx context.Context, req IPsIncludeRangeRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/ipsIncludeRange"

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
