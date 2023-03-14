// Lists all the images. A image is a template which can be used to deploy machines
package image

import "repos.digitalenergy.online/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to image
type Image struct {
	client interfaces.Caller
}

// Builder for image endpoint
func New(client interfaces.Caller) *Image {
	return &Image{
		client: client,
	}
}
