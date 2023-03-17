package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for disable service
type DisableRequest struct {
	// ID of the service to disable
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`
}

// Disable disables service.
// Disabling a service technically means setting model status
// of all computes and service itself to DISABLED and stopping all computes.
func (b BService) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/delete"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
