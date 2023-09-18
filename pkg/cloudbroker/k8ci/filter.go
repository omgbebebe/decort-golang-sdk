package k8ci

// FilterByID returns ListK8CI with specified ID.
func (lkc ListK8CI) FilterByID(id uint64) ListK8CI {
	predicate := func(ikc ItemK8CI) bool {
		return ikc.RecordK8CIList.ID == id
	}

	return lkc.FilterFunc(predicate)
}

// FilterByName returns ListK8CI with specified Name.
func (lkc ListK8CI) FilterByName(name string) ListK8CI {
	predicate := func(ikc ItemK8CI) bool {
		return ikc.RecordK8CIList.Name == name
	}

	return lkc.FilterFunc(predicate)
}

// FilterByStatus returns ListK8CI with specified Status.
func (lkc ListK8CI) FilterByStatus(status string) ListK8CI {
	predicate := func(ikc ItemK8CI) bool {
		return ikc.Status == status
	}

	return lkc.FilterFunc(predicate)
}

// FilterByWorkerImageID returns ListK8CI with specified WorkerImageID.
func (lkc ListK8CI) FilterByWorkerImageID(workerImageID uint64) ListK8CI {
	predicate := func(ikc ItemK8CI) bool {
		return ikc.WorkerImageID == workerImageID
	}

	return lkc.FilterFunc(predicate)
}

// FilterByLBImageID returns ListK8CI with specified LBImageID.
func (lkc ListK8CI) FilterByLBImageID(lbImageID uint64) ListK8CI {
	predicate := func(ikc ItemK8CI) bool {
		return ikc.LBImageID == lbImageID
	}

	return lkc.FilterFunc(predicate)
}

// FilterByMasterImageID returns ListK8CI with specified MasterImageID.
func (lkc ListK8CI) FilterByMasterImageID(masterImageID uint64) ListK8CI {
	predicate := func(ikc ItemK8CI) bool {
		return ikc.MasterImageID == masterImageID
	}

	return lkc.FilterFunc(predicate)
}

// FilterFunc allows filtering ListK8CI based on a user-specified predicate.
func (lkc ListK8CI) FilterFunc(predicate func(ItemK8CI) bool) ListK8CI {
	var result ListK8CI

	for _, item := range lkc.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(result.Data))

	return result
}

// FindOne returns first found ItemK8CI
// If none was found, returns an empty struct.
func (lkc ListK8CI) FindOne() ItemK8CI {
	if len(lkc.Data) == 0 {
		return ItemK8CI{}
	}

	return lkc.Data[0]
}
