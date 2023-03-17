package cloudapi

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/image"
)

// Accessing the Image method group
func (ca *CloudAPI) Image() *image.Image {
	return image.New(ca.client)
}
