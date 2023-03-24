package vins

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for enable several VINSes
type MassEnableRequest struct {
	// VINS IDs
	// Required: true
	VINSIDs []uint64 `url:"vinsIds" json:"vinsIds" validate:"min=1"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// MassEnable start jobs to enable several VINSes
func (v VINS) MassEnable(ctx context.Context, req MassEnableRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/massEnable"

	_, err = v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
