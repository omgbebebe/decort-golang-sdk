package cloudapi

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/kvmppc"

// Accessing the KVMPPC method group
func (ca *CloudAPI) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(ca.client)
}
