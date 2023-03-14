package cloudbroker

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudbroker/account"
)

// Accessing the Account method group
func (cb *CloudBroker) Account() *account.Account {
	return account.New(cb.client)
}
