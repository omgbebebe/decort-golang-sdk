package cloudbroker

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/k8s"
)

// Accessing the K8S method group
func (cb *CloudBroker) K8S() *k8s.K8S {
	return k8s.New(cb.client)
}
