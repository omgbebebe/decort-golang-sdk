package sep

// FilterByID returns ListSEP with specified ID.
func (lsep ListSEP) FilterByID(id uint64) ListSEP {
	predicate := func(rsep RecordSEP) bool {
		return rsep.ID == id
	}

	return lsep.FilterFunc(predicate)
}

// FilterByName returns ListSEP with specified Name.
func (lsep ListSEP) FilterByName(name string) ListSEP {
	predicate := func(rsep RecordSEP) bool {
		return rsep.Name == name
	}

	return lsep.FilterFunc(predicate)
}

// FilterByObjStatus returns ListSEP with specified ObjStatus.
func (lsep ListSEP) FilterByObjStatus(objStatus string) ListSEP {
	predicate := func(rsep RecordSEP) bool {
		return rsep.ObjStatus == objStatus
	}

	return lsep.FilterFunc(predicate)
}

// FilterByTechStatus returns ListSEP with specified TechStatus.
func (lsep ListSEP) FilterByTechStatus(techStatus string) ListSEP {
	predicate := func(rsep RecordSEP) bool {
		return rsep.TechStatus == techStatus
	}

	return lsep.FilterFunc(predicate)
}

// FilterByType returns ListSEP with specified Type.
func (lsep ListSEP) FilterByType(sepType string) ListSEP {
	predicate := func(rsep RecordSEP) bool {
		return rsep.Type == sepType
	}

	return lsep.FilterFunc(predicate)
}

// FilterFunc allows filtering ListSEP based on user-specified predicate.
func (lsep ListSEP) FilterFunc(predicate func(RecordSEP) bool) ListSEP {
	var result ListSEP

	for _, item := range lsep.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(result.Data))

	return result
}

// FindOne returns first found RecordSEP
// If none was found, returns an empty struct.
func (lsep ListSEP) FindOne() RecordSEP {
	if len(lsep.Data) == 0 {
		return RecordSEP{}
	}

	return lsep.Data[0]
}
