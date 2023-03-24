package vins

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete several VINSes
type MassDeleteRequest struct {
	// VINS IDs
	// Required: true
	VINSIDs []uint64 `url:"vinsIds" json:"vinsIds" validate:"min=1"`

	// Set to true if you want force delete non-empty VINS. Primarily,
	// VINS is considered non-empty if it has VMs connected to it,
	// and force flag will detach them from the VINS being deleted.
	// Otherwise method will return an error
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`

	// Set to true if you want to destroy VINS and all linked resources, if any, immediately.
	// Otherwise, they will be placed into recycle bin and could be restored later
	// within the recycle bins purge period
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// MassDelete start jobs to delete several VINSes
func (v VINS) MassDelete(ctx context.Context, req MassDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/massDelete"

	_, err = v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
