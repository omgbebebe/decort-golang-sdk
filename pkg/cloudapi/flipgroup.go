package cloudapi

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/flipgroup"
)

// Accessing the FLIPGroup method group
func (ca *CloudAPI) FLIPGroup() *flipgroup.FLIPGroup {
	return flipgroup.New(ca.client)
}
