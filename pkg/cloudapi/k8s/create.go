package k8s

import (
	"context"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// type Params []string

// Request struct for create kubernetes cluster
type CreateRequest struct {
	// Name of Kubernetes cluster
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Resource Group ID for cluster placement
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// ID of Kubernetes catalog item (k8sci) for cluster
	// Required: true
	K8SCIID uint64 `url:"k8ciId" json:"k8ciId" validate:"required"`

	// Name for first worker group created with cluster
	// Required: true
	WorkerGroupName string `url:"workerGroupName" json:"workerGroupName" validate:"required,workerGroupName"`

	// Network plugin
	// Must be one of these values: flannel, weawenet, calico
	// Required: true
	NetworkPlugin string `url:"networkPlugin" json:"networkPlugin" validate:"required,networkPlugin"`

	// ID of SEP to create boot disks for master nodes. Uses images SEP ID if not set
	// Required: false
	MasterSEPID uint64 `url:"masterSepId,omitempty" json:"masterSepId,omitempty"`

	// Pool to use if master SEP ID is set, can be also empty if needed to be chosen by system
	// Required: false
	MasterSEPPool string `url:"masterSepPool,omitempty" json:"masterSepPool,omitempty"`

	// ID of SEP to create boot disks for default worker nodes group. Uses images SEP ID if not set
	// Required: false
	WorkerSEPID uint64 `url:"workerSepId,omitempty" json:"workerSepId,omitempty"`

	// Pool to use if worker SEP ID is set, can be also empty if needed to be chosen by system
	// Required: false
	WorkerSEPPool string `url:"workerSepPool,omitempty" json:"workerSepPool,omitempty"`

	// List of strings with labels for default worker group
	// i.e: ["label1=value1", "label2=value2"]
	// Required: false
	Labels []string `url:"labels,omitempty" json:"labels,omitempty"`

	// List of strings with taints for default worker group
	// i.e: ["key1=value1:NoSchedule", "key2=value2:NoExecute"]
	// Required: false
	Taints []string `url:"taints,omitempty" json:"taints,omitempty"`

	// List of strings with annotations for worker group
	// i.e: ["key1=value1", "key2=value2"]
	// Required: false
	Annotations []string `url:"annotations,omitempty" json:"annotations,omitempty"`

	// Number of master nodes to create
	// Required: false
	MasterNum uint `url:"masterNum,omitempty" json:"masterNum,omitempty"`

	// Master node CPU count
	// Required: false
	MasterCPU uint `url:"masterCpu,omitempty" json:"masterCpu,omitempty"`

	// Master node RAM volume in MB
	// Required: false
	MasterRAM uint `url:"masterRam,omitempty" json:"masterRam,omitempty"`

	// Master node boot disk size in GB If 0 is specified, size is defined by the OS image size
	// Required: false
	MasterDisk uint `url:"masterDisk,omitempty" json:"masterDisk,omitempty"`

	// Number of worker nodes to create in default worker group
	// Required: false
	WorkerNum uint `url:"workerNum,omitempty" json:"workerNum,omitempty"`

	// Worker node CPU count
	// Required: false
	WorkerCPU uint `url:"workerCpu,omitempty" json:"workerCpu,omitempty"`

	// Worker node RAM volume in MB
	// Required: false
	WorkerRAM uint `url:"workerRam,omitempty" json:"workerRam,omitempty"`

	// Worker node boot disk size in GB. If 0 is specified, size is defined by the OS image size
	// Required: false
	WorkerDisk uint `url:"workerDisk,omitempty" json:"workerDisk,omitempty"`

	// ID of the external network to connect load balancer and cluster ViNS. If 0 is specified, external network selects automatically to
	// Required: false
	ExtNetID uint64 `url:"extnetId,omitempty" json:"extnetId,omitempty"`

	// ID of the ViNS to connect k8s cluster. If nothing is specified, ViNS will be created automatically
	// Required: false
	VinsId uint64 `url:"vinsId,omitempty" json:"vinsId,omitempty"`

	// Create Kubernetes cluster with masters nodes behind load balancer if true.
	// Otherwise give all cluster nodes direct external addresses from selected ExtNet
	// Required: false
	WithLB bool `url:"withLB" json:"withLB"`

	// Custom sysctl values for Load Balancer instance. Applied on boot
	// Required: false
	LbSysctlParams string `url:"-" json:"lbSysctlParams,omitempty" validate:"omitempty,dive"`

	// Use Highly Available schema for LB deploy
	// Required: false
	HighlyAvailable bool `url:"highlyAvailable,omitempty" json:"highlyAvailable,omitempty"`

	// Optional extra Subject Alternative Names (SANs) to use for the API Server serving certificate. Can be both IP addresses and DNS names
	// Required: false
	AdditionalSANs []string `url:"additionalSANs,omitempty" json:"additionalSANs,omitempty"`

	// Is used to define settings and actions that should be performed before any other component in the cluster starts.
	// It allows you to configure things like node registration, network setup, and other initialization tasks. insert a valid JSON string with all levels of nesting
	// Required: false
	InitConfiguration string `url:"initConfiguration,omitempty" json:"initConfiguration,omitempty"`

	// Is used to define global settings and configurations for the entire cluster.
	// It includes parameters such as cluster name, DNS settings, authentication methods, and other cluster-wide configurations.
	// Insert a valid JSON string with all levels of nesting
	// Required: false
	ClusterConfiguration string `url:"clusterConfiguration,omitempty" json:"clusterConfiguration,omitempty"`

	// Is used to configure the behavior and settings of the Kubelet, which is the primary node agent that runs on each node in the cluster.
	// It includes parameters such as node IP address, resource allocation, pod eviction policies, and other Kubelet-specific configurations.
	// Insert a valid JSON string with all levels of nesting
	// Required: false
	KubeletConfiguration string `url:"kubeletConfiguration,omitempty" json:"kubeletConfiguration,omitempty"`

	// Is used to configure the behavior and settings of the Kube-proxy, which is responsible for network proxying and load balancing within the cluster.
	// It includes parameters such as proxy mode, cluster IP ranges, and other Kube-proxy specific configurations.
	// Insert a valid JSON string with all levels of nesting
	// Required: false
	KubeProxyConfiguration string `url:"kubeProxyConfiguration,omitempty" json:"kubeProxyConfiguration,omitempty"`

	// Is used to configure the behavior and settings for joining a node to a cluster.
	// It includes parameters such as the cluster's control plane endpoint, token, and certificate key. insert a valid JSON string with all levels of nesting
	// Required: false
	JoinConfiguration string `url:"joinConfiguration,omitempty" json:"joinConfiguration,omitempty"`

	// Text description of this Kubernetes cluster
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Meta data for working group computes, format YAML "user_data": 1111
	// Required: false
	UserData string `url:"userData,omitempty" json:"userData,omitempty"`

	// Use only selected ExtNet for infrastructure connections
	// Required: false
	ExtNetOnly bool `url:"extnetOnly,omitempty" json:"extnetOnly,omitempty"`

	// Insert ssl certificate in x509 pem format
	// Required: false
	OidcCertificate string `url:"oidcCertificate,omitempty" json:"oidcCertificate,omitempty"`
}

// type wrapperCreateRequest struct {
// 	CreateRequest
// 	Params []string `url:"lbSysctlParams,omitempty"`
// }

// Create creates a new Kubernetes cluster in the specified Resource Group
func (k8s K8S) Create(ctx context.Context, req CreateRequest) (string, error) {

	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	// var params []string

	// if len(req.LbSysctlParams) != 0 {
	// 	params = make([]string, 0, len(req.LbSysctlParams))

	// 	for r := range req.LbSysctlParams {
	// 		b, err := json.Marshal(req.LbSysctlParams[r])
	// 		if err != nil {
	// 			return "", err
	// 		}

	// 		params = append(params, string(b))
	// 	}
	// } else {
	// 	params = []string{"[]"}
	// }

	// reqWrapped := wrapperCreateRequest{
	// 	CreateRequest: req,
	// 	Params:        params,
	// }

	url := "/cloudapi/k8s/create"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
