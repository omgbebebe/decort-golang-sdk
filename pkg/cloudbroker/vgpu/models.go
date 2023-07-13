package vgpu

type ItemVGPU struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	//Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// VGPU ID
	ID uint64 `json:"id"`

	// Last claimed by
	LastClaimedBy uint64 `json:"lastClaimedBy"`

	// Last update time
	LastUpdateTime uint64 `json:"lastUpdateTime"`

	// Mode
	Mode string `json:"mode"`

	// PCI Slot
	PCISlot interface{} `json:"pciSlot"`

	// PGPUID
	PGPUID uint64 `json:"pgpuid"`

	// Profile ID
	ProfileID interface{} `json:"profileId"`

	// RAM
	RAM uint64 `json:"ram"`

	// Reference ID
	ReferenceID interface{} `json:"referenceId"`

	// RGID
	RGID uint64 `json:"rgId"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`

	// VMID
	VMID uint64 `json:"vmid"`
}

// List of VGPU
type ListVGPU struct {
	// Data
	Data []ItemVGPU `json:"data"`

	// Entry count
	EntryCount uint64 `json:"entryCount"`
}
