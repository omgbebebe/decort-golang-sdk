package cloudapi

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/bservice"

// Accessing the BService method group
func (ca *CloudAPI) BService() *bservice.BService {
	return bservice.New(ca.client)
}
