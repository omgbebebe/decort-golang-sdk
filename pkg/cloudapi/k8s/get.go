package k8s

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get detailed information about kubernetes cluster
type GetRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`
}

// Get gets information about Kubernetes cluster as a RecordK8S struct
func (k8s K8S) Get(ctx context.Context, req GetRequest) (*RecordK8S, error) {
	res, err := k8s.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordK8S{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets information about Kubernetes cluster as an array of bytes
func (k8s K8S) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/k8s/get"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
