// API Actor for managing ComputeCI. This actor is a final API for admin to manage ComputeCI
package computeci

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
)

// Structure for creating request to computeci
type ComputeCI struct {
	client interfaces.Caller
}

// Builder for computeci endpoints
func New(client interfaces.Caller) *ComputeCI {
	return &ComputeCI{
		client,
	}
}
