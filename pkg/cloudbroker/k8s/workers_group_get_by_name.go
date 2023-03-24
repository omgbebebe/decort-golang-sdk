package k8s

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get information about worker group
type WorkersGroupGetByNameRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// Worker group name
	// Required: true
	GroupName string `url:"groupName" json:"groupName" validate:"required"`
}

// WorkersGroupGetByName gets worker group metadata by name
func (k K8S) WorkersGroupGetByName(ctx context.Context, req WorkersGroupGetByNameRequest) (*RecordK8SGroup, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/k8s/workersGroupGetByName"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordK8SGroup{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
