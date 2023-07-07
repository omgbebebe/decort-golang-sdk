package k8ci

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create K8CI instance
type CreateRequest struct {
	// Name of catalog item
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Version tag
	// Required: true
	Version string `url:"version" json:"version" validate:"required"`

	// Optional description
	// Required: false
	Description string `url:"description,omitempty" json:"description,omitempty"`

	// Image ID for master K8S node
	// Required: true
	MasterImageID uint64 `url:"masterImageId" json:"masterImageId" validate:"required"`

	// Compute driver
	// Should be one of:
	//	- KVM_X86
	//	- KVM_PPC
	//	- etc
	// Required: true
	MasterDriver string `url:"masterDriver" json:"masterDriver" validate:"driver"`

	// Image ID for worker K8S node
	// Required: true
	WorkerImageID uint64 `url:"workerImageId" json:"workerImageId" validate:"required"`

	// Compute driver
	// Should be one of
	//	- KVM_X86
	//	- KVM_PPC
	//	- etc
	// Required: true
	WorkerDriver string `url:"workerDriver" json:"workerDriver" validate:"driver"`

	// List of account IDs, which have access to this item.
	// If empty, any account has access
	// Required: false
	SharedWith []uint64 `url:"sharedWith,omitempty" json:"sharedWith,omitempty"`

	// Policy limit on maximum number of master nodes
	// Required: true
	MaxMasterCount uint64 `url:"maxMasterCount" json:"maxMasterCount" validate:"required"`

	// Policy limit on maximum number of worker nodes
	// Required: true
	MaxWorkerCount uint64 `url:"maxWorkerCount" json:"maxWorkerCount" validate:"required"`

	// Network plugins
	// Values of slice must be flannel, weawenet or calico
	//Required: true
	NetworkPlugins []string `url:"networkPlugins" json:"networkPlugins" validate:"required,networkPlugins"`
}

// Create creates a new K8CI instance
func (k K8CI) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/k8ci/create"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
