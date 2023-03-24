package extnet

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for exclude range of IPs
type IPsExcludeRangeRequest struct {
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

// IPsExcludeRange exclude range of IPs from external network pool
func (e ExtNet) IPsExcludeRange(ctx context.Context, req IPsExcludeRangeRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/ipsExcludeRange"

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
