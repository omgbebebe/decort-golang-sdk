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

// SortByCreatedTime sorts ListDisksUnattached by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lu ListDisksUnattached) SortByCreatedTime(inverse bool) ListDisksUnattached {
	if len(lu) < 2 {
		return lu
	}

	sort.Slice(lu, func(i, j int) bool {
		if inverse {
			return lu[i].CreatedTime > lu[j].CreatedTime
		}

		return lu[i].CreatedTime < lu[j].CreatedTime
	})

	return lu
}

// SortByDestructionTime sorts ListDisksUnattached by the DestructionTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lu ListDisksUnattached) SortByDestructionTime(inverse bool) ListDisksUnattached {
	if len(lu) < 2 {
		return lu
	}

	sort.Slice(lu, func(i, j int) bool {
		if inverse {
			return lu[i].DestructionTime > lu[j].DestructionTime
		}

		return lu[i].DestructionTime < lu[j].DestructionTime
	})

	return lu
}

// SortByDeletedTime sorts ListDisksUnattached by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lu ListDisksUnattached) SortByDeletedTime(inverse bool) ListDisksUnattached {
	if len(lu) < 2 {
		return lu
	}

	sort.Slice(lu, func(i, j int) bool {
		if inverse {
			return lu[i].DeletedTime > lu[j].DeletedTime
		}

		return lu[i].DeletedTime < lu[j].DeletedTime
	})

	return lu
}
