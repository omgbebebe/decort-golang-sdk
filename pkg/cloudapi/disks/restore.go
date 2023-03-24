package disks

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for restore a deleted unattached disk
type RestoreRequest struct {
	// ID of the disk to restore
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// Reason for restoring the disk
	// Required: true
	Reason string `url:"reason" json:"reason" validate:"required"`
}

// Restore restore a deleted unattached disk from recycle bin
func (d Disks) Restore(ctx context.Context, req RestoreRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/disks/restore"

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
