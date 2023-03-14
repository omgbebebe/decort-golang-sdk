package disks

import "sort"

// SortByCreatedTime sorts ListDisks by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ld ListDisks) SortByCreatedTime(inverse bool) ListDisks {
	if len(ld) < 2 {
		return ld
	}

	sort.Slice(ld, func(i, j int) bool {
		if inverse {
			return ld[i].CreatedTime > ld[j].CreatedTime
		}

		return ld[i].CreatedTime < ld[j].CreatedTime
	})

	return ld
}

// SortByDestructionTime sorts ListDisks by the DestructionTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ld ListDisks) SortByDestructionTime(inverse bool) ListDisks {
	if len(ld) < 2 {
		return ld
	}

	sort.Slice(ld, func(i, j int) bool {
		if inverse {
			return ld[i].DestructionTime > ld[j].DestructionTime
		}

		return ld[i].DestructionTime < ld[j].DestructionTime
	})

	return ld
}

// SortByDeletedTime sorts ListDisks by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ld ListDisks) SortByDeletedTime(inverse bool) ListDisks {
	if len(ld) < 2 {
		return ld
	}

	sort.Slice(ld, func(i, j int) bool {
		if inverse {
			return ld[i].DeletedTime > ld[j].DeletedTime
		}

		return ld[i].DeletedTime < ld[j].DeletedTime
	})

	return ld
}
