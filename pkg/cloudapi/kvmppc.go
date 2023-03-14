package cloudapi

import "repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudapi/kvmppc"

// Accessing the KVMPPC method group
func (ca *CloudAPI) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(ca.client)
}
