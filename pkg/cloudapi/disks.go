package cloudapi

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/disks"
)

// Accessing the Disks method group
func (ca *CloudAPI) Disks() *disks.Disks {
	return disks.New(ca.client)
}
