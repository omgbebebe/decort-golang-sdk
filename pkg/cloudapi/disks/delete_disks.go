package disks

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for multiple disks
type DisksDeleteRequest struct {
	// List of disk ids to delete
	// Required: true
	DisksIDs []uint64 `url:"diskIds" json:"diskIds" validate:"required"`

	// Reason for deleting the disks
	// Required: true
	Reason string `url:"reason" json:"reason" validate:"required"`

	// Whether to completely delete the disks, works only with non attached disks
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`
}

// DeleteDisks deletes multiple disks permanently
func (d Disks) DeleteDisks(ctx context.Context, req DisksDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/disks/deleteDisks"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
