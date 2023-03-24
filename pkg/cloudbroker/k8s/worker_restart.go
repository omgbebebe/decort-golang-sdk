package k8s

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for restart worker node
type WorkerRestartRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"requireD"`

	// ID of the workers compute group
	// Required: true
	WorkersGroupID uint64 `url:"workersGroupId" json:"workersGroupId" validate:"required"`

	// Compute ID of worker node to restart
	// Required: true
	WorkerID uint64 `url:"workerId" json:"workerId" validate:"required"`
}

// WorkerRestart soft restart (reboot OS) worker node of the kubernetes cluster
func (k8s K8S) WorkerRestart(ctx context.Context, req WorkerRestartRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/k8s/workerRestart"

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
