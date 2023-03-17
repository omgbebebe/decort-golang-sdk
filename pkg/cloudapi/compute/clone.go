package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for clone compute instance
type CloneRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Name of the clone
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Timestamp of the parent's snapshot to create clone from
	// Required: false
	SnapshotTimestamp uint64 `url:"snapshotTimestamp,omitempty" json:"snapshotTimestamp,omitempty"`

	// Name of the parent's snapshot to create clone from
	// Required: false
	SnapshotName string `url:"snapshotName,omitempty" json:"snapshotName,omitempty"`
}

// Clone clones compute instance
func (c Compute) Clone(ctx context.Context, req CloneRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/clone"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
