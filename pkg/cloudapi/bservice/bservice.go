// API Actor for managing Compute Group. This actor is a final API for endusers to manage Compute Group
package bservice

import "repos.digitalenergy.online/BASIS/decort-golang-sdk/interfaces"

// Structure for creating request to bservice
type BService struct {
	client interfaces.Caller
}

// Builder for bservice endpoints
func New(client interfaces.Caller) *BService {
	return &BService{
		client,
	}
}
