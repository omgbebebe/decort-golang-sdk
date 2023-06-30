package account

// FilterByID returns ListDeleted with specified ID.
func (ld ListDeleted) FilterByID(id uint64) ListDeleted {
	predicate := func(ia ItemAccount) bool {
		return ia.ID == id
	}

	return ld.FilterFunc(predicate)
}

// FilterByName returns ListDeleted with specified Name.
func (ld ListDeleted) FilterByName(name string) ListDeleted {
	predicate := func(ia ItemAccount) bool {
		return ia.Name == name
	}

	return ld.FilterFunc(predicate)
}

// FilterByStatus returns ListDeleted with specified Status.
func (ld ListDeleted) FilterByStatus(status string) ListDeleted {
	predicate := func(ia ItemAccount) bool {
		return ia.Status == status
	}

	return ld.FilterFunc(predicate)
}

// FilterByUserGroupID returns ListDeleted with specified UserGroupID.
func (ld ListDeleted) FilterByUserGroupID(userGroupID string) ListDeleted {
	predicate := func(ia ItemAccount) bool {
		acl := ia.ACL

		for _, item := range acl {
			if item.UserGroupID == userGroupID {
				return true
			}
		}

		return false
	}

	return ld.FilterFunc(predicate)
}

// FilterByCompany returns ListDeleted with specified Company.
func (ld ListDeleted) FilterByCompany(company string) ListDeleted {
	predicate := func(ia ItemAccount) bool {
		return ia.Company == company
	}

	return ld.FilterFunc(predicate)
}

// FilterByCreatedBy returns ListDeleted created by specified user.
func (ld ListDeleted) FilterByCreatedBy(createdBy string) ListDeleted {
	predicate := func(ia ItemAccount) bool {
		return ia.CreatedBy == createdBy
	}

	return ld.FilterFunc(predicate)
}

// FilterFunc allows filtering ListDeleted based on a user-specified predicate.
func (ld ListDeleted) FilterFunc(predicate func(ItemAccount) bool) ListDeleted {
	var result ListDeleted

	for _, item := range ld {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemAccount.
// If none was found, returns an empty struct.
func (ld ListDeleted) FindOne() ItemAccount {
	if len(ld) == 0 {
		return ItemAccount{}
	}

	return ld[0]
}
