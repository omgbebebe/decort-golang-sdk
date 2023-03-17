package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for remove parent Compute Group
// relation to the specified Compute Group
type GroupParentRemoveRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
    ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// ID of the Compute Group
	// Required: true
    CompGroupID uint64 `url:"compgroupId" json:"compgroupId" validate:"required"`

	// ID of the parent Compute Group
	// to remove from the current Compute Group
	// Required: true
    ParentID uint64 `url:"parentId" json:"parentId" validate:"required"`
}

// GroupParentRemove removes parent Compute Group relation to the specified Compute Group
func (b BService) GroupParentRemove(ctx context.Context, req GroupParentRemoveRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/groupParentRemove"

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
