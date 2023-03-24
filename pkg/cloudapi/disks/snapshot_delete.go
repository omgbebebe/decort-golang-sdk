package disks

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete snapshot
type SnapshotDeleteRequest struct {
	// ID of disk to delete
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// Label of the snapshot to delete
	// Required: true
	Label string `url:"label" json:"label" validate:"required"`
}

// SnapshotDelete deletes a snapshot
func (d Disks) SnapshotDelete(ctx context.Context, req SnapshotDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/disks/snapshotDelete"

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
