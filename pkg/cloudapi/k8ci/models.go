package k8ci

// Detailed information about K8CI
type ItemK8CI struct {
	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Description
	Description string `json:"desc"`

	// ID
	ID uint64 `json:"id"`

	// LB image ID
	LBImageID uint64 `json:"lbImageId"`

	// Name
	Name string `json:"name"`

	// Network plugins
	NetworkPlugins []string `json:"networkPlugins"`

	// Status
	Status string `json:"status"`

	// Version
	Version string `json:"version"`
}

// List of K8CI
type ListK8CI struct {
	Data []ItemK8CI `json:"data"`

	EntryCount uint64 `json:"entryCount"`
}

// Main information about K8CI
type RecordK8CI struct {
	// Description
	Description string `json:"desc"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Network plugins
	NetworkPlugins []string `json:"networkPlugins"`

	// Version
	Version string `json:"version"`
}
