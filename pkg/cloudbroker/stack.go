package cloudbroker

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/stack"
)

// Accessing the Stack method group
func (cb *CloudBroker) Stack() *stack.Stack {
	return stack.New(cb.client)
}
