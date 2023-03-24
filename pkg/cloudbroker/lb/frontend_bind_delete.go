package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete bind
type FrontendBindDeleteRequest struct {
	// ID of the load balancer instance to FrontendBindDelete
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`

	// Name of the frontend to delete
	// Required: true
	FrontendName string `url:"frontendName" json:"frontendName" validate:"required"`

	// Name of the binding to delete
	// Required: true
	BindingName string `url:"bindingName" json:"bindingName" validate:"required"`
}

// FrontendBindDelete deletes binding from the specified load balancer frontend
func (lb LB) FrontendBindDelete(ctx context.Context, req FrontendBindDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/lb/frontendBindDelete"

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
