package k8ci

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for disable K8CI
type DisableRequest struct {
	// K8CI ID
	// Required: true
	K8CIID uint64 `url:"k8ciId" json:"k8ciId" validate:"required"`
}

// Disable disables K8CI
func (k K8CI) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/k8ci/disable"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
