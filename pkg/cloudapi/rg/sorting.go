package rg

import "sort"

// SortByCreatedTime sorts ListResourceGroups by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListResourceGroups) SortByCreatedTime(inverse bool) ListResourceGroups {
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

// SortByUpdatedTime sorts ListResourceGroups by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListResourceGroups) SortByUpdatedTime(inverse bool) ListResourceGroups {
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

// SortByDeletedTime sorts ListResourceGroups by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lrg ListResourceGroups) SortByDeletedTime(inverse bool) ListResourceGroups {
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
