package cloudbroker

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudbroker/compute"
)

// Accessing the Compute method group
func (cb *CloudBroker) Compute() *compute.Compute {
	return compute.New(cb.client)
}
