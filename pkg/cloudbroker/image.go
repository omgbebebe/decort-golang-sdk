package cloudbroker

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/image"
)

// Accessing the Image method group
func (cb *CloudBroker) Image() *image.Image {
	return image.New(cb.client)
}
