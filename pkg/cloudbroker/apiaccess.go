package cloudbroker

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/apiaccess"

// Accessing the APIAccess method group
func (cb *CloudBroker) APIAccess() *apiaccess.APIAccess {
	return apiaccess.New(cb.client)
}
