package grid

// FilterByID returns ListGrids with specified ID.
func (lg ListGrids) FilterByID(id uint64) ListGrids {
	predicate := func(rg ItemGridList) bool {
		return rg.ID == id
	}

	return lg.FilterFunc(predicate)
}

// FilterByGID returns ListGrids with specified GID.
func (lg ListGrids) FilterByGID(gid uint64) ListGrids {
	predicate := func(rg ItemGridList) bool {
		return rg.GID == gid
	}

	return lg.FilterFunc(predicate)
}

// FilterByName returns ListGrids with specified Name.
func (lg ListGrids) FilterByName(name string) ListGrids {
	predicate := func(rg ItemGridList) bool {
		return rg.Name == name
	}

	return lg.FilterFunc(predicate)
}

// FilterByLocationCode returns ListGrids with specified LocationCode.
func (lg ListGrids) FilterByLocationCode(locationCode string) ListGrids {
	predicate := func(rg ItemGridList) bool {
		return rg.LocationCode == locationCode
	}

	return lg.FilterFunc(predicate)
}

// FilterFunc allows filtering ListGrids based on a user-specified predicate.
func (lg ListGrids) FilterFunc(predicate func(ItemGridList) bool) ListGrids {
	var result ListGrids

	for _, item := range lg.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(result.Data))

	return result
}

// FindOne returns first found RecordGrid.
// If none was found, returns an empty struct.
func (lg ListGrids) FindOne() ItemGridList {
	if len(lg.Data) == 0 {
		return ItemGridList{}
	}

	return lg.Data[0]
}
