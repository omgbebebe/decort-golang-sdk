package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for detach PCI device
type DetachPCIDeviceRequest struct {
	// Identifier compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Pci device ID
	// Required: true
	DeviceID uint64 `url:"deviceId" json:"deviceId" validate:"required"`
}

// DetachPCIDevice detach PCI device
func (c Compute) DetachPCIDevice(ctx context.Context, req DetachPCIDeviceRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/detachPciDevice"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
