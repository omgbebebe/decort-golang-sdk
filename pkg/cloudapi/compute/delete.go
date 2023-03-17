package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete compute
type DeleteRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Delete permanently
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`

	// Set True if you want to detach data disks (if any) from the compute before its deletion
	// Required: false
	DetachDisks bool `url:"detachDisks,omitempty" json:"detachDisks,omitempty"`
}

// Delete deletes compute
func (c Compute) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/delete"

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
