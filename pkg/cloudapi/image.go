package cloudapi

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudapi/image"
)

// Accessing the Image method group
func (ca *CloudAPI) Image() *image.Image {
	return image.New(ca.client)
}
