package cloudapi

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/stack"
)

// Accessing the Stack method group
func (ca *CloudAPI) Stack() *stack.Stack {
	return stack.New(ca.client)
}
