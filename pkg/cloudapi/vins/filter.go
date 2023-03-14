package vins

// FilterByID returns ListVINS with specified ID.
func (lv ListVINS) FilterByID(id uint64) ListVINS {
	predicate := func(iv ItemVINS) bool {
		return iv.ID == id
	}

	return lv.FilterFunc(predicate)
}

// FilterByName returns ListVINS with specified Name.
func (lv ListVINS) FilterByName(name string) ListVINS {
	predicate := func(iv ItemVINS) bool {
		return iv.Name == name
	}

	return lv.FilterFunc(predicate)
}

// FilterByAccountID returns ListVINS with specified AccountID.
func (lv ListVINS) FilterByAccountID(accountID uint64) ListVINS {
	predicate := func(iv ItemVINS) bool {
		return iv.AccountID == accountID
	}

	return lv.FilterFunc(predicate)
}

// FilterByCreatedBy returns ListVINS created by specified user.
func (lv ListVINS) FilterByCreatedBy(createdBy string) ListVINS {
	predicate := func(iv ItemVINS) bool {
		return iv.CreatedBy == createdBy
	}

	return lv.FilterFunc(predicate)
}

// FilterByUpdatedBy returns ListVINS updated by specified user.
func (lv ListVINS) FilterByUpdatedBy(updatedBy string) ListVINS {
	predicate := func(iv ItemVINS) bool {
		return iv.UpdatedBy == updatedBy
	}

	return lv.FilterFunc(predicate)
}

// FilterByDeletedBy returns ListVINS deleted by specified user.
func (lv ListVINS) FilterByDeletedBy(deletedBy string) ListVINS {
	predicate := func(iv ItemVINS) bool {
		return iv.DeletedBy == deletedBy
	}

	return lv.FilterFunc(predicate)
}

// FilterFunc allows filtering ListVINS based on a user-specified predicate.
func (lv ListVINS) FilterFunc(predicate func(ItemVINS) bool) ListVINS {
	var result ListVINS

	for _, item := range lv {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemVINS
// If none was found, returns an empty struct.
func (lv ListVINS) FindOne() ItemVINS {
	if len(lv) == 0 {
		return ItemVINS{}
	}

	return lv[0]
}
