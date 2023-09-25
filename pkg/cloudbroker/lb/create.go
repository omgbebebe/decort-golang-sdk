package lb

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

type Params []string

// Request struct for create load balancer
type CreateRequest struct {
	// ID of the resource group where this load balancer instance will be located
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Name of the load balancer.
	// Must be unique among all load balancers in this resource group
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// External network to connect this load balancer to
	// Required: false
	ExtNetID uint64 `url:"extnetId" json:"extnetId"`

	// Internal network (VINS) to connect this load balancer to
	// Required: false
	VINSID uint64 `url:"vinsId" json:"vinsId"`

	// Custom  sysctl values for Load Balancer instance. Applied on boot
	// Required: false
	SysctlParams Params `url:"-" json:"sysctlParams,omitempty" validate:"omitempty,dive"`

	// Use Highly Available schema for LB deploy
	// Required: false
	HighlyAvailable bool `url:"highlyAvailable,omitempty" json:"highlyAvailable,omitempty"`

	// Start now Load balancer
	// Required: false
	Start bool `url:"start" json:"start"`

	// Text description of this load balancer
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

type wrapperCreateRequest struct {
	CreateRequest
	Params []string `url:"sysctlParams,omitempty"`
}

// Create method will create a new load balancer instance
func (lb LB) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	if req.ExtNetID == 0 && req.VINSID == 0 {
		return 0, errors.New("vinsId and extNetId cannot be both in the value 0")
	}

	var params []string

	if len(req.SysctlParams) != 0 {
		params = make([]string, 0, len(req.SysctlParams))

		for r := range req.SysctlParams {
			b, err := json.Marshal(req.SysctlParams[r])
			if err != nil {
				return 0, err
			}

			params = append(params, string(b))
		}
	} else {
		params = []string{}
	}

	reqWrapped := wrapperCreateRequest{
		CreateRequest: req,
		Params:        params,
	}

	url := "/cloudbroker/lb/create"

	res, err := lb.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
