package grid

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for purge logs
type PurgeLogsRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid" json:"gid" validate:"required"`

	// Age of the records to remove, e.g. -1h for records older than 1 hour, -1w - one week, etc
	// Required: true
	Age string `url:"age" json:"age" validate:"required"`
}

// PurgeLogs clear Log and ECO records that are older than the specified age.
// By default, records older than one week are removed
func (g Grid) PurgeLogs(ctx context.Context, req PurgeLogsRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/purgeLogs"

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
