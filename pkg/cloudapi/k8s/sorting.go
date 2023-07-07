package k8s

import "sort"

// SortByCreatedTime sorts ListK8SClusters by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lkc ListK8SClusters) SortByCreatedTime(inverse bool) ListK8SClusters {
	if len(lkc.Data) < 2 {
		return lkc
	}

	sort.Slice(lkc.Data, func(i, j int) bool {
		if inverse {
			return lkc.Data[i].CreatedTime > lkc.Data[j].CreatedTime
		}

		return lkc.Data[i].CreatedTime < lkc.Data[j].CreatedTime
	})

	return lkc
}

// SortByUpdatedTime sorts ListK8SClusters by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lkc ListK8SClusters) SortByUpdatedTime(inverse bool) ListK8SClusters {
	if len(lkc.Data) < 2 {
		return lkc
	}

	sort.Slice(lkc.Data, func(i, j int) bool {
		if inverse {
			return lkc.Data[i].UpdatedTime > lkc.Data[j].UpdatedTime
		}

		return lkc.Data[i].UpdatedTime < lkc.Data[j].UpdatedTime
	})

	return lkc
}

// SortByDeletedTime sorts ListK8SClusters by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lkc ListK8SClusters) SortByDeletedTime(inverse bool) ListK8SClusters {
	if len(lkc.Data) < 2 {
		return lkc
	}

	sort.Slice(lkc.Data, func(i, j int) bool {
		if inverse {
			return lkc.Data[i].DeletedTime > lkc.Data[j].DeletedTime
		}

		return lkc.Data[i].DeletedTime < lkc.Data[j].DeletedTime
	})

	return lkc
}
