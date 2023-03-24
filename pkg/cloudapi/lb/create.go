package lb

import (
	"context"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create load balancer
type CreateRequest struct {
	// ID of the resource group where this load balancer instance will be located
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Name of the load balancer.
	// Must be unique among all load balancers in this Resource Group
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// External network to connect this load balancer to
	// Required: true
	ExtNetID uint64 `url:"extnetId" json:"extnetId" validate:"required"`

	// Internal network (VINS) to connect this load balancer to
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// Start now Load balancer
	// Required: true
	Start bool `url:"start" json:"start" validate:"required"`

	// Text description of this load balancer
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

// Create method will create a new load balancer instance
func (l LB) Create(ctx context.Context, req CreateRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/lb/create"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
