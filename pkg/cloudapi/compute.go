package cloudapi

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudapi/compute"
)

// Accessing the Compute method group
func (ca *CloudAPI) Compute() *compute.Compute {
	return compute.New(ca.client)
}
