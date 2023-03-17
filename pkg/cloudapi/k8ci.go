package cloudapi

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/k8ci"
)

// Accessing the K8CI method group
func (ca *CloudAPI) K8CI() *k8ci.K8CI {
	return k8ci.New(ca.client)
}
