package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for grant access to compute
type UserGrantRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Name of the user to add
	// Required: true
	Username string `url:"userName" json:"userName" validate:"required"`

	// Access type
	// Should be one of:
	//	- 'R' for Read only
	//	- 'RCX' for Write
	//	- 'ARCXDU' for Admin
	// Required: true
	AccessType string `url:"accesstype" json:"accesstype" validate:"accountAccessType"`
}

// UserGrant grant user access to the compute
func (c Compute) UserGrant(ctx context.Context, req UserGrantRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/userGrant"

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
