package grid

// FilterByID returns ListGrids with specified ID.
func (lg ListGrids) FilterByID(id uint64) ListGrids {
	predicate := func(rg RecordGrid) bool {
		return rg.ID == id
	}

	return lg.FilterFunc(predicate)
}

// FilterByName returns ListGrids with specified Name.
func (lg ListGrids) FilterByName(name string) ListGrids {
	predicate := func(rg RecordGrid) bool {
		return rg.Name == name
	}

	return lg.FilterFunc(predicate)
}

// FilterByLocationCode returns ListGrids with specified LocationCode.
func (lg ListGrids) FilterByLocationCode(locationCode string) ListGrids {
	predicate := func(rg RecordGrid) bool {
		return rg.LocationCode == locationCode
	}

	return lg.FilterFunc(predicate)
}

// FilterFunc allows filtering ListGrids based on a user-specified predicate.
func (lg ListGrids) FilterFunc(predicate func(RecordGrid) bool) ListGrids {
	var result ListGrids

	for _, item := range lg {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found RecordGrid.
// If none was found, returns an empty struct.
func (lg ListGrids) FindOne() RecordGrid {
	if len(lg) == 0 {
		return RecordGrid{}
	}

	return lg[0]
}
