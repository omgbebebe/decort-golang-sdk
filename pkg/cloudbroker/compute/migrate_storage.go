package compute

import (
	"context"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for migration
type MigrateStorageRequest struct {
	// ID of the compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// SEP ID to migrate disks
	// Required: true
	SEPID uint64 `url:"sepId" json:"sepId" validate:"required"`

	// SEP pool name to migrate disks
	// Required: true
	PoolName string `url:"poolName" json:"poolName" validate:"required"`

	// Target stack ID
	// Required: true
	StackID uint64 `url:"stackId" json:"stackId" validate:"required"`

	// Async API call
	// Required: true
	Sync bool `url:"sync" json:"sync" validate:"required"`
}

// MigrateStorage gets complex compute migration
// Compute will be migrated to specified stack, and compute disks will
// be migrated to specified SEP to specified pool.
// This action can take up to 84 hours
func (c Compute) MigrateStorage(ctx context.Context, req MigrateStorageRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/migrateStorage"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
