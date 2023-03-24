package compute

import (
	"context"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get console URL
type GetConsoleURLRequest struct {
	// ID of compute instance to get console for
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// GetConsoleURL gets computes console URL
func (c Compute) GetConsoleURL(ctx context.Context, req GetConsoleURLRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/getConsoleUrl"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	result = strings.ReplaceAll(result, "\\", "")

	return result, nil
}
