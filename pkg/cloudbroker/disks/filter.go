package disks

// FilterByID returns ListDisks with specified ID.
func (ld ListDisks) FilterByID(id uint64) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		return idisk.ID == id
	}

	return ld.FilterFunc(predicate)
}

// FilterByName returns ListDisks with specified Name.
func (ld ListDisks) FilterByName(name string) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		return idisk.Name == name
	}

	return ld.FilterFunc(predicate)
}

// FilterByStatus returns ListDisks with specified Status.
func (ld ListDisks) FilterByStatus(status string) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		return idisk.Status == status
	}

	return ld.FilterFunc(predicate)
}

// FilterByTechStatus returns ListDisks with specified TechStatus.
func (ld ListDisks) FilterByTechStatus(techStatus string) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		return idisk.TechStatus == techStatus
	}

	return ld.FilterFunc(predicate)
}

// FilterByImageID returns ListDisks with specified ImageID.
func (ld ListDisks) FilterByImageID(imageID uint64) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		return idisk.ImageID == imageID
	}

	return ld.FilterFunc(predicate)
}

// FilterFunc allows filtering ListDisks based on a user-specified predicate.
func (ld ListDisks) FilterFunc(predicate func(ItemDisk) bool) ListDisks {
	var result ListDisks

	for _, item := range ld {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemDisk
// If none was found, returns an empty struct.
func (ld ListDisks) FindOne() ItemDisk {
	if len(ld) == 0 {
		return ItemDisk{}
	}

	return ld[0]
}
