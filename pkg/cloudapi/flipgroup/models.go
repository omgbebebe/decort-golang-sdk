package flipgroup

// Main information about FLIPGroup
type RecordFLIPGroup struct {
	// Default GW
	DefaultGW string `json:"defaultGW"`

	// ID
	ID uint64 `json:"id"`

	// IP
	IP string `json:"ip"`

	// Name
	Name string `json:"name"`

	// Network mask
	NetMask uint64 `json:"netmask"`
}

// Detailed information about FLIPGroup
type ItemFLIPGroup struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// List of client IDs
	ClientIDs []uint64 `json:"clientIds"`

	// Client type
	ClientType string `json:"clientType"`

	// Connection ID
	ConnID uint64 `json:"connId"`

	// Connection type
	ConnType string `json:"connType"`

	// Default GW
	DefaultGW string `json:"defaultGW"`

	// Description
	Description string `json:"desc"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// IP
	IP string `json:"ip"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Network ID
	NetID uint64 `json:"netId"`

	// Network type
	NetType string `json:"netType"`

	// NetMask
	NetMask uint64 `json:"netmask"`

	// Status
	Status string `json:"status"`
}

// List of FLIPGroup
type ListFLIPGroups struct {
	Data []ItemFLIPGroup `json:"data"`

	EntryCount uint64 `json:"entryCount"`
}
