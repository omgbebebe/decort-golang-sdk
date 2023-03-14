package cloudbroker

import "repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudbroker/rg"

// Accessing the RG method group
func (cb *CloudBroker) RG() *rg.RG {
	return rg.New(cb.client)
}
