package cloudbroker

import "repos.digitalenergy.online/BASIS/decort-golang-sdk/pkg/cloudbroker/grid"

// Accessing the Grid method group
func (cb *CloudBroker) Grid() *grid.Grid {
	return grid.New(cb.client)
}
