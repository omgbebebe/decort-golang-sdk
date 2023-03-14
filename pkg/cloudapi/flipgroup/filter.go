package flipgroup

// FilterByID returns ListFLIPGroups with specified ID.
func (lfg ListFLIPGroups) FilterByID(id uint64) ListFLIPGroups {
	predicate := func(ifg ItemFLIPGroup) bool {
		return ifg.ID == id
	}

	return lfg.FilterFunc(predicate)
}

// FilterByAccountID returns ListFLIPGroups with specified AccountID.
func (lfg ListFLIPGroups) FilterByAccountID(accountID uint64) ListFLIPGroups {
	predicate := func(ifg ItemFLIPGroup) bool {
		return ifg.AccountID == accountID
	}

	return lfg.FilterFunc(predicate)
}

// FilterByRGID returns ListFLIPGroups with specified RGID.
func (lfg ListFLIPGroups) FilterByRGID(rgID uint64) ListFLIPGroups {
	predicate := func(ifg ItemFLIPGroup) bool {
		return ifg.RGID == rgID
	}

	return lfg.FilterFunc(predicate)
}

// FilterByName returns ListFLIPGroups with specified Name.
func (lfg ListFLIPGroups) FilterByName(name string) ListFLIPGroups {
	predicate := func(ifg ItemFLIPGroup) bool {
		return ifg.Name == name
	}

	return lfg.FilterFunc(predicate)
}

// FilterByStatus returns ListFLIPGroups with specified Status.
func (lfg ListFLIPGroups) FilterByStatus(status string) ListFLIPGroups {
	predicate := func(ifg ItemFLIPGroup) bool {
		return ifg.Status == status
	}

	return lfg.FilterFunc(predicate)
}

// FilterByCreatedBy returns ListFLIPGroups created by specified user.
func (lfg ListFLIPGroups) FilterByCreatedBy(createdBy string) ListFLIPGroups {
	predicate := func(ifg ItemFLIPGroup) bool {
		return ifg.CreatedBy == createdBy
	}

	return lfg.FilterFunc(predicate)
}

// FilterByUpdatedBy returns ListFLIPGroups updated by specified user.
func (lfg ListFLIPGroups) FilterByUpdatedBy(updatedBy string) ListFLIPGroups {
	predicate := func(ifg ItemFLIPGroup) bool {
		return ifg.UpdatedBy == updatedBy
	}

	return lfg.FilterFunc(predicate)
}

// FilterByDeletedBy returns ListFLIPGroups deleted by specified user.
func (lfg ListFLIPGroups) FilterByDeletedBy(deletedBy string) ListFLIPGroups {
	predicate := func(ifg ItemFLIPGroup) bool {
		return ifg.DeletedBy == deletedBy
	}

	return lfg.FilterFunc(predicate)
}

// FilterFunc allows filtering ListFLIPGroups based on a user-specified predicate.
func (lfg ListFLIPGroups) FilterFunc(predicate func(ItemFLIPGroup) bool) ListFLIPGroups {
	var result ListFLIPGroups

	for _, item := range lfg {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemFLIPGroup
// If none was found, returns an empty struct.
func (lfg ListFLIPGroups) FindOne() ItemFLIPGroup {
	if len(lfg) == 0 {
		return ItemFLIPGroup{}
	}

	return lfg[0]
}
