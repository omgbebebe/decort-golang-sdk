package bservice

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list existing snapshots
type SnapshotListRequest struct {
	// ID of the Basic Service
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`
}

// SnapshotList gets list existing snapshots of the Basic Service
func (b BService) SnapshotList(ctx context.Context, req SnapshotListRequest) (ListSnapshots, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/snapshotList"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListSnapshots{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
