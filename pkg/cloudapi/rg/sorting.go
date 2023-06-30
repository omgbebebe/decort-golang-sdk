package rg

import "sort"

// SortByCreatedTime sorts ListResourceGroups by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListResourceGroups) SortByCreatedTime(inverse bool) ListResourceGroups {
	if len(lrg.Data) < 2 {
		return lrg
	}

	sort.Slice(lrg.Data, func(i, j int) bool {
		if inverse {
			return lrg.Data[i].CreatedTime > lrg.Data[j].CreatedTime
		}

		return lrg.Data[i].CreatedTime < lrg.Data[j].CreatedTime
	})

	return lrg
}

// SortByUpdatedTime sorts ListResourceGroups by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListResourceGroups) SortByUpdatedTime(inverse bool) ListResourceGroups {
	if len(lrg.Data) < 2 {
		return lrg
	}

	sort.Slice(lrg.Data, func(i, j int) bool {
		if inverse {
			return lrg.Data[i].UpdatedTime > lrg.Data[j].UpdatedTime
		}

		return lrg.Data[i].UpdatedTime < lrg.Data[j].UpdatedTime
	})

	return lrg
}

// SortByDeletedTime sorts ListResourceGroups by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListResourceGroups) SortByDeletedTime(inverse bool) ListResourceGroups {
	if len(lrg.Data) < 2 {
		return lrg
	}

	sort.Slice(lrg.Data, func(i, j int) bool {
		if inverse {
			return lrg.Data[i].DeletedTime > lrg.Data[j].DeletedTime
		}

		return lrg.Data[i].DeletedTime < lrg.Data[j].DeletedTime
	})

	return lrg
}
