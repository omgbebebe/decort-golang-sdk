package k8s

import (
	"context"
	"encoding/json"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for add worker to a kubernetes cluster
type WorkerAddRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// ID of the workers compute group
	// Required: true
	WorkersGroupID uint64 `url:"workersGroupId" json:"workersGroupId" validate:"required"`

	// How many worker nodes to add
	// Required: true
	Num uint64 `url:"num" json:"num" validate:"required"`
}

// WorkerAdd add worker nodes to a Kubernetes cluster
func (k8s K8S) WorkerAdd(ctx context.Context, req WorkerAddRequest) ([]uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/k8s/workerAdd"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	result := make([]uint64, 0)

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
