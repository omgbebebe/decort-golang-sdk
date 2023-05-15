package cloudbroker

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/backup"

// Accessing the Backup method group
func (cb *CloudBroker) Backup() *backup.Backup {
	return backup.New(cb.client)
}
