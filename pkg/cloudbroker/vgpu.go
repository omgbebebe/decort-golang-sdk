package cloudbroker

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/vgpu"

// Accessing the VGPU method group
func (cb *CloudBroker) VGPU() *vgpu.VGPU {
	return vgpu.New(cb.client)
}
