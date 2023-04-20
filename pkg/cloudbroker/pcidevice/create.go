package pcidevice

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
)

// Request struct for creating PCI device
type CreateRequest struct {
	// StackID
	// Required: true
	StackID uint64 `url:"stackId" json:"stackId" validate:"required"`

	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Name of device
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// PCI address of the device
	// Must be in format 0000:1f:2b.0
	// Required: true
	HWPath string `url:"hwPath" json:"hwPath" validate:"required,hwPath"`

	// Description, just for information
	// Required: false
	Description string `url:"description,omitempty" json:"description,omitempty"`
}

// Create creates PCI Device
func (p PCIDevice) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/pcidevice/create"

	res, err := p.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
