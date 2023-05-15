package backup

import (
	"context"
	"encoding/json"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strings"
)

// Request struct for restoring disk from backup
type RestoreDiskFromBackupRequest struct {
	// Compute ID
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Disk ID
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// Backup path
	BackupPath string `url:"backupPath" json:"backupPath" validate:"required"`

	// Backup file
	BackupFile string `url:"backupFile" json:"backupFile" validate:"required"`

	// Async API Call
	// For async call use RestoreDiskFromBackupAsync
	// For sync call use RestoreDiskFromBackup
	// Required: true
	async bool `url:"async"`
}

// RestoreDiskFromBackup restores disk from backup
func (b Backup) RestoreDiskFromBackup(ctx context.Context, req RestoreDiskFromBackupRequest) (ListInfoRestoredDisk, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	req.async = false

	url := "/cloudbroker/backup/restoreDiskFromBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	result := make(ListInfoRestoredDisk, 0)

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// RestoreDiskFromBackupAsync restores disk from backup
func (b Backup) RestoreDiskFromBackupAsync(ctx context.Context, req RestoreDiskFromBackupRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	req.async = true

	url := "/cloudbroker/backup/restoreDiskFromBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
