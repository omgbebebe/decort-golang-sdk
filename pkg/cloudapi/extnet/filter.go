package extnet

// FilterByID returns ListExtNets with specified ID.
func (lenet ListExtNets) FilterByID(id uint64) ListExtNets {
	predicate := func(iexnet ItemExtNet) bool {
		return iexnet.ID == id
	}

	return lenet.FilterFunc(predicate)
}

// FilterByName returns ListExtNets with specified Name.
func (lenet ListExtNets) FilterByName(name string) ListExtNets {
	predicate := func(iexnet ItemExtNet) bool {
		return iexnet.Name == name
	}

	return lenet.FilterFunc(predicate)
}

// FilterByStatus returns ListExtNets with specified Status.
func (lenet ListExtNets) FilterByStatus(status string) ListExtNets {
	predicate := func(iexnet ItemExtNet) bool {
		return iexnet.Status == status
	}

	return lenet.FilterFunc(predicate)
}

// FilterFunc allows filtering ListExtNets based on a user-specified predicate.
func (lenet ListExtNets) FilterFunc(predicate func(ItemExtNet) bool) ListExtNets {
	var result ListExtNets

	for _, item := range lenet {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemExtNet
// If none was found, returns an empty struct.
func (lenet ListExtNets) FindOne() ItemExtNet {
	if len(lenet) == 0 {
		return ItemExtNet{}
	}

	return lenet[0]
}
