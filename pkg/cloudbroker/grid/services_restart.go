package grid

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for restart services
type ServicesRestartRequest struct {
	// Grid (platform) ID
	// Required: true
    GID uint64 `url:"gid" json:"gid" validate:"required"`

	// Node ID
	// Required: true
    NID uint64 `url:"nid" json:"nid" validate:"required"`
}

// ServicesRestart restarts decort services on the node
func (g Grid) ServicesRestart(ctx context.Context, req ServicesRestartRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/servicesRestart"

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
