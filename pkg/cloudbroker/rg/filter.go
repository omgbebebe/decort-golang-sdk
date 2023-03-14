package rg

// FilterByID returns ListRG with specified ID.
func (lrg ListRG) FilterByID(id uint64) ListRG {
	predicate := func(irg ItemRG) bool {
		return irg.ID == id
	}

	return lrg.FilterFunc(predicate)
}

// FilterByName returns ListRG with specified Name.
func (lrg ListRG) FilterByName(name string) ListRG {
	predicate := func(irg ItemRG) bool {
		return irg.Name == name
	}

	return lrg.FilterFunc(predicate)
}

// FilterByCreatedBy return ListRG created by specified user.
func (lrg ListRG) FilterByCreatedBy(createdBy string) ListRG {
	predicate := func(irg ItemRG) bool {
		return irg.CreatedBy == createdBy
	}

	return lrg.FilterFunc(predicate)
}

// FilterByStatus returns ListRG with specified Status.
func (lrg ListRG) FilterByStatus(status string) ListRG {
	predicate := func(irg ItemRG) bool {
		return irg.Status == status
	}

	return lrg.FilterFunc(predicate)
}

// FilterByLockStatus return ListRG with specified LockStatus.
func (lrg ListRG) FilterByLockStatus(lockStatus string) ListRG {
	predicate := func(irg ItemRG) bool {
		return irg.LockStatus == lockStatus
	}

	return lrg.FilterFunc(predicate)
}

// FilterByDefNetType returns ListRG with specified DefNetType.
func (lrg ListRG) FilterByDefNetType(defNetType string) ListRG {
	predicate := func(irg ItemRG) bool {
		return irg.DefNetType == defNetType
	}

	return lrg.FilterFunc(predicate)
}

// FilterByDefNetID returns ListRG with specified DefNetID.
func (lrg ListRG) FilterByDefNetID(defNetID int64) ListRG {
	predicate := func(irg ItemRG) bool {
		return irg.DefNetID == defNetID
	}

	return lrg.FilterFunc(predicate)
}

// FilterFunc allows filtering ListRG based on a user-specified predicate.
func (lrg ListRG) FilterFunc(predicate func(ItemRG) bool) ListRG {
	var result ListRG

	for _, item := range lrg {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemRG.
// If none was found, returns an empty struct.
func (lrg ListRG) FindOne() ItemRG {
	if len(lrg) == 0 {
		return ItemRG{}
	}

	return lrg[0]
}
