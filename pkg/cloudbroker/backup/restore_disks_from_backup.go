package backup

import (
	"context"
	"encoding/json"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strings"
)

type BackupFile struct {
	// Disk ID
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// Backup path
	BackupPath string `url:"backupPath" json:"backupPath" validate:"required"`

	// Backup file
	BackupFile string `url:"backupFile" json:"backupFile" validate:"required"`
}

// Request struct for restoring disks from backup
type RestoreDisksFromBackupRequest struct {
	// Compute ID
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	//Backup files
	BackupFiles []BackupFile `url:"backupFiles" json:"backupFiles" validate:"required,dive"`

	// Async API Call
	// For async call use RestoreDisksFromBackupAsync
	// For sync call use RestoreDisksFromBackup
	// Required: true
	async bool `url:"async"`
}

// RestoreDisksFromBackup restores disks from backup
func (b Backup) RestoreDisksFromBackup(ctx context.Context, req RestoreDisksFromBackupRequest) (ListInfoRestoredDisk, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	req.async = false

	url := "/cloudbroker/backup/restoreDisksFromBackup"

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

// RestoreDisksFromBackupAsync restores disks from backup
func (b Backup) RestoreDisksFromBackupAsync(ctx context.Context, req RestoreDisksFromBackupRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	req.async = true

	url := "/cloudbroker/backup/restoreDisksFromBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
