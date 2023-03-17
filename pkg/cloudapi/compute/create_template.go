package compute

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create template
type CreateTemplateRequest struct {
	// ID of the compute to create template from
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Name to assign to the template being created
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Async API call
	// For async call use CreateTemplateAsync
	// For sync call use CreateTemplate
	// Required: true
	async bool `url:"async"`
}

// CreateTemplate create template from compute instance
func (c Compute) CreateTemplate(ctx context.Context, req CreateTemplateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	req.async = false

	url := "/cloudapi/compute/createTemplate"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// CreateTemplateAsync create template from compute instance
func (c Compute) CreateTemplateAsync(ctx context.Context, req CreateTemplateRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	req.async = true

	url := "/cloudapi/compute/createTemplate"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
