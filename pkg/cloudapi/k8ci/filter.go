package k8ci

// FilterByID returns ListK8CI with specified ID.
func (lkc ListK8CI) FilterByID(id uint64) ListK8CI {
	predicate := func(ikc ItemK8CI) bool {
		return ikc.RecordK8CI.ID == id
	}

	return lkc.FilterFunc(predicate)
}

// FilterByName returns ListK8CI with specified Name.
func (lkc ListK8CI) FilterByName(name string) ListK8CI {
	predicate := func(ikc ItemK8CI) bool {
		return ikc.RecordK8CI.Name == name
	}

	return lkc.FilterFunc(predicate)
}

// FilterFunc allows filtering ListK8CI based on a user-specified predicate.
func (lkc ListK8CI) FilterFunc(predicate func(ItemK8CI) bool) ListK8CI {
	var result ListK8CI

	for _, item := range lkc {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemK8CI
// If none was found, returns an empty struct.
func (lkc ListK8CI) FindOne() ItemK8CI {
	if len(lkc) == 0 {
		return ItemK8CI{}
	}

	return lkc[0]
}
