package pcidevice

import "repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to PCI device
type PCIDevice struct {
	client interfaces.Caller
}

// Builder for PCI device endpoints
func New(client interfaces.Caller) *PCIDevice {
	return &PCIDevice{
		client: client,
	}
}
