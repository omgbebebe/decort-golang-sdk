package extnet

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for set external network as default
type SetDefaultRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id" validate:"required"`
}

// SetDefault sets external network as default for platform
func (e ExtNet) SetDefault(ctx context.Context, req SetDefaultRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/setDefault"

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
