package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete basic service
type DeleteRequest struct {
	// ID of the BasicService to be delete
	// Required: true
    ServiceID uint64 `url:"serviceId" json:"serviceId" validate:"required"`

	// If set to False, Basic service will be deleted to recycle bin. Otherwise destroyed immediately
	// Required: false
	Permanently bool `url:"permanently,omitempty" json:"permanently,omitempty"`
}

// Delete deletes BasicService instance
func (b BService) Delete(ctx context.Context, req DeleteRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/delete"

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
