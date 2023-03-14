package lb

// FilterByID returns ListLB with specified ID.
func (ll ListLB) FilterByID(id uint64) ListLB {
	predicate := func(ill ItemLoadBalancer) bool {
		return ill.ID == id
	}

	return ll.FilterFunc(predicate)
}

// FilterByName returns ListLB with specified Name.
func (ll ListLB) FilterByName(name string) ListLB {
	predicate := func(ill ItemLoadBalancer) bool {
		return ill.Name == name
	}

	return ll.FilterFunc(predicate)
}

// FilterByExtNetID returns ListLB with specified ExtNetID.
func (ll ListLB) FilterByExtNetID(extNetID uint64) ListLB {
	predicate := func(ill ItemLoadBalancer) bool {
		return ill.ExtNetID == extNetID
	}

	return ll.FilterFunc(predicate)
}

// FilterByImageID returns ListLB with specified ImageID.
func (ll ListLB) FilterByImageID(imageID uint64) ListLB {
	predicate := func(ill ItemLoadBalancer) bool {
		return ill.ImageID == imageID
	}

	return ll.FilterFunc(predicate)
}

// FilterFunc allows filtering ListLB based on a user-specified predicate.
func (ll ListLB) FilterFunc(predicate func(ItemLoadBalancer) bool) ListLB {
	var result ListLB

	for _, item := range ll {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found ItemLoadBalancer
// If none was found, returns an empty struct.
func (ll ListLB) FindOne() ItemLoadBalancer {
	if len(ll) == 0 {
		return ItemLoadBalancer{}
	}

	return ll[0]
}
