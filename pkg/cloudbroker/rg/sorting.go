package rg

import "sort"

// SortByCreatedTime sorts ListRG by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListRG) SortByCreatedTime(inverse bool) ListRG {
	if len(lrg) < 2 {
		return lrg
	}

	sort.Slice(lrg, func(i, j int) bool {
		if inverse {
			return lrg[i].CreatedTime > lrg[j].CreatedTime
		}

		return lrg[i].CreatedTime < lrg[j].CreatedTime
	})

	return lrg
}

// SortByUpdatedTime sorts ListRG by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListRG) SortByUpdatedTime(inverse bool) ListRG {
	if len(lrg) < 2 {
		return lrg
	}

	sort.Slice(lrg, func(i, j int) bool {
		if inverse {
			return lrg[i].UpdatedTime > lrg[j].UpdatedTime
		}

		return lrg[i].UpdatedTime < lrg[j].UpdatedTime
	})

	return lrg
}

// SortByDeletedTime sorts ListRG by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListRG) SortByDeletedTime(inverse bool) ListRG {
	if len(lrg) < 2 {
		return lrg
	}

	sort.Slice(lrg, func(i, j int) bool {
		if inverse {
			return lrg[i].DeletedTime > lrg[j].DeletedTime
		}

		return lrg[i].DeletedTime < lrg[j].DeletedTime
	})

	return lrg
}
