package rg

// FilterByID returns ListResourceGroups with specified ID.
func (lrg ListResourceGroups) FilterByID(id uint64) ListResourceGroups {
	predicate := func(irg ItemResourceGroup) bool {
		return irg.ID == id
	}

	return lrg.FilterFunc(predicate)
}

// FilterByName returns ListResourceGroups with specified Name.
func (lrg ListResourceGroups) FilterByName(name string) ListResourceGroups {
	predicate := func(irg ItemResourceGroup) bool {
		return irg.Name == name
	}

	return lrg.FilterFunc(predicate)
}

// FilterByCreatedBy return ListResourceGroups created by specified user.
func (lrg ListResourceGroups) FilterByCreatedBy(createdBy string) ListResourceGroups {
	predicate := func(irg ItemResourceGroup) bool {
		return irg.CreatedBy == createdBy
	}

	return lrg.FilterFunc(predicate)
}

// FilterByStatus returns ListResourceGroups with specified Status.
func (lrg ListResourceGroups) FilterByStatus(status string) ListResourceGroups {
	predicate := func(irg ItemResourceGroup) bool {
		return irg.Status == status
	}

	return lrg.FilterFunc(predicate)
}

// FilterByLockStatus return ListResourceGroups with specified LockStatus.
func (lrg ListResourceGroups) FilterByLockStatus(lockStatus string) ListResourceGroups {
	predicate := func(irg ItemResourceGroup) bool {
		return irg.LockStatus == lockStatus
	}

	return lrg.FilterFunc(predicate)
}

// FilterByDefNetType returns ListResourceGroups with specified DefNetType.
func (lrg ListResourceGroups) FilterByDefNetType(defNetType string) ListResourceGroups {
	predicate := func(irg ItemResourceGroup) bool {
		return irg.DefNetType == defNetType
	}

	return lrg.FilterFunc(predicate)
}

// FilterFunc allows filtering ListResourceGroups based on a user-specified predicate.
func (lrg ListResourceGroups) FilterFunc(predicate func(irg ItemResourceGroup) bool) ListResourceGroups {
	var result ListResourceGroups

	for _, rgItem := range lrg {
		if predicate(rgItem) {
			result = append(result, rgItem)
		}
	}

	return result
}

// FindOne returns first found ItemResourceGroup.
// If none was found, returns an empty struct.
func (lrg ListResourceGroups) FindOne() ItemResourceGroup {
	if len(lrg) == 0 {
		return ItemResourceGroup{}
	}

	return lrg[0]
}
