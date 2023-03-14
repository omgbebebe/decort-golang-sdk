package bservice

// FilterByID returns ListBasicServices with specified ID.
func (lbs ListBasicServices) FilterByID(id uint64) ListBasicServices {
	predicate := func(ibs ItemBasicService) bool {
		return ibs.ID == id
	}

	return lbs.FilterFunc(predicate)
}

// FilterByName returns ListBasicServices with specified Name.
func (lbs ListBasicServices) FilterByName(name string) ListBasicServices {
	predicate := func(ibs ItemBasicService) bool {
		return ibs.Name == name
	}

	return lbs.FilterFunc(predicate)
}

// FilterByRGID returns ListBasicServices with specified RGID.
func (lbs ListBasicServices) FilterByRGID(rgID uint64) ListBasicServices {
	predicate := func(ibs ItemBasicService) bool {
		return ibs.RGID == rgID
	}

	return lbs.FilterFunc(predicate)
}

// FilterByStatus returns ListBasicServices with specified Status.
func (lbs ListBasicServices) FilterByStatus(status string) ListBasicServices {
	predicate := func(ibs ItemBasicService) bool {
		return ibs.Status == status
	}

	return lbs.FilterFunc(predicate)
}

// FilterByTechStatus returns ListBasicServices with specified TechStatus.
func (lbs ListBasicServices) FilterByTechStatus(techStatus string) ListBasicServices {
	predicate := func(ibs ItemBasicService) bool {
		return ibs.TechStatus == techStatus
	}

	return lbs.FilterFunc(predicate)
}

// FilterFunc allows filtering ListResourceGroups based on a user-specified predicate.
func (lbs ListBasicServices) FilterFunc(predicate func(ItemBasicService) bool) ListBasicServices {
	var result ListBasicServices

	for _, item := range lbs {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemBasicService
// If none was found, returns an empty struct.
func (lbs ListBasicServices) FindOne() ItemBasicService {
	if len(lbs) == 0 {
		return ItemBasicService{}
	}

	return lbs[0]
}
