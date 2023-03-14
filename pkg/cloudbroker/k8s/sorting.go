package k8s

import "sort"

// SortByCreatedTime sorts ListK8S by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lkc ListK8S) SortByCreatedTime(inverse bool) ListK8S {
	if len(lkc) < 2 {
		return lkc
	}

	sort.Slice(lkc, func(i, j int) bool {
		if inverse {
			return lkc[i].CreatedTime > lkc[j].CreatedTime
		}

		return lkc[i].CreatedTime < lkc[j].CreatedTime
	})

	return lkc
}

// SortByUpdatedTime sorts ListK8S by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lkc ListK8S) SortByUpdatedTime(inverse bool) ListK8S {
	if len(lkc) < 2 {
		return lkc
	}

	sort.Slice(lkc, func(i, j int) bool {
		if inverse {
			return lkc[i].UpdatedTime > lkc[j].UpdatedTime
		}

		return lkc[i].UpdatedTime < lkc[j].UpdatedTime
	})

	return lkc
}

// SortByDeletedTime sorts ListK8S by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lkc ListK8S) SortByDeletedTime(inverse bool) ListK8S {
	if len(lkc) < 2 {
		return lkc
	}

	sort.Slice(lkc, func(i, j int) bool {
		if inverse {
			return lkc[i].DeletedTime > lkc[j].DeletedTime
		}

		return lkc[i].DeletedTime < lkc[j].DeletedTime
	})

	return lkc
}
