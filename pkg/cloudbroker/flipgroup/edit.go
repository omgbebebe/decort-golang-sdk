package flipgroup

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for edit FLIPGroup
type EditRequest struct {
	// FLIPGroup ID
	// Required: true
	FLIPGroupID uint64 `url:"flipgroupId" json:"flipgroupId" validate:"required"`

	// FLIPGroup name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// FLIPGroup description
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

// Edit edits FLIPGroup fields
func (f FLIPGroup) Edit(ctx context.Context, req EditRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/flipgroup/edit"

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
