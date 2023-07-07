package lb

import "sort"

// SortByCreatedTime sorts ListLB by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ll ListLB) SortByCreatedTime(inverse bool) ListLB {
	if len(ll.Data) < 2 {
		return ll
	}

	sort.Slice(ll.Data, func(i, j int) bool {
		if inverse {
			return ll.Data[i].CreatedTime > ll.Data[j].CreatedTime
		}

		return ll.Data[i].CreatedTime < ll.Data[j].CreatedTime
	})

	return ll
}

// SortByUpdatedTime sorts ListLB by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ll ListLB) SortByUpdatedTime(inverse bool) ListLB {
	if len(ll.Data) < 2 {
		return ll
	}

	sort.Slice(ll.Data, func(i, j int) bool {
		if inverse {
			return ll.Data[i].UpdatedTime > ll.Data[j].UpdatedTime
		}

		return ll.Data[i].UpdatedTime < ll.Data[j].UpdatedTime
	})

	return ll
}

// SortByDeletedTime sorts ListLB by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ll ListLB) SortByDeletedTime(inverse bool) ListLB {
	if len(ll.Data) < 2 {
		return ll
	}

	sort.Slice(ll.Data, func(i, j int) bool {
		if inverse {
			return ll.Data[i].DeletedTime > ll.Data[j].DeletedTime
		}

		return ll.Data[i].DeletedTime < ll.Data[j].DeletedTime
	})

	return ll
}
