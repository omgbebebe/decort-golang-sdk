package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete snapshot
type SnapshotDeleteRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Text label of snapshot to delete
	// Required: true
	Label string `url:"label" json:"label" validate:"required"`
}

// SnapshotDelete delete specified compute snapshot
func (c Compute) SnapshotDelete(ctx context.Context, req SnapshotDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/snapshotDelete"

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
