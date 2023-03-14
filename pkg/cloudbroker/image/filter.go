package image

// FilterById returns ListImages with specified ID.
func (li ListImages) FilterById(id uint64) ListImages {
	predicate := func(ri RecordImage) bool {
		return ri.ID == id
	}

	return li.FilterFunc(predicate)
}

// FilterByName returns ListImages with specified Name.
func (li ListImages) FilterByName(name string) ListImages {
	predicate := func(ri RecordImage) bool {
		return ri.Name == name
	}

	return li.FilterFunc(predicate)
}

// FilterByStatus returns ListImages with specified Status.
func (li ListImages) FilterByStatus(status string) ListImages {
	predicate := func(ri RecordImage) bool {
		return ri.Status == status
	}

	return li.FilterFunc(predicate)
}

// FilterByTechStatus returns ListImages with specified TechStatus.
func (li ListImages) FilterByTechStatus(techStatus string) ListImages {
	predicate := func(ri RecordImage) bool {
		return ri.TechStatus == techStatus
	}

	return li.FilterFunc(predicate)
}

// FilterByBootType returns ListImages with specified BootType.
func (li ListImages) FilterByBootType(bootType string) ListImages {
	predicate := func(ri RecordImage) bool {
		return ri.BootType == bootType
	}

	return li.FilterFunc(predicate)
}

// FilterFunc allows filtering ListImages based on a user-specified predicate.
func (li ListImages) FilterFunc(predicate func(RecordImage) bool) ListImages {
	var result ListImages

	for _, item := range li {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found RecordImage
// If none was found, returns an empty struct.
func (li ListImages) FindOne() RecordImage {
	if len(li) == 0 {
		return RecordImage{}
	}

	return li[0]
}
