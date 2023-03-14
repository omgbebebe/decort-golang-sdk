// API Actor for managing Disk. This actor is a final API for admin to manage Disk
package disks

import "repos.digitalenergy.online/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to disks
type Disks struct {
	client interfaces.Caller
}

// Builder for disks endpoints
func New(client interfaces.Caller) *Disks {
	return &Disks{
		client: client,
	}
}
