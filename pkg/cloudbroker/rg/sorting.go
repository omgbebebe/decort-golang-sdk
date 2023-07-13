package rg

import "sort"

// SortByCreatedTime sorts ListRG by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListRG) SortByCreatedTime(inverse bool) ListRG {
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

// SortByUpdatedTime sorts ListRG by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListRG) SortByUpdatedTime(inverse bool) ListRG {
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

// SortByDeletedTime sorts ListRG by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListRG) SortByDeletedTime(inverse bool) ListRG {
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
