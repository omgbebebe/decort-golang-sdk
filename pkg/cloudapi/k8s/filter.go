package k8s

// FilterByID returns ListK8SClusters with specified ID.
func (lkc ListK8SClusters) FilterByID(id uint64) ListK8SClusters {
	predicate := func(ikc ItemK8SCluster) bool {
		return ikc.ID == id
	}

	return lkc.FilterFunc(predicate)
}

// FilterByName returns ListK8SClusters with specified Name.
func (lkc ListK8SClusters) FilterByName(name string) ListK8SClusters {
	predicate := func(ikc ItemK8SCluster) bool {
		return ikc.Name == name
	}

	return lkc.FilterFunc(predicate)
}

// FilterByAccountID returns ListK8SClusters with specified AccountID.
func (lkc ListK8SClusters) FilterByAccountID(accountID uint64) ListK8SClusters {
	predicate := func(ikc ItemK8SCluster) bool {
		return ikc.AccountID == accountID
	}

	return lkc.FilterFunc(predicate)
}

// FilterByRGID returns ListK8SClusters with specified RGID.
func (lkc ListK8SClusters) FilterByRGID(rgID uint64) ListK8SClusters {
	predicate := func(ikc ItemK8SCluster) bool {
		return ikc.RGID == rgID
	}

	return lkc.FilterFunc(predicate)
}

// FilterByStatus returns ListK8SClusters with specified Status.
func (lkc ListK8SClusters) FilterByStatus(status string) ListK8SClusters {
	predicate := func(ikc ItemK8SCluster) bool {
		return ikc.Status == status
	}

	return lkc.FilterFunc(predicate)
}

// FilterByTechStatus returns ListK8SClusters with specified TechStatus.
func (lkc ListK8SClusters) FilterByTechStatus(techStatus string) ListK8SClusters {
	predicate := func(ikc ItemK8SCluster) bool {
		return ikc.TechStatus == techStatus
	}

	return lkc.FilterFunc(predicate)
}

// FilterByCreatedBy returns ListK8SClusters created by specified user.
func (lkc ListK8SClusters) FilterByCreatedBy(createdBy string) ListK8SClusters {
	predicate := func(ikc ItemK8SCluster) bool {
		return ikc.CreatedBy == createdBy
	}

	return lkc.FilterFunc(predicate)
}

// FilterByDeletedBy returns ListK8SClusters deleted by specified user.
func (lkc ListK8SClusters) FilterByDeletedBy(deletedBy string) ListK8SClusters {
	predicate := func(ikc ItemK8SCluster) bool {
		return ikc.DeletedBy == deletedBy
	}

	return lkc.FilterFunc(predicate)
}

// FilterFunc allows filtering ListK8SClusters based on a user-specified predicate.
func (lkc ListK8SClusters) FilterFunc(predicate func(ItemK8SCluster) bool) ListK8SClusters {
	var result ListK8SClusters

	for _, item := range lkc.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(result.Data))

	return result
}

// FindOne returns first found ItemK8SCluster
// If none was found, returns an empty struct.
func (lkc ListK8SClusters) FindOne() ItemK8SCluster {
	if len(lkc.Data) == 0 {
		return ItemK8SCluster{}
	}

	return lkc.Data[0]
}
