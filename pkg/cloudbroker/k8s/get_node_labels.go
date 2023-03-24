package k8s

import (
	"context"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get worker node labels
type GetNodeLabelsRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// Compute ID of worker node
	// Required: true
	NodeID uint64 `url:"nodeId" json:"nodeId" validate:"required"`
}

// GetNodeLabels gets kubernetes cluster worker node labels
func (k K8S) GetNodeLabels(ctx context.Context, req GetNodeLabelsRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/k8s/getNodeLabels"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
