package cloudbroker

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/vins"
)

// Accessing the VINS method group
func (cb *CloudBroker) VINS() *vins.VINS {
	return vins.New(cb.client)
}
