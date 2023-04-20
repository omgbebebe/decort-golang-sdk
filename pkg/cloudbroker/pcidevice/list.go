package pcidevice

import (
	"context"
	"encoding/json"
	"net/http"
)

// List gets list all pci devices
func (p PCIDevice) List(ctx context.Context) (ListPCIDevices, error) {
	url := "/cloudbroker/pcidevice/list"

	res, err := p.client.DecortApiCall(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	list := ListPCIDevices{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
