package sizes

// FilterByID returns ListSizes with specified ID.
func (ls ListSizes) FilterByID(id uint64) ListSizes {
	predicate := func(is ItemSize) bool {
		return is.ID == id
	}

	return ls.FilterFunc(predicate)
}

// FilterByName returns ListSizes with specified Name.
func (ls ListSizes) FilterByName(name string) ListSizes {
	predicate := func(is ItemSize) bool {
		return is.Name == name
	}

	return ls.FilterFunc(predicate)
}

// FilterFunc allows filtering ListSizes based on a user-specified predicate.
func (ls ListSizes) FilterFunc(predicate func(ItemSize) bool) ListSizes {
	var result ListSizes

	for _, item := range ls.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(result.Data))

	return result
}

// FindOne returns first found ItemSize
// If none was found, returns an empty struct.
func (ls ListSizes) FindOne() ItemSize {
	if len(ls.Data) == 0 {
		return ItemSize{}
	}

	return ls.Data[0]
}
