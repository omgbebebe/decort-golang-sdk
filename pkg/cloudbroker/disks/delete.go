package disks

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request for delete disk
type DeleteRequest struct {
	// ID of disk to delete
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// Detach disk from machine first
	// Required: false
	Detach bool `url:"detach,omitempty" json:"detach,omitempty"`

	// Whether to completely delete the disk, works only with non attached disks
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`

	// Reason to delete
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Delete deletes disk by ID
func (d Disks) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/disks/delete"

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
