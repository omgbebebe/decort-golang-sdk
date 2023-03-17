package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get audit records
type AuditsRequest struct {
	// ID of the compute
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// Audits gets audit records for the specified compute object
func (c Compute) Audits(ctx context.Context, req AuditsRequest) (ListAudits, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/audits"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListAudits{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
