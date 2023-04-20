package k8s

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for add workers group
type WorkersGroupAddRequest struct {
	// Kubernetes cluster ID
	// Required: true
	K8SID uint64 `url:"k8sId" json:"k8sId" validate:"required"`

	// Worker group name
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// ID of SEP to create boot disks for default worker nodes group. Uses images SEP ID if not set
	// Required: false
	WorkerSEPID uint64 `url:"workerSepId,omitempty" json:"workerSepId,omitempty"`

	// Pool to use if worker SEP ID is set, can be also empty if needed to be chosen by system
	// Required: false
	WorkerSEPPool string `url:"workerSepPool,omitempty" json:"workerSepPool,omitempty"`

	// List of strings with labels for worker group
	// i.e: ["label1=value1", "label2=value2"]
	// Required: false
	Labels []string `url:"labels,omitempty" json:"labels,omitempty"`

	// List of strings with taints for worker group
	// i.e: ["key1=value1:NoSchedule", "key2=value2:NoExecute"]
	// Required: false
	Taints []string `url:"taints,omitempty" json:"taints,omitempty"`

	// List of strings with annotations for worker group
	// i.e: ["key1=value1", "key2=value2"]
	// Required: false
	Annotations []string `url:"annotations,omitempty" json:"annotations,omitempty"`

	// Number of worker nodes to create
	// Required: false
	WorkerNum uint64 `url:"workerNum,omitempty" json:"workerNum,omitempty"`

	// Worker node CPU count
	// Required: false
	WorkerCPU uint64 `url:"workerCpu,omitempty" json:"workerCpu,omitempty"`

	// Worker node RAM volume in MB
	// Required: false
	WorkerRAM uint64 `url:"workerRam,omitempty" json:"workerRam,omitempty"`

	// Worker node boot disk size in GB If 0 is specified, size is defined by the OS image size
	// Required: false
	WorkerDisk uint64 `url:"workerDisk,omitempty" json:"workerDisk,omitempty"`
}

// WorkersGroupAdd adds workers group to Kubernetes cluster
func (k8s K8S) WorkersGroupAdd(ctx context.Context, req WorkersGroupAddRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/k8s/workersGroupAdd"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
