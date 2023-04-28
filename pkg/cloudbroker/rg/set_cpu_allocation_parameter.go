package rg

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
)

// Request for setting CPU allocation parameter
type SetCPUAllocationParameterRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// CPU allocation parameter.
	// If "strict" VM can't be run if not enough CPU resources.
	// "loose" allow running VM if not enough resources.
	// Required: true
	StrictLoose string `url:"strict_loose" json:"strict_loose" validate:"required,strict_loose"`
}

// SetCPUAllocationParameter sets CPU allocation parameter
func (r RG) SetCPUAllocationParameter(ctx context.Context, req SetCPUAllocationParameterRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/rg/setCpuAllocationParameter"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
