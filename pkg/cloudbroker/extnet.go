package cloudbroker

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudbroker/extnet"
)

// Accessing the ExtNet method group
func (cb *CloudBroker) ExtNet() *extnet.ExtNet {
	return extnet.New(cb.client)
}
