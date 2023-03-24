package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update load balancer
type UpdateRequest struct {
	// ID of the load balancer to update
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`

	// New description of this load balancer.
	// If omitted, current description is retained
	// Required: true
	Description string `url:"desc" json:"desc" validate:"required"`
}

// Update updates some of load balancer attributes
func (l LB) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/lb/update"

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
