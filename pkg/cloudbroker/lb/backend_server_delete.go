package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete server definition
type BackendServerDeleteRequest struct {
	// ID of the load balancer instance to BackendServerDelete
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`

	// Must match one of the existing backens - name of the backend to add servers to
	// Required: true
	BackendName string `url:"backendName" json:"backendName" validate:"required"`

	// Must be unique among all servers defined for this backend - name of the server definition to add
	// Required: true
	ServerName string `url:"serverName" json:"serverName" validate:"required"`
}

// BackendServerDelete deletes server definition from the backend on the specified load balancer.
// Warning: you cannot undo this action!
func (lb LB) BackendServerDelete(ctx context.Context, req BackendServerDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/lb/backendServerDelete"

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
