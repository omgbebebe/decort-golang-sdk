package account

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
)

// Request for setting CPU allocation ratio
type SetCPUAllocationRatioRequest struct {
	// Account ID
	// Required: true
	AccountID uint64 `url:"accountId" json:"accoutnId" validate:"required"`

	// CPU allocation ratio, i.e. one pCPU = ratio*vCPU
	// Required: true
	Ratio float64 `url:"ratio" json:"ratio" validate:"required"`
}

// SetCPUAllocationRatio sets CPU allocation ratio
func (a Account) SetCPUAllocationRatio(ctx context.Context, req SetCPUAllocationRatioRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/setCpuAllocationRatio"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
