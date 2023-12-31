// Lists all the configured flavors available.
// A flavor is a combination of amount of compute capacity(CU) and disk capacity(GB).
package sizes

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
)

// Structure for creatig request to sizes
type Sizes struct {
	client interfaces.Caller
}

// Builder for sizes endpoints
func New(client interfaces.Caller) *Sizes {
	return &Sizes{
		client,
	}
}
