package grid

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
)

// Request for setting CPU allocation ratio for computes
type SetCPUAllocationRatioForVMRequest struct {
	// Grid ID
	// Required: true
	GridID uint64 `url:"gridId" json:"gridId" validate:"required"`

	// Default CPU allocation ratio for computes
	// Required: true
	Ratio float64 `url:"ratio" json:"ratio" validate:"required"`
}

// SetCPUAllocationRatio sets CPU allocation ratio for computes
func (g Grid) SetCPUAllocationRatioForVM(ctx context.Context, req SetCPUAllocationRatioForVMRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/setCpuAllocationRatioForVM"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
