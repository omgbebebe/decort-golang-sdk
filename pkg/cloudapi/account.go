package cloudapi

import (
	"repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudapi/account"
)

// Accessing the Account method group
func (ca *CloudAPI) Account() *account.Account {
	return account.New(ca.client)
}
