// API Actor API for managing account
package account

import "repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to account
type Account struct {
	client interfaces.Caller
}

// Builder for account endpoints
func New(client interfaces.Caller) *Account {
	return &Account{
		client: client,
	}
}
