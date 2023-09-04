package grid

// Resource information
type Resources struct {
	// Current resources
	Current RecordResource `json:"Current"`

	// Reserved resources
	Reserved RecordResource `json:"Reserved"`
}

// Resource consumption information
type RecordResourcesConsumption struct {
	// Current resources
	Consumed RecordResource `json:"Consumed"`

	// Reserved resources
	Reserved RecordResource `json:"Reserved"`

	// GID
	GID uint64 `json:"id"`
}

type ListResourceConsumption struct {
	// Data
	Data []RecordResourcesConsumption `json:"data"`

	// Entry count
	EntryCount uint64 `json:"entryCount"`
}

// Resource details
type RecordResource struct {
	// Number of CPU
	CPU uint64 `json:"cpu"`

	// Disk size
	DiskSize float64 `json:"disksize"`

	// Disk size max
	DiskSizeMax int64 `json:"disksizemax"`

	// External IPs
	ExtIPs uint64 `json:"extips"`

	// External traffic
	ExtTraffic uint64 `json:"exttraffic"`

	// Number of GPU
	GPU uint64 `json:"gpu"`

	// Number of RAM
	RAM uint64 `json:"ram"`

	// SEPs
	SEPs map[string]map[string]DiskUsage `json:"seps"`
}

// Disk usage
type DiskUsage struct {
	// Disk size
	DiskSize float64 `json:"disksize"`

	// Disk size max
	DiskSizeMax float64 `json:"disksizemax"`
}

// Detailed information about grid
type RecordGrid struct {
	// Flag
	Flag string `json:"flag"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Location code
	LocationCode string `json:"locationCode"`

	// Name
	Name string `json:"name"`
}

// Information about grid
type ItemGridList struct {
	// Resource information
	Resources Resources `json:"Resources"`

	// Flag
	Flag string `json:"flag"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Location code
	LocationCode string `json:"locationCode"`

	// Name
	Name string `json:"name"`
}

// List Grids
type ListGrids struct {
	//Data
	Data []ItemGridList `json:"data"`

	// Entry count
	EntryCount uint64 `json:"entryCount"`
}

// List emails
type ListEmails struct {
	//Data
	Data []string `json:"data"`

	// Entry count
	EntryCount uint64 `json:"entryCount"`
}
