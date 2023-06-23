package cloudbroker

import "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/user"

func (cb *CloudBroker) User() *user.User {
	return user.New(cb.client)
}
