package vins

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for enable VINS
type EnableRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Enable enables VINS by ID
func (v VINS) Enable(ctx context.Context, req EnableRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/enable"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
