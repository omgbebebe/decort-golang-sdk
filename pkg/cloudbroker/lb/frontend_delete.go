package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete frontend
type FrontendDeleteRequest struct {
	// ID of the load balancer instance to FrontendDelete
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`

	// Name of the frontend to delete
	// Required: true
	FrontendName string `url:"frontendName" json:"frontendName" validate:"required"`
}

// FrontendDelete deletes frontend from the specified load balancer.
// Warning: you cannot undo this action!
func (lb LB) FrontendDelete(ctx context.Context, req FrontendDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/lb/frontendDelete"

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
