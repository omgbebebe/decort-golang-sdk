package backup

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
)

// Structure for creating request to backup
type Backup struct {
	client interfaces.Caller
}

// Builder for backup endpoints
func New(client interfaces.Caller) *Backup {
	return &Backup{
		client: client,
	}
}
