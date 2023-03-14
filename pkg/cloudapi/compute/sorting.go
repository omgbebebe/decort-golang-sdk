package compute

import "sort"

// SortByCPU sorts ListComputes by the CPU core amount in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lc ListComputes) SortByCPU(inverse bool) ListComputes {
	if len(lc) < 2 {
		return lc
	}

	sort.Slice(lc, func(i, j int) bool {
		if inverse {
			return lc[i].CPU > lc[j].CPU
		}

		return lc[i].CPU < lc[j].CPU
	})

	return lc
}

// SortByRAM sorts ListComputes by the RAM amount in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lc ListComputes) SortByRAM(inverse bool) ListComputes {
	if len(lc) < 2 {
		return lc
	}

	sort.Slice(lc, func(i, j int) bool {
		if inverse {
			return lc[i].RAM > lc[j].RAM
		}

		return lc[i].RAM < lc[j].RAM
	})

	return lc
}

// SortByCreatedTime sorts ListComputes by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lc ListComputes) SortByCreatedTime(inverse bool) ListComputes {
	if len(lc) < 2 {
		return lc
	}

	sort.Slice(lc, func(i, j int) bool {
		if inverse {
			return lc[i].CreatedTime > lc[j].CreatedTime
		}

		return lc[i].CreatedTime < lc[j].CreatedTime
	})

	return lc
}

// SortByUpdatedTime sorts ListComputes by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lc ListComputes) SortByUpdatedTime(inverse bool) ListComputes {
	if len(lc) < 2 {
		return lc
	}

	sort.Slice(lc, func(i, j int) bool {
		if inverse {
			return lc[i].UpdatedTime > lc[j].UpdatedTime
		}

		return lc[i].UpdatedTime < lc[j].UpdatedTime
	})

	return lc
}

// SortByDeletedTime sorts ListComputes by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lc ListComputes) SortByDeletedTime(inverse bool) ListComputes {
	if len(lc) < 2 {
		return lc
	}

	sort.Slice(lc, func(i, j int) bool {
		if inverse {
			return lc[i].DeletedTime > lc[j].DeletedTime
		}

		return lc[i].DeletedTime < lc[j].DeletedTime
	})

	return lc
}
