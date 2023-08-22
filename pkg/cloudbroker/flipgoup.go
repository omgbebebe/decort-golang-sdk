package cloudbroker

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/flipgroup"
)

// Accessing the FLIPGroup method group
func (cb *CloudBroker) FLIPGroup() *flipgroup.FLIPGroup {
	return flipgroup.New(cb.client)
}
