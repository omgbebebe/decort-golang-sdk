package grid

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for set allocation
type SetCPUAllocationRatioRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gridId" json:"gridId" validate:"required"`

	// Allocation ratio
	// Required: true
	Ratio float64 `url:"ratio" json:"ratio" validate:"required"`
}

// SetCPUAllocationRatio sets CPU allocation ratio
func (g Grid) SetCPUAllocationRatio(ctx context.Context, req SetCPUAllocationRatioRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/setCpuAllocationRatio"

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
