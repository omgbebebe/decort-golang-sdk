package k8s

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for add worker to a kubernetes cluster
type UpdateWorkerNodesMetaDataRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// ID of the workers compute group
	// Required: true
	WorkersGroupID uint64 `url:"workersGroupId" json:"workersGroupId" validate:"required"`

	// Meta data for working group computes, format YAML "user_data": 1111
	// Required: true
	UserData string `url:"userData" json:"userData" validate:"required"`
}

// WorkerAdd adds worker nodes to a kubernetes cluster
func (k K8S) UpdateWorkerNodesMetaData(ctx context.Context, req UpdateWorkerNodesMetaDataRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/k8s/updateWorkerNodesMetaData"

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
