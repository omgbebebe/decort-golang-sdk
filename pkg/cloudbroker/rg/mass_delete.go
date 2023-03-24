package rg

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete several resource groups
type MassDeleteRequest struct {
	// IDs of the resource groups
	// Required: true
	RGIDs []uint64 `url:"rgIds" json:"rgIds" validate:"min=1"`

	// Set to true if you want force delete non-empty resource groups
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`

	// Set to true if you want to destroy resource group and all linked
	// resources, if any, immediately.
	// Otherwise, they will be placed into recycle bin and could be
	// restored later within recycle bins purge period
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// MassDelete starts jobs to delete several resource groups
func (r RG) MassDelete(ctx context.Context, req MassDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/rg/massDelete"

	_, err = r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
