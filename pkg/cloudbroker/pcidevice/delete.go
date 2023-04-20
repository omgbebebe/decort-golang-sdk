package pcidevice

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
)

// Request struct for deleting PCI device
type DeleteRequest struct {
	// PCI device ID
	// Required: true
	DeviceID uint64 `url:"deviceId" json:"deviceId" validate:"required"`

	// Force delete
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`
}

// Delete PCI device
func (p PCIDevice) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/pcidevice/delete"

	res, err := p.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
