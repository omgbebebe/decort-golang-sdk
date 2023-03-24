package vins

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete VINS
type DeleteRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// Set to True if you want force delete non-empty VINS.
	// Primarily, VINS is considered non-empty if it has virtual machines connected to it,
	// and force flag will detach them from the VINS being deleted.
	// Otherwise method will return an error
	// Required: false
	Force bool `url:"force,omitempty" json:"force,omitempty"`

	// Set to True if you want to destroy VINS and all linked resources, if any, immediately.
	// Otherwise, they will be placed into recycle bin and could be restored later within the recycle bin's purge period
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Delete deletes VINS
func (v VINS) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/delete"

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
