package cloudapi

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/extnet"
)

// Accessing the ExtNet method group
func (ca *CloudAPI) ExtNet() *extnet.ExtNet {
	return extnet.New(ca.client)
}
