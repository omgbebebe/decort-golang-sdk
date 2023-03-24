package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for restart load balancer
type RestartRequest struct {
	// ID of the load balancer instance to restart
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`
}

// Restart restarts specified load balancer instance
func (l LB) Restart(ctx context.Context, req RestartRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/lb/restart"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
