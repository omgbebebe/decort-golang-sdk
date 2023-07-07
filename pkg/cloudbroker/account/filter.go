package account

// FilterByID returns ListAccounts with specified ID.
func (la ListAccounts) FilterByID(id uint64) ListAccounts {
	predicate := func(ia ItemAccount) bool {
		return ia.ID == id
	}

	return la.FilterFunc(predicate)
}

// FilterByName returns ListAccounts with specified Name.
func (la ListAccounts) FilterByName(name string) ListAccounts {
	predicate := func(ia ItemAccount) bool {
		return ia.Name == name
	}

	return la.FilterFunc(predicate)
}

// FilterByStatus returns ListAccounts with specified Status.
func (la ListAccounts) FilterByStatus(status string) ListAccounts {
	predicate := func(ia ItemAccount) bool {
		return ia.Status == status
	}

	return la.FilterFunc(predicate)
}

// FilterByUserGroupID returns ListAccounts with specified UserGroupID.
func (la ListAccounts) FilterByUserGroupID(userGroupID string) ListAccounts {
	predicate := func(ia ItemAccount) bool {
		acl := ia.ACL

		for _, item := range acl {
			if item.UserGroupID == userGroupID {
				return true
			}
		}

		return false
	}

	return la.FilterFunc(predicate)
}

// FilterByCompany returns ListAccounts with specified Company.
func (la ListAccounts) FilterByCompany(company string) ListAccounts {
	predicate := func(ia ItemAccount) bool {
		return ia.Company == company
	}

	return la.FilterFunc(predicate)
}

// FilterByCreatedBy returns ListAccounts created by specified user.
func (la ListAccounts) FilterByCreatedBy(createdBy string) ListAccounts {
	predicate := func(ia ItemAccount) bool {
		return ia.CreatedBy == createdBy
	}

	return la.FilterFunc(predicate)
}

// FilterFunc allows filtering ListAccounts based on a user-specified predicate.
func (la ListAccounts) FilterFunc(predicate func(ItemAccount) bool) ListAccounts {
	var result ListAccounts

	for _, item := range la.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(result.Data))

	return result
}

// FindOne returns first found ItemAccount.
// If none was found, returns an empty struct.
func (la ListAccounts) FindOne() ItemAccount {
	if len(la.Data) == 0 {
		return ItemAccount{}
	}

	return la.Data[0]
}
