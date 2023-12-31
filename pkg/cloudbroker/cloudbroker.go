// List of method groups for the admin
package cloudbroker

import "repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to CloudBroker groups
type CloudBroker struct {
	client interfaces.Caller
}

// Builder to get access to CloudBroker
func New(client interfaces.Caller) *CloudBroker {
	return &CloudBroker{
		client: client,
	}
}
