package cloudapi

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudapi/vins"
)

// Accessing the VINS method group
func (ca *CloudAPI) VINS() *vins.VINS {
	return vins.New(ca.client)
}
