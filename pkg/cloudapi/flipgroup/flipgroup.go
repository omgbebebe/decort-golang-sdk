// API to manage FLIPGroup instances
package flipgroup

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/interfaces"
)

// Structure for creating request to FLIPGroup
type FLIPGroup struct {
	client interfaces.Caller
}

// Builder for FLIPGroup endpoints
func New(client interfaces.Caller) *FLIPGroup {
	return &FLIPGroup{
		client,
	}
}
