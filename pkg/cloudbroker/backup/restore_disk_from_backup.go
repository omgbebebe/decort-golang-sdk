package backup

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
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
}

type wrapperRestoreDiskFromBackupRequest struct {
	RestoreDiskFromBackupRequest

	Async bool `url:"async"`
}

// RestoreDiskFromBackup restores disk from backup
func (b Backup) RestoreDiskFromBackup(ctx context.Context, req RestoreDiskFromBackupRequest) (ListInfoRestoredDisk, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	reqWrapped := wrapperRestoreDiskFromBackupRequest{
		RestoreDiskFromBackupRequest: req,
		Async:                        false,
	}

	url := "/cloudbroker/backup/restoreDiskFromBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
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

	reqWrapped := wrapperRestoreDiskFromBackupRequest{
		RestoreDiskFromBackupRequest: req,
		Async:                        true,
	}

	url := "/cloudbroker/backup/restoreDiskFromBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
