package pcidevice

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct to get list of pci devices
type ListRequest struct {
	// Find by id
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by computeId
	// Required: false
	ComputeID uint64 `url:"computeId,omitempty" json:"computeId,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by rgId
	// Required: false
	RGID uint64 `url:"rgId,omitempty" json:"rgId,omitempty"`

	// Find by status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of all pci devices as a ListPCIDevices struct
func (p PCIDevice) List(ctx context.Context, req ListRequest) (*ListPCIDevices, error) {
	res, err := p.ListRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	list := ListPCIDevices{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// ListRaw gets list of all pci devices as an array of bytes
func (p PCIDevice) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudbroker/pcidevice/list"

	res, err := p.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
