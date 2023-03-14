package k8ci

import "sort"

// SortByCreatedTime sorts ListK8CI by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lkc ListK8CI) SortByCreatedTime(inverse bool) ListK8CI {
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
