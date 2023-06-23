package backup

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for deleting disk backup
type DeleteDiskBackupRequest struct {
	// Backup path
	BackupPath string `url:"backupPath" json:"backupPath" validate:"required"`

	// Backup file
	BackupFile string `url:"backupFile" json:"backupFile" validate:"required"`
}

type wrapperDeleteDiskBackupRequest struct {
	DeleteDiskBackupRequest

	Async bool `url:"async"`
}

// DeleteDiskBackup deletes disk backup
func (b Backup) DeleteDiskBackup(ctx context.Context, req DeleteDiskBackupRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	reqWrapped := wrapperDeleteDiskBackupRequest{
		DeleteDiskBackupRequest: req,
		Async:                   false,
	}

	url := "/cloudbroker/backup/deleteDiskBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
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

	reqWrapped := wrapperDeleteDiskBackupRequest{
		DeleteDiskBackupRequest: req,
		Async:                   true,
	}

	url := "/cloudbroker/backup/deleteDiskBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
