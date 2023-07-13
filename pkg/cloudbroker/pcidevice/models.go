package pcidevice

// Main information about PCI device
type ItemPCIDevice struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// Compute ID
	ComputeID uint64 `json:"computeId"`

	// Description
	Description string `json:"description"`

	// GUID
	GUID uint64 `json:"guid"`

	// HwPath
	HwPath string `json:"hwPath"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Stack ID
	StackID uint64 `json:"stackId"`

	// Status
	Status string `json:"status"`

	// System name
	SystemName string `json:"systemName"`
}

// List PCI devices
type ListPCIDevices struct {
	// Data
	Data []ItemPCIDevice `json:"data"`

	// Entry count
	EntryCount uint64 `json:"entryCount"`
}

