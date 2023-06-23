package user

import "repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to User
type User struct {
	client interfaces.Caller
}

// Builder for User endpoints
func New(client interfaces.Caller) *User {
	return &User{
		client: client,
	}
}
