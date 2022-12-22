package k8s

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for hard reset kubernetes cluster
type WorkerResetRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId"`

	// ID of the workers compute group
	// Required: true
	WorkersGroupID uint64 `url:"workersGroupId"`

	// Compute ID of worker node to reset
	// Required: true
	WorkerID uint64 `url:"workerId"`
}

func (krq WorkerResetRequest) validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID must be set")
	}
	if krq.WorkersGroupID == 0 {
		return errors.New("validation-error: field WorkersGroupID must be set")
	}
	if krq.WorkerID == 0 {
		return errors.New("validation-error: field WorkerID must be set")
	}

	return nil
}

// WorkerReset hard reset (compute start + stop) worker node of the kubernetes cluster
func (k K8S) WorkerReset(ctx context.Context, req WorkerResetRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/k8s/workerReset"

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
