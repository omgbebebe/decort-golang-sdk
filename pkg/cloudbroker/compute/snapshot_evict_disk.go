package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for evict specified disk
type SnapshotEvictDiskRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" valdiate:"required"`

	// ID of the disk instance
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`
}

// SnapshotEvictDisk evict specified disk from all snapshots of a compute instance
func (c Compute) SnapshotEvictDisk(ctx context.Context, req SnapshotEvictDiskRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/snapshotEvictDisk"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
