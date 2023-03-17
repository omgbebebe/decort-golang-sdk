package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for resize the group
type GroupResizeRequest struct {
	// ID of the Basic Service of Compute Group
	// Required: true
	ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// ID of the Compute Group to resize
	// Required: true
	CompGroupID uint64 `url:"compgroupId" json:"compgroupId" validate:"required"`

	// Either delta or absolute value of computes
	// Required: true
	Count int64 `url:"count" json:"count" validate:"required"`

	// Either delta or absolute value of computes
	// Should be one of:
	//	- ABSOLUTE
	//	- RELATIVE
	// Required: true
	Mode string `url:"mode" json:"mode" validate:"bserviceMode"`
}

// GroupResize resize the group by changing the number of computes
func (b BService) GroupResize(ctx context.Context, req GroupResizeRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/groupResize"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
