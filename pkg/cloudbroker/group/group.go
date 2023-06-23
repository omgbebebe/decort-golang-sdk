package group

import "repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to group
type Group struct {
	client interfaces.Caller
}

// Builder for group endpoints
func New(client interfaces.Caller) *Group {
	return &Group{
		client: client,
	}
}
