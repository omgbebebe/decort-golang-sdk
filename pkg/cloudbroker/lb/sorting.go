package lb

import "sort"

// SortByCreatedTime sorts ListLB by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ll ListLB) SortByCreatedTime(inverse bool) ListLB {
	if len(ll) < 2 {
		return ll
	}

	sort.Slice(ll, func(i, j int) bool {
		if inverse {
			return ll[i].CreatedTime > ll[j].CreatedTime
		}

		return ll[i].CreatedTime < ll[j].CreatedTime
	})

	return ll
}

// SortByUpdatedTime sorts ListLB by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ll ListLB) SortByUpdatedTime(inverse bool) ListLB {
	if len(ll) < 2 {
		return ll
	}

	sort.Slice(ll, func(i, j int) bool {
		if inverse {
			return ll[i].UpdatedTime > ll[j].UpdatedTime
		}

		return ll[i].UpdatedTime < ll[j].UpdatedTime
	})

	return ll
}

// SortByDeletedTime sorts ListLB by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ll ListLB) SortByDeletedTime(inverse bool) ListLB {
	if len(ll) < 2 {
		return ll
	}

	sort.Slice(ll, func(i, j int) bool {
		if inverse {
			return ll[i].DeletedTime > ll[j].DeletedTime
		}

		return ll[i].DeletedTime < ll[j].DeletedTime
	})

	return ll
}
