package compute

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get compute logs
type GetLogRequest struct {
	// ID of compute instance to get log for
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Path to log file
	// Required: true
	Path string `url:"path" json:"path" validate:"required"`
}

// GetLog gets compute's log file by path
func (c Compute) GetLog(ctx context.Context, req GetLogRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/getLog"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
