package k8s

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get information about group of kubernetes cluster
type FindGroupByLabelRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// List of labels to search
	// Required: true
	Labels []string `url:"labels" json:"labels" validate:"min=1"`

	// If true and more than one label provided, select only groups that have all provided labels.
	// If false - groups that have at least one label
	// Required: false
	Strict bool `url:"strict,omitempty" json:"strict,omitempty"`
}

// FindGroupByLabel find worker group information by one on more labels
func (k8s K8S) FindGroupByLabel(ctx context.Context, req FindGroupByLabelRequest) (ListK8SGroups, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/k8s/findGroupByLabel"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListK8SGroups{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
