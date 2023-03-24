package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create frontend
type FrontendCreateRequest struct {
	// ID of the load balancer instance to FrontendCreate
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`

	// Must be unique among all frontends of
	// this load balancer - name of the new frontend to create
	// Required: true
	FrontendName string `url:"frontendName" json:"frontendName" validate:"required"`

	// Should be one of the backends existing on
	// this load balancer - name of the backend to use
	// Required: true
	BackendName string `url:"backendName" json:"backendName" validate:"required"`
}

// FrontendCreate creates new frontend on the specified load balancer
func (l LB) FrontendCreate(ctx context.Context, req FrontendCreateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/lb/frontendCreate"

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
