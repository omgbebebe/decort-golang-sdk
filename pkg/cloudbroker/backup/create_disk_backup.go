package backup

import (
	"context"
	"encoding/json"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strings"
)

// Request struct for creating disk backup
type CreateDiskBackupRequest struct {
	// Compute ID
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Disk ID
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// Backup path
	// Required: true
	BackupPath string `url:"backupPath" json:"backupPath" validate:"required"`

	// Async API Call
	// For async call use CreateDiskBackupAsync
	// For sync call use CreateDiskBackup
	// Required: true
	async bool `url:"async"`
}

// CreateDiskBackup creates disk backup
func (b Backup) CreateDiskBackup(ctx context.Context, req CreateDiskBackupRequest) (ListInfoBackup, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	req.async = false

	url := "/cloudbroker/backup/createDiskBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	result := make(ListInfoBackup, 0)

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateDiskBackupAsync creates disk backup
func (b Backup) CreateDiskBackupAsync(ctx context.Context, req CreateDiskBackupRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	req.async = true

	url := "/cloudbroker/backup/createDiskBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
