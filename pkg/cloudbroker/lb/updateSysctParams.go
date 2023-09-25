package lb

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update sysct params for lb
type UpdateSysctParamsRequest struct {
	// ID of the LB instance
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`

	// Custom sysctl values for Load Balancer instance. Applied on boot
	// Required: true
	SysctlParams Params `url:"-" json:"sysctlParams" validate:"required,dive"`
}

type wrapperUpdateSysctParamsRequest struct {
	UpdateSysctParamsRequest
	Params []string `url:"sysctlParams" validate:"required"`
}

// Create method will create a new load balancer instance
func (l LB) UpdateSysctParams(ctx context.Context, req UpdateSysctParamsRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	var params []string

	if len(req.SysctlParams) != 0 {
		params = make([]string, 0, len(req.SysctlParams))

		for r := range req.SysctlParams {
			b, err := json.Marshal(req.SysctlParams[r])
			if err != nil {
				return false, err
			}

			params = append(params, string(b))
		}
	} else {
		params = []string{}
	}

	reqWrapped := wrapperUpdateSysctParamsRequest{
		UpdateSysctParamsRequest: req,
		Params:                   params,
	}

	url := "/cloudbroker/lb/updateSysctParams"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
