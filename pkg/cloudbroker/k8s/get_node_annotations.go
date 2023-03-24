package k8s

import (
	"context"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get node annotations
type GetNodeAnnotationsRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// Node ID
	// Required: true
	NodeID uint64 `url:"nodeId" json:"nodeId" validate:"required"`
}

// GetNodeAnnotations gets kubernetes cluster worker node annotations
func (k K8S) GetNodeAnnotations(ctx context.Context, req GetNodeAnnotationsRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/k8s/getNodeAnnotations"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
