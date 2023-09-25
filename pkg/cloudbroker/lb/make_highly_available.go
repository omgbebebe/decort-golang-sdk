package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for make Load Balancer Highly available
type HighlyAvailableRequest struct {
	// ID of the LB instance
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`
}

// Make Load Balancer Highly available
func (l LB) HighlyAvailable(ctx context.Context, req HighlyAvailableRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/lb/makeHighlyAvailable"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
