package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list snapshots
type SnapshotListRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// SnapshotList gets list compute snapshots
func (c Compute) SnapshotList(ctx context.Context, req SnapshotListRequest) (ListSnapShots, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/snapshotList"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListSnapShots{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
