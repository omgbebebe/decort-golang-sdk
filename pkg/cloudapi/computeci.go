package cloudapi

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/computeci"
)

// Accessing the ComputeCI method group
func (ca *CloudAPI) ComputeCI() *computeci.ComputeCI {
	return computeci.New(ca.client)
}
