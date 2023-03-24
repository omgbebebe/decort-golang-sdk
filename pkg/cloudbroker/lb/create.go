package lb

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create load balancer
type CreateRequest struct {
	// ID of the resource group where this load balancer instance will be located
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Name of the load balancer.
	// Must be unique among all load balancers in this resource group
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// OS image ID to create load balancer from
	// Required: false
	ImageID uint64 `url:"imageId,omitempty" json:"imageId,omitempty"`

	// External network to connect this load balancer to
	// Required: true
	ExtNetID uint64 `url:"extnetId" json:"extnetId" validate:"required"`

	// Internal network (VINS) to connect this load balancer to
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// Start now Load balancer
	// Required: false
	Start bool `url:"start" json:"start" validate:"required"`

	// Text description of this load balancer
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

// Create method will create a new load balancer instance
func (lb LB) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/lb/create"

	res, err := lb.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
