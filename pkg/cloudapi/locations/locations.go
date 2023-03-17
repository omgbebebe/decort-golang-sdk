// API Actor api for managing locations
package locations

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
)

// Structure for creating request to locations
type Locations struct {
	client interfaces.Caller
}

// Builder for locations endpoints
func New(client interfaces.Caller) *Locations {
	return &Locations{
		client,
	}
}
