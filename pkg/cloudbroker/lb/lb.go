// API to manage load balancer instance
package lb

import "repos.digitalenergy.online/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to load balancer
type LB struct {
	client interfaces.Caller
}

// Builder for load balancer
func New(client interfaces.Caller) *LB {
	return &LB{
		client: client,
	}
}
