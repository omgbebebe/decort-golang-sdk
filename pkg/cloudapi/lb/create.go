package lb

import (
	"context"
	"errors"
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
	// Required: false
	ExtNetID uint64 `url:"extnetId" json:"extnetId"`

	// Internal network (VINS) to connect this load balancer to
	// Required: false
	VINSID uint64 `url:"vinsId" json:"vinsId"`

	// Start now Load balancer
	// Required: false
	Start bool `url:"start" json:"start"`

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

	if req.ExtNetID == 0 && req.VINSID == 0 {
		return "", errors.New ("vinsId and extNetId cannot be both in the value 0")
	}

	url := "/cloudapi/lb/create"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
