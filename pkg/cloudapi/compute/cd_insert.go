package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for insert new CD image
type CDInsertRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// ID of CD-ROM image
	// Required: true
	CDROMID uint64 `url:"cdromId" json:"cdromId" validate:"required"`
}

// CDInsert insert new CD image to compute's CD-ROM
func (c Compute) CDInsert(ctx context.Context, req CDInsertRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/cdInsert"

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
