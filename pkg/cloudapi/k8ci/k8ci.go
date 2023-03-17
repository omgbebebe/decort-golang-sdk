// API to manage K8CI instances
package k8ci

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
)

// Structure for creating request to K8CI
type K8CI struct {
	client interfaces.Caller
}

// Builder for K8CI endpoints
func New(client interfaces.Caller) *K8CI {
	return &K8CI{
		client,
	}
}
