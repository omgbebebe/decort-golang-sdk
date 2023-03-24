package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for reset config
type ConfigResetRequest struct {
	// ID of the load balancer instance to ConfigReset
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`
}

// ConfigReset reset current software configuration of the specified load balancer.
// Warning: this action cannot be undone!
func (lb LB) ConfigReset(ctx context.Context, req ConfigResetRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/lb/configReset"

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
