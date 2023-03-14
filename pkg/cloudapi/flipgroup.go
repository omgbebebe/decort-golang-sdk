package cloudapi

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudapi/flipgroup"
)

// Accessing the FLIPGroup method group
func (ca *CloudAPI) FLIPGroup() *flipgroup.FLIPGroup {
	return flipgroup.New(ca.client)
}
