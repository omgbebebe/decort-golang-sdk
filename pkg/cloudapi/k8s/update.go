package k8s

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for update kubernetes cluster
type UpdateRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// New name to set.
	// If empty string is passed, name is not updated
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// New description to set.
	// If empty string is passed, description is not updated
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

// Update updates name or description of Kubernetes cluster
func (k8s K8S) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/k8s/update"

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
