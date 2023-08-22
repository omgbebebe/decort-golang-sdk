package stack

// Main information about stack
type InfoStack struct {
	// CKey
	Ckey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	//API URL
	APIURL string `json:"apiUrl"`

	//API key
	Apikey string `json:"apikey"`

	// App ID
	AppID string `json:"appId"`

	// CPU allocation ratio
	CPUAllocationRatio float64 `json:"cpu_allocation_ratio"`

	// Description
	Description string `json:"desc"`

	// Descr
	Descr string `json:"descr"`

	// Drivers
	Drivers []string `json:"drivers"`

	// Eco
	Eco interface{} `json:"eco"`

	// Error
	Error uint64 `json:"error"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`
	// List image IDs
	Images []uint64 `json:"images"`

	// Login
	Login string `json:"login"`

	// Mem allocation ratio
	MemAllocationRatio float64 `json:"mem_allocation_ratio"`

	// Name
	Name string `json:"name"`

	// Packegas
	Packages Packages `json:"packages"`

	//Password
	Password string `json:"passwd"`

	// Reference ID
	ReferenceID string `json:"referenceId"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`
}

// List of stacks
type ListStacks struct {

	//List
	Data []InfoStack `json:"data"`

	//Entry count
	EntryCount uint64 `json:"entryCount"`
}

// Package
type Packages struct {

	// LibvirtBin
	LibvirtBin LibvirtBin `json:"libvirt-bin"`

	// Lvm2Lockd
	Lvm2Lockd Lvm2Lockd `json:"lvm2-lockd"`

	// OpenvswitchCommon
	OpenvswitchCommon OpenvswitchCommon `json:"openvswitch-common"`

	// OpenvswitchSwitch
	OpenvswitchSwitch OpenvswitchSwitch `json:"openvswitch-switch"`

	// QemuSystemX86
	QemuSystemX86 QemuSystemX86 `json:"qemu-system-x86"`

	// Sanlock
	Sanlock Sanlock `json:"sanlock"`
}

// LibvirtBin
type LibvirtBin struct {

	// InstalledSize
	InstalledSize string `json:"installed_size"`

	// Version
	Ver string `json:"ver"`
}

// Lvm2Lockd
type Lvm2Lockd struct {

	// InstalledSize
	InstalledSize string `json:"installed_size"`

	// Version
	Ver string `json:"ver"`
}

// OpenvswitchCommon
type OpenvswitchCommon struct {

	// InstalledSize
	InstalledSize string `json:"installed_size"`

	// Version
	Ver string `json:"ver"`
}

// OpenvswitchSwitch
type OpenvswitchSwitch struct {

	// InstalledSize
	InstalledSize string `json:"installed_size"`

	// Version
	Ver string `json:"ver"`
}

// QemuSystemX86
type QemuSystemX86 struct {

	// InstalledSize
	InstalledSize string `json:"installed_size"`

	// Version
	Ver string `json:"ver"`
}

// Sanlock
type Sanlock struct {

	// InstalledSize
	InstalledSize string `json:"installed_size"`

	// Version
	Ver string `json:"ver"`
}
