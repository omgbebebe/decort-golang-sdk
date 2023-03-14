package compute

// FilterByID returns ListComputes with specified ID.
func (lc ListComputes) FilterByID(id uint64) ListComputes {
	predicate := func(ic ItemCompute) bool {
		return ic.ID == id
	}

	return lc.FilterFunc(predicate)
}

// FilterByName returns ListComputes with specified Name.
func (lc ListComputes) FilterByName(name string) ListComputes {
	predicate := func(ic ItemCompute) bool {
		return ic.Name == name
	}

	return lc.FilterFunc(predicate)
}

// FilterByStatus returns ListComputes with specified Status.
func (lc ListComputes) FilterByStatus(status string) ListComputes {
	predicate := func(ic ItemCompute) bool {
		return ic.Status == status
	}

	return lc.FilterFunc(predicate)
}

// FilterByTechStatus returns ListComputes with specified TechStatus.
func (lc ListComputes) FilterByTechStatus(techStatus string) ListComputes {
	predicate := func(ic ItemCompute) bool {
		return ic.TechStatus == techStatus
	}

	return lc.FilterFunc(predicate)
}

// FilterByDiskID return ListComputes with specified DiskID.
func (lc ListComputes) FilterByDiskID(diskID uint64) ListComputes {
	predicate := func(ic ItemCompute) bool {
		for _, disk := range ic.Disks {
			if disk.ID == diskID {
				return true
			}
		}
		return false
	}

	return lc.FilterFunc(predicate)
}

// FilterFunc allows filtering ListComputes based on a user-specified predicate.
func (lc ListComputes) FilterFunc(predicate func(ItemCompute) bool) ListComputes {
	var result ListComputes

	for _, item := range lc {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemCompute
// If none was found, returns an empty struct.
func (lc ListComputes) FindOne() ItemCompute {
	if len(lc) == 0 {
		return ItemCompute{}
	}

	return lc[0]
}
