package cloudbroker

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/k8ci"
)

// Accessing the K8CI method group
func (cb *CloudBroker) K8CI() *k8ci.K8CI {
	return k8ci.New(cb.client)
}
