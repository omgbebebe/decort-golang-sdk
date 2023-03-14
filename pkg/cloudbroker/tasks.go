package cloudbroker

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudbroker/tasks"
)

// Accessing the tasks method group
func (cb *CloudBroker) Tasks() *tasks.Tasks {
	return tasks.New(cb.client)
}
