// API Actor api, this actor is the final api a enduser uses to manage his resources
package disks

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
)

// Structure for creating request to disks
type Disks struct {
	client interfaces.Caller
}

// Builder for disks endpoints
func New(client interfaces.Caller) *Disks {
	return &Disks{
		client,
	}
}
