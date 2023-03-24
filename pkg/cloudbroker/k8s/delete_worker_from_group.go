package k8s

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete worker from group
type DeleteWorkerFromGroupRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// ID of the workers compute group
	// Required: true
	WorkersGroupID uint64 `url:"workersGroupId" json:"workersGroupId" validate:"required"`

	// Compute ID of worker node to delete
	// Required: true
	WorkerID uint64 `url:"workerId" json:"workerId" validate:"required"`
}

// DeleteWorkerFromGroup deletes worker compute from workers group in selected kubernetes cluster
func (k K8S) DeleteWorkerFromGroup(ctx context.Context, req DeleteWorkerFromGroupRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/k8s/deleteWorkerFromGroup"

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
