package cloudapi

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/sizes"
)

// Accessing the Sizes method group
func (ca *CloudAPI) Sizes() *sizes.Sizes {
	return sizes.New(ca.client)
}
