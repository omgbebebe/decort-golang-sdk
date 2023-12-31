package cloudapi

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/rg"

// Accessing the RG method group
func (ca *CloudAPI) RG() *rg.RG {
	return rg.New(ca.client)
}
