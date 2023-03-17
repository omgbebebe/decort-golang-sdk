package cloudbroker

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/kvmppc"

// Accessing the KVMPPC method group
func (cb *CloudBroker) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(cb.client)
}
