package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest to get information about compute
type GetRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Reason to action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Get gets information about compute as a RecordCompute struct
func (c Compute) Get(ctx context.Context, req GetRequest) (*RecordCompute, error) {
	res, err := c.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordCompute{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets information about compute as an array of bytes
func (c Compute) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/get"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
