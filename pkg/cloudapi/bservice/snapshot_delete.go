package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete snapshot
type SnapshotDeleteRequest struct {
	// ID of the Basic Service
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// Label of the snapshot
	// Required: true
	Label string `url:"label" json:"label" validate:"required"`
}

// SnapshotDelete delete snapshot of the Basic Service
func (b BService) SnapshotDelete(ctx context.Context, req SnapshotDeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/snapshotDelete"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
