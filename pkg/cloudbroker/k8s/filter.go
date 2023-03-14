package k8s

// FilterByID returns ListK8S with specified ID.
func (lkc ListK8S) FilterByID(id uint64) ListK8S {
	predicate := func(ikc ItemK8S) bool {
		return ikc.ID == id
	}

	return lkc.FilterFunc(predicate)
}

// FilterByName returns ListK8S with specified Name.
func (lkc ListK8S) FilterByName(name string) ListK8S {
	predicate := func(ikc ItemK8S) bool {
		return ikc.Name == name
	}

	return lkc.FilterFunc(predicate)
}

// FilterByAccountID returns ListK8S with specified AccountID.
func (lkc ListK8S) FilterByAccountID(accountID uint64) ListK8S {
	predicate := func(ikc ItemK8S) bool {
		return ikc.AccountID == accountID
	}

	return lkc.FilterFunc(predicate)
}

// FilterByRGID returns ListK8S with specified RGID.
func (lkc ListK8S) FilterByRGID(rgID uint64) ListK8S {
	predicate := func(ikc ItemK8S) bool {
		return ikc.RGID == rgID
	}

	return lkc.FilterFunc(predicate)
}

// FilterByStatus returns ListK8S with specified Status.
func (lkc ListK8S) FilterByStatus(status string) ListK8S {
	predicate := func(ikc ItemK8S) bool {
		return ikc.Status == status
	}

	return lkc.FilterFunc(predicate)
}

// FilterByTechStatus returns ListK8S with specified TechStatus.
func (lkc ListK8S) FilterByTechStatus(techStatus string) ListK8S {
	predicate := func(ikc ItemK8S) bool {
		return ikc.TechStatus == techStatus
	}

	return lkc.FilterFunc(predicate)
}

// FilterByCreatedBy returns ListK8S created by specified user.
func (lkc ListK8S) FilterByCreatedBy(createdBy string) ListK8S {
	predicate := func(ikc ItemK8S) bool {
		return ikc.CreatedBy == createdBy
	}

	return lkc.FilterFunc(predicate)
}

// FilterByDeletedBy returns ListK8S deleted by specified user.
func (lkc ListK8S) FilterByDeletedBy(deletedBy string) ListK8S {
	predicate := func(ikc ItemK8S) bool {
		return ikc.DeletedBy == deletedBy
	}

	return lkc.FilterFunc(predicate)
}

// FilterFunc allows filtering ListK8S based on a user-specified predicate.
func (lkc ListK8S) FilterFunc(predicate func(ItemK8S) bool) ListK8S {
	var result ListK8S

	for _, item := range lkc {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemK8S
// If none was found, returns an empty struct.
func (lkc ListK8S) FindOne() ItemK8S {
	if len(lkc) == 0 {
		return ItemK8S{}
	}

	return lkc[0]
}
