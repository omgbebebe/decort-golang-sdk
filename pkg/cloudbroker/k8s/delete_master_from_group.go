package k8s

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete master from group
type DeleteMasterFromGroupRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// ID of the masters compute group
	// Required: true
	MasterGroupID uint64 `url:"masterGroupId" json:"masterGroupId" validate:"required"`

	// List of Compute IDs of master nodes to delete
	// Required: true
	MasterIDs []string `url:"masterIds" json:"masterIds" validate:"min=1"`
}

// DeleteMasterFromGroup deletes compute from masters group in selected kubernetes cluster
func (k K8S) DeleteMasterFromGroup(ctx context.Context, req DeleteMasterFromGroupRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/k8s/deleteMasterFromGroup"

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
