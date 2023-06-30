package locations

// FilterByID returns ListLocations with specified ID.
func (ll ListLocations) FilterByID(id uint64) ListLocations {
	predicate := func(il ItemLocation) bool {
		return il.ID == id
	}

	return ll.FilterFunc(predicate)
}

// FilterByName returns ListLocations with specified Name.
func (ll ListLocations) FilterByName(name string) ListLocations {
	predicate := func(il ItemLocation) bool {
		return il.Name == name
	}

	return ll.FilterFunc(predicate)
}

// FilterByGID returns ListLocations with specified GID.
func (ll ListLocations) FilterByGID(gid uint64) ListLocations {
	predicate := func(il ItemLocation) bool {
		return il.GID == gid
	}

	return ll.FilterFunc(predicate)
}

// FilterFunc allows filtering ListLocations based on a user-specified predicate.
func (ll ListLocations) FilterFunc(predicate func(ItemLocation) bool) ListLocations {
	var result ListLocations

	for _, item := range ll.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(result.Data))

	return result
}

// FindOne returns first found ItemLocation
// If none was found, returns an empty struct.
func (ll ListLocations) FindOne() ItemLocation {
	if len(ll.Data) == 0 {
		return ItemLocation{}
	}

	return ll.Data[0]
}
