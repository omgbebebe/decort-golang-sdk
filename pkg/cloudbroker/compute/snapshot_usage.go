package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get compute snapshot real size on storage
type SnapshotUsageRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Specify to show usage exact for this snapshot.
	// Leave empty for get usage for all compute snapshots
	// Required: false
	Label string `url:"label,omitempty" json:"label,omitempty"`
}

// SnapshotUsage Get compute snapshot real size on storage.
// Always returns list of json objects, and first json object contains summary about all related
// snapshots.
func (c Compute) SnapshotUsage(ctx context.Context, req SnapshotUsageRequest) (ListSnapshotUsage, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/snapshotUsage"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListSnapshotUsage{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
