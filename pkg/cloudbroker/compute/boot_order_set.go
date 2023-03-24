package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for set boot order
type BootOrderSetRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// List of boot devices
	// Should be one of:
	//	- cdrom
	//	- network
	//	- hd
	// Required: true
	Order []string `url:"order" json:"order" validate:"min=1,computeOrder"`
}

// BootOrderSet sets compute boot order
func (c Compute) BootOrderSet(ctx context.Context, req BootOrderSetRequest) ([]string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/bootOrderSet"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	orders := make([]string, 0)

	err = json.Unmarshal(res, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
