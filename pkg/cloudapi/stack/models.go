package stack

// Main information about stack
type InfoStack struct {
	// CPU allocation ratio
	CPUAllocationRatio float64 `json:"cpu_allocation_ratio"`

	// Descr
	Descr string `json:"descr"`

	// Drivers
	Drivers []string `json:"drivers"`

	// ID
	ID uint64 `json:"id"`

	// Mem allocation ratio
	MemAllocationRatio float64 `json:"mem_allocation_ratio"`

	// Name
	Name string `json:"name"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`
}

// Information about stack in list
type ItemStack struct {
	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`
}

// List of stacks
type ListStacks struct {

	//List
	Data []ItemStack `json:"data"`

	//Entry count
	EntryCount uint64 `json:"entryCount"`
}
