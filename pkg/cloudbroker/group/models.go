package group

type ItemGroup struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"meta"`

	// Is active
	Actice bool `json:"active"`

	// Description
	Description string `json:"description"`

	// Domain
	Domain string `json:"domain"`

	// GID
	GID uint64 `json:"gid"`

	// GUID
	GUID string `json:"guid"`

	// ID
	ID string

	// Last check
	LastCheck uint64 `json:"lastcheck"`

	// Roles
	Roles []interface{} `json:"roles"`

	// Users
	Users []string `json:"users"`
}

type ListGroups struct {
	Data []ItemGroup `json:"data"`

	EntryCount uint64 `json:"entryCount"`
}
