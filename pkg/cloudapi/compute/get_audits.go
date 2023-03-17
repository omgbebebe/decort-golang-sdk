package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get compute audits
type GetAuditsRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// GetAudits gets compute audits
func (c Compute) GetAudits(ctx context.Context, req GetAuditsRequest) (ListShortAudits, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/getAudits"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListShortAudits{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
