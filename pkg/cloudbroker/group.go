package cloudbroker

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/group"

// Accessing the Group method group
func (cb *CloudBroker) Group() *group.Group {
	return group.New(cb.client)
}
