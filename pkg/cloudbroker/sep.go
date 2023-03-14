package cloudbroker

import "repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudbroker/sep"

// Accessing the SEP method group
func (cb *CloudBroker) SEP() *sep.SEP {
	return sep.New(cb.client)
}
