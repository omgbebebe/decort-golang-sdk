package compute

import (
	"context"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for cleanup resources after finished migration
type MigrateStorageCleanUpRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`
}

// MigrateStorageCleanUp cleanup resources after finished (success of failed) complex compute migration.
// If the migration was successful, then old disks will be removed, else new (target) disks will be removed.
// Do it wisely!
func (c Compute) MigrateStorageCleanUp(ctx context.Context, req MigrateStorageCleanUpRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/migrateStorageCleanup"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
