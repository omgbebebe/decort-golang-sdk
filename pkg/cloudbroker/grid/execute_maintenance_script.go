package grid

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for execute script
type ExecuteMaintenanceScriptRequest struct {
	// Grid (platform) ID
	// Required: true
	GID string `url:"gid" json:"gid" validate:"required"`

	// Type of nodes you want to apply the action on
	// Required: true
	NodesType string `url:"nodestype" json:"nodestype" validate:"required"`

	// The script you want to run
	// Required: true
	Script string `url:"script" json:"script" validate:"required"`
}

// ExecuteMaintenanceScript executes maintenance script
func (g Grid) ExecuteMaintenanceScript(ctx context.Context, req ExecuteMaintenanceScriptRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/executeMaintenanceScript"

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
