// List of method groups for the user
package cloudapi

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/interfaces"
)

// Structure for creating request to CloudAPI groups
type CloudAPI struct {
	client interfaces.Caller
}

// Builder to get access to  CloudAPI
func New(client interfaces.Caller) *CloudAPI {
	return &CloudAPI{
		client: client,
	}
}
