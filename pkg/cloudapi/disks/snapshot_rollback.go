package disks

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for rollback snapshot
type SnapshotRollbackRequest struct {
	// ID of the disk
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// Label of the snapshot to rollback
	// Required: false
	Label string `url:"label,omitempty" json:"label,omitempty"`

	// Timestamp of the snapshot to rollback
	// Required: false
	TimeStamp uint64 `url:"timestamp,omitempty" json:"timestamp,omitempty"`
}

// SnapshotRollback rollback an individual disk snapshot
func (d Disks) SnapshotRollback(ctx context.Context, req SnapshotRollbackRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/disks/snapshotRollback"

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
