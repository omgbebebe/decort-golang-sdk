// API to manage KVM PowerPC compute instances (PPC VMs)
package kvmppc

import "repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to KVMPPC
type KVMPPC struct {
	client interfaces.Caller
}

// Builder for KVMPPC endpoints
func New(client interfaces.Caller) *KVMPPC {
	return &KVMPPC{
		client: client,
	}
}
