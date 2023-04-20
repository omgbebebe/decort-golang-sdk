package vgpu

import "repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to VGPU
type VGPU struct {
	client interfaces.Caller
}

// Builder for VGPU endpoints
func New(client interfaces.Caller) *VGPU {
	return &VGPU{
		client: client,
	}
}
