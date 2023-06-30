package k8ci

import "sort"

// SortByCreatedTime sorts ListK8CI by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lkc ListK8CI) SortByCreatedTime(inverse bool) ListK8CI {
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
