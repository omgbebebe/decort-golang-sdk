package cloudbroker

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/pcidevice"

// Accessing the PCI Device method group
func (cb *CloudBroker) PCIDevice() *pcidevice.PCIDevice {
	return pcidevice.New(cb.client)
}
