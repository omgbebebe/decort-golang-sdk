package vins

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for disable several VINSes
type MassDisableRequest struct {
	// VINS IDs
	// Required: true
	VINSIDs []uint64 `url:"vinsIds" json:"vinsIds" validate:"min=1"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// MassDisable start jobs to disable several VINSes
func (v VINS) MassDisable(ctx context.Context, req MassDisableRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/massDisable"

	_, err = v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
