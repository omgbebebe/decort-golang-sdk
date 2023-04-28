package compute

import (
	"context"
	"encoding/json"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for getting boot order
type BootOrderGetRequest struct {
	// Compute ID
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// BootOrderGet gets actual compute boot order information
func (c Compute) BootOrderGet(ctx context.Context, req BootOrderGetRequest) ([]string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/bootOrderGet"

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
