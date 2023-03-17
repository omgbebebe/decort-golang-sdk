package cloudbroker

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/lb"

// Accessing the LB method group
func (cb *CloudBroker) LB() *lb.LB {
	return lb.New(cb.client)
}
