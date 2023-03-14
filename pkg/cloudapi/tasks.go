package cloudapi

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudapi/tasks"
)

// Accessing the Tasks method group
func (ca *CloudAPI) Tasks() *tasks.Tasks {
	return tasks.New(ca.client)
}
