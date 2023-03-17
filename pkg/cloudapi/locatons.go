package cloudapi

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/locations"

// Accessing the Locations method group
func (ca *CloudAPI) Locations() *locations.Locations {
	return locations.New(ca.client)
}
