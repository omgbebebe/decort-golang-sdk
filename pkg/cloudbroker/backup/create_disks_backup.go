package backup

import (
	"context"
	"encoding/json"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
	"strings"
)

type Disk struct {
	// Disk ID
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// Backup path
	BackupPath string `url:"backupPath" json:"backupPath" validate:"required"`
}

// Request struct for creating disks backup
type CreateDisksBackupRequest struct {
	// Compute ID
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Disks
	Disks []Disk `url:"disks" json:"disks" validate:"required,dive"`

	// Async API Call
	// For async call use CreateDisksBackupAsync
	// For sync call use CreateDisksBackup
	// Required: true
	async bool `url:"async"`
}

// CreateDisksBackup creates disks backup
func (b Backup) CreateDisksBackup(ctx context.Context, req CreateDisksBackupRequest) (ListInfoBackup, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	req.async = false

	url := "/cloudbroker/backup/createDisksBackup"

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

// CreateDisksBackupAsync creates disks backup
func (b Backup) CreateDisksBackupAsync(ctx context.Context, req CreateDisksBackupRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	req.async = true

	url := "/cloudbroker/backup/createDisksBackup"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
