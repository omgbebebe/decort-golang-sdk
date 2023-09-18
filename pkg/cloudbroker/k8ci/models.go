package k8ci

// Main information about K8CI in List
type ItemK8CI struct {
	// Created time
	CreatedTime uint64 `json:"createdTime"`
	// Detailed information about K8CI
	RecordK8CIList
}

// List K8CI
type ListK8CI struct {
	//Data
	Data []ItemK8CI `json:"data"`

	// Entry count
	EntryCount uint64 `json:"entryCount"`
}

// Detailed information about K8CI in List
type RecordK8CIList struct {
	// Description
	Description string `json:"desc"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Load balancer image ID
	LBImageID uint64 `json:"lbImageId"`

	// Master driver
	MasterDriver string `json:"masterDriver"`

	// Master image ID
	MasterImageID uint64 `json:"masterImageId"`

	// Max master count
	MaxMasterCount uint64 `json:"maxMasterCount"`

	// Max worker count
	MaxWorkerCount uint64 `json:"maxWorkerCount"`

	// Name
	Name string `json:"name"`

	// Shared with
	SharedWith []uint64 `json:"sharedWith"`

	// Status
	Status string `json:"status"`

	// Version
	Version string `json:"version"`

	// Worker driver
	WorkerDriver string `json:"workerDriver"`

	// Worker image ID
	WorkerImageID uint64 `json:"workerImageId"`
}

// Detailed information about K8CI 
type RecordK8CI struct {
	// Description
	Description string `json:"desc"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Load balancer image ID
	LBImageID uint64 `json:"lbImageId"`

	// Master driver
	MasterDriver string `json:"masterDriver"`

	// Master image ID
	MasterImageID uint64 `json:"masterImageId"`

	// Max master count
	MaxMasterCount uint64 `json:"maxMasterCount"`

	// Max worker count
	MaxWorkerCount uint64 `json:"maxWorkerCount"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	//NetworkPlugins
	NetworkPlugins []string `json:"networkPlugins"`

	// Shared with
	SharedWith []uint64 `json:"sharedWith"`

	// Status
	Status string `json:"status"`

	// Version
	Version string `json:"version"`

	// Worker driver
	WorkerDriver string `json:"workerDriver"`

	// Worker image ID
	WorkerImageID uint64 `json:"workerImageId"`
}
