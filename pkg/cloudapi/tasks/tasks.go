// User API tasks interface
package tasks

import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
)

// Structure for creating request to tasks
type Tasks struct {
	client interfaces.Caller
}

// Builder for tasks endpoints
func New(client interfaces.Caller) *Tasks {
	return &Tasks{
		client,
	}
}
