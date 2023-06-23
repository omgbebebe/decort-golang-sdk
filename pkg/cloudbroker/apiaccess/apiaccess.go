package apiaccess

import "repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to APIAccess
type APIAccess struct {
	client interfaces.Caller
}

// Builder for APIAccess endpoints
func New(client interfaces.Caller) *APIAccess {
	return &APIAccess{
		client: client,
	}
}
