package backup

import (
	"context"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strconv"
	"strings"
)

// Request struct for deleting disk backup
type DeleteDiskBackupRequest struct {
	// Backup path
	BackupPath string `url:"backupPath" json:"backupPath" validate:"required"`

	// Backup file
	BackupFile string `url:"backupFile" json:"backupFile" validate:"required"`

	// Async API Call
	// For async call use DeleteDiskBackupAsync
	// For sync call use DeleteDiskBackup
	// Required: true
	async bool `url:"async"`
}

// DeleteDiskBackup deletes disk backup
func (b Backup) DeleteDiskBackup(ctx context.Context, req DeleteDiskBackupRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	req.async = false

	url := "/cloudbroker/backup/deleteDiskBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}

// DeleteDiskBackupAsync deletes disk backup
func (b Backup) DeleteDiskBackupAsync(ctx context.Context, req DeleteDiskBackupRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	req.async = true

	url := "/cloudbroker/backup/deleteDiskBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
