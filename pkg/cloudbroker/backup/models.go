package backup

// Main info about backup
type InfoBackup struct {
	// Compute ID
	ComputeID uint64 `json:"computeId"`

	// Disk ID
	DiskID uint64 `json:"diskId"`

	// Backup path
	BackupPath string `json:"backupPath"`
}

// CreateDisksBackup response
type ListInfoBackup []InfoBackup

// RestoreDiskFromFile response
type InfoRestoredDisk struct {
	// Compute ID
	ComputeID uint64 `json:"computeId"`

	// Disk ID
	DiskID uint64 `json:"diskId"`
}

// RestoreDisksFromFile response
type ListInfoRestoredDisk []InfoRestoredDisk
