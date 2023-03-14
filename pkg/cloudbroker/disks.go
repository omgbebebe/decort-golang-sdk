package cloudbroker

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudbroker/disks"
)

// Accessing the Disks method group
func (cb *CloudBroker) Disks() *disks.Disks {
	return disks.New(cb.client)
}
