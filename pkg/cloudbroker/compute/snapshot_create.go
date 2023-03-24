package compute

import (
	"context"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create snapshot
type SnapshotCreateRequest struct {
	// ID of the compute instance to create snapshot for
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Text label for snapshot.
	// Must be unique among this compute snapshots
	// Required: true
	Label string `url:"label" json:"label" validate:"required"`
}

// SnapshotCreate create compute snapshot
func (c Compute) SnapshotCreate(ctx context.Context, req SnapshotCreateRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/snapshotCreate"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
