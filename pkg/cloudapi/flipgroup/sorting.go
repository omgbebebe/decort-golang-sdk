package flipgroup

import "sort"

// SortByCreatedTime sorts ListFLIPGroups by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lfg ListFLIPGroups) SortByCreatedTime(inverse bool) ListFLIPGroups {
	if len(lfg) < 2 {
		return lfg
	}

	sort.Slice(lfg, func(i, j int) bool {
		if inverse {
			return lfg[i].CreatedTime > lfg[j].CreatedTime
		}

		return lfg[i].CreatedTime < lfg[j].CreatedTime
	})

	return lfg
}

// SortByUpdatedTime sorts ListFLIPGroups by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lfg ListFLIPGroups) SortByUpdatedTime(inverse bool) ListFLIPGroups {
	if len(lfg) < 2 {
		return lfg
	}

	sort.Slice(lfg, func(i, j int) bool {
		if inverse {
			return lfg[i].UpdatedTime > lfg[j].UpdatedTime
		}

		return lfg[i].UpdatedTime < lfg[j].UpdatedTime
	})

	return lfg
}

// SortByDeletedTime sorts ListFLIPGroups by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lfg ListFLIPGroups) SortByDeletedTime(inverse bool) ListFLIPGroups {
	if len(lfg) < 2 {
		return lfg
	}

	sort.Slice(lfg, func(i, j int) bool {
		if inverse {
			return lfg[i].DeletedTime > lfg[j].DeletedTime
		}

		return lfg[i].DeletedTime < lfg[j].DeletedTime
	})

	return lfg
}
