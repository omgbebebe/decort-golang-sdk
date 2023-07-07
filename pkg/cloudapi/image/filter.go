package image

// FilterByID returns ListImages with specified ID.
func (li ListImages) FilterByID(id uint64) ListImages {
	predicate := func(ii ItemImage) bool {
		return ii.ID == id
	}

	return li.FilterFunc(predicate)
}

// FilterByName returns ListImages with specified Name.
func (li ListImages) FilterByName(name string) ListImages {
	predicate := func(ii ItemImage) bool {
		return ii.Name == name
	}

	return li.FilterFunc(predicate)
}

// FilterByStatus returns ListImages with specified Status.
func (li ListImages) FilterByStatus(status string) ListImages {
	predicate := func(ii ItemImage) bool {
		return ii.Status == status
	}

	return li.FilterFunc(predicate)
}

// FilterByBootType returns ListImages with specified BootType.
func (li ListImages) FilterByBootType(bootType string) ListImages {
	predicate := func(ii ItemImage) bool {
		return ii.BootType == bootType
	}

	return li.FilterFunc(predicate)
}

// FilterFunc allows filtering ListImages based on a user-specified predicate.
func (li ListImages) FilterFunc(predicate func(ItemImage) bool) ListImages {
	var result ListImages

	for _, item := range li.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(result.Data))

	return result
}

// FindOne returns first found ItemImage
// If none was found, returns an empty struct.
func (li ListImages) FindOne() ItemImage {
	if len(li.Data) == 0 {
		return ItemImage{}
	}

	return li.Data[0]
}
