package extnet

// FilterByID returns ListExtNet with specified ID.
func (lenet ListExtNet) FilterByID(id uint64) ListExtNet {
	predicate := func(iexnet ItemExtNet) bool {
		return iexnet.ID == id
	}

	return lenet.FilterFunc(predicate)
}

// FilterByName returns ListExtNet with specified Name.
func (lenet ListExtNet) FilterByName(name string) ListExtNet {
	predicate := func(iexnet ItemExtNet) bool {
		return iexnet.Name == name
	}

	return lenet.FilterFunc(predicate)
}

// FilterByStatus returns ListExtNet with specified Status.
func (lenet ListExtNet) FilterByStatus(status string) ListExtNet {
	predicate := func(iexnet ItemExtNet) bool {
		return iexnet.Status == status
	}

	return lenet.FilterFunc(predicate)
}

// FilterFunc allows filtering ListExtNet based on a user-specified predicate.
func (lenet ListExtNet) FilterFunc(predicate func(ItemExtNet) bool) ListExtNet {
	var result ListExtNet

	for _, item := range lenet {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemExtNet
// If none was found, returns an empty struct.
func (lenet ListExtNet) FindOne() ItemExtNet {
	if len(lenet) == 0 {
		return ItemExtNet{}
	}

	return lenet[0]
}
