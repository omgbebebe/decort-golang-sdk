package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete backend
type BackendDeleteRequest struct {
	// ID of the load balancer instance to BackendDelete
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`

	// Cannot be emtpy string - name of the backend to delete
	// Required: true
	BackendName string `url:"backendName" json:"backendName" validate:"required"`
}

// BackendDelete deletes backend from the specified load balancer.
// Warning: you cannot undo this action!
func (lb LB) BackendDelete(ctx context.Context, req BackendDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/lb/backendDelete"

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
