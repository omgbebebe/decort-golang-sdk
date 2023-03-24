package k8s

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get node taints
type GetNodeTaintsRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// Node ID
	// Required: true
	NodeID uint64 `url:"nodeId" json:"nodeId" validate:"required"`
}

// GetNodeTaints gets kubernetes cluster worker node taints
func (k8s K8S) GetNodeTaints(ctx context.Context, req GetNodeTaintsRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/k8s/getNodeTaints"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
