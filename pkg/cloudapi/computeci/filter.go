package computeci

// FilterByID returns ListComputeCI with specified ID.
func (lci ListComputeCI) FilterByID(id uint64) ListComputeCI {
	predicate := func(ic ItemComputeCI) bool {
		return ic.ID == id
	}

	return lci.FilterFunc(predicate)
}

// FilterByName returns ListComputeCI with specified Name.
func (lci ListComputeCI) FilterByName(name string) ListComputeCI {
	predicate := func(ic ItemComputeCI) bool {
		return ic.Name == name
	}

	return lci.FilterFunc(predicate)
}

// FilterByStatus returns ListComputeCI with specified Status.
func (lci ListComputeCI) FilterByStatus(status string) ListComputeCI {
	predicate := func(ic ItemComputeCI) bool {
		return ic.Status == status
	}

	return lci.FilterFunc(predicate)
}

// FilterFunc allows filtering ListComputeCI based on a user-specified predicate.
func (lci ListComputeCI) FilterFunc(predicate func(ItemComputeCI) bool) ListComputeCI {
	var result ListComputeCI

	for _, item := range lci {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemComputeCI
// If none was found, returns an empty struct.
func (lci ListComputeCI) FindOne() ItemComputeCI {
	if len(lci) == 0 {
		return ItemComputeCI{}
	}

	return lci[0]
}
