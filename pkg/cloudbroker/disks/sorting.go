package disks

import "sort"

// SortByCreatedTime sorts ListDisks by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ld ListDisks) SortByCreatedTime(inverse bool) ListDisks {
	if len(ld.Data) < 2 {
		return ld
	}

	sort.Slice(ld.Data, func(i, j int) bool {
		if inverse {
			return ld.Data[i].CreatedTime > ld.Data[j].CreatedTime
		}

		return ld.Data[i].CreatedTime < ld.Data[j].CreatedTime
	})

	return ld
}

// SortByDestructionTime sorts ListDisks by the DestructionTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ld ListDisks) SortByDestructionTime(inverse bool) ListDisks {
	if len(ld.Data) < 2 {
		return ld
	}

	sort.Slice(ld.Data, func(i, j int) bool {
		if inverse {
			return ld.Data[i].DestructionTime > ld.Data[j].DestructionTime
		}

		return ld.Data[i].DestructionTime < ld.Data[j].DestructionTime
	})

	return ld
}

// SortByDeletedTime sorts ListDisks by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (ld ListDisks) SortByDeletedTime(inverse bool) ListDisks {
	if len(ld.Data) < 2 {
		return ld
	}

	sort.Slice(ld.Data, func(i, j int) bool {
		if inverse {
			return ld.Data[i].DeletedTime > ld.Data[j].DeletedTime
		}

		return ld.Data[i].DeletedTime < ld.Data[j].DeletedTime
	})

	return ld
}
