// API Actor for managing Compute. This actor is a final API for admin to manage Compute
package compute

import "repos.digitalenergy.online/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to compute
type Compute struct {
	client interfaces.Caller
}

// Builder for compute endpoints
func New(client interfaces.Caller) *Compute {
	return &Compute{
		client: client,
	}
}
