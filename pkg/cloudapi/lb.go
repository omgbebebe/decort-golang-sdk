package cloudapi

import "repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudapi/lb"

// Accessing the LB method group
func (ca *CloudAPI) LB() *lb.LB {
	return lb.New(ca.client)
}
