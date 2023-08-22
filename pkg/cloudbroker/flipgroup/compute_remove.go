package flipgroup

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for remove compute instance
type ComputeRemoveRequest struct {
	// ID of the Floating IP group to remove compute instance from
	// Required: true
	FLIPGroupID uint64 `url:"flipgroupId" json:"flipgroupId" validate:"required"`

	// ID of the compute instance to remove
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// ComputeRemove remove compute instance from the Floating IP group
func (f FLIPGroup) ComputeRemove(ctx context.Context, req ComputeRemoveRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/flipgroup/computeRemove"

	res, err := f.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
