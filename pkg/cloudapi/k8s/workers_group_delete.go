package k8s

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete workers group
type WorkersGroupDeleteRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// Worker group ID
	// Required: true
	WorkersGroupID uint64 `url:"workersGroupId" json:"workersGroupId" validate:"required"`
}

// WorkersGroupDelete deletes workers group from Kubernetes cluster
func (k8s K8S) WorkersGroupDelete(ctx context.Context, req WorkersGroupDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/k8s/workersGroupDelete"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
