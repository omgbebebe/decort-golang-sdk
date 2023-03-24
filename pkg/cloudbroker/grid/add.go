package grid

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for location code
type AddRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid" json:"gid" validate:"required"`

	// Name of the location
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Location code typicly used in dns names
	// Required: true
	LocationCode string `url:"locationcode" json:"locationcode" validate:"required"`
}

// Add location code (e.g. DNS name of this grid)
func (g Grid) Add(ctx context.Context, req AddRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/add"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
