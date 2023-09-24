// Lists all the stack.
package stack

import "repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to stack
type Stack struct {
	client interfaces.Caller
}

// Builder for stack endpoint
func New(client interfaces.Caller) *Stack {
	return &Stack{
		client: client,
	}
}
