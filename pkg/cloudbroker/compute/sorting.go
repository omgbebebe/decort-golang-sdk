package compute

import "sort"

// SortByCPU sorts ListComputes by the CPU core amount in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lc ListComputes) SortByCPU(inverse bool) ListComputes {
	if len(lc.Data) < 2 {
		return lc
	}

	sort.Slice(lc.Data, func(i, j int) bool {
		if inverse {
			return lc.Data[i].CPUs > lc.Data[j].CPUs
		}

		return lc.Data[i].CPUs < lc.Data[j].CPUs
	})

	return lc
}

// SortByRAM sorts ListComputes by the RAM amount in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lc ListComputes) SortByRAM(inverse bool) ListComputes {
	if len(lc.Data) < 2 {
		return lc
	}

	sort.Slice(lc.Data, func(i, j int) bool {
		if inverse {
			return lc.Data[i].RAM > lc.Data[j].RAM
		}

		return lc.Data[i].RAM < lc.Data[j].RAM
	})

	return lc
}

// SortByCreatedTime sorts ListComputes by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lc ListComputes) SortByCreatedTime(inverse bool) ListComputes {
	if len(lc.Data) < 2 {
		return lc
	}

	sort.Slice(lc.Data, func(i, j int) bool {
		if inverse {
			return lc.Data[i].CreatedTime > lc.Data[j].CreatedTime
		}

		return lc.Data[i].CreatedTime < lc.Data[j].CreatedTime
	})

	return lc
}

// SortByUpdatedTime sorts ListComputes by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lc ListComputes) SortByUpdatedTime(inverse bool) ListComputes {
	if len(lc.Data) < 2 {
		return lc
	}

	sort.Slice(lc.Data, func(i, j int) bool {
		if inverse {
			return lc.Data[i].UpdatedTime > lc.Data[j].UpdatedTime
		}

		return lc.Data[i].UpdatedTime < lc.Data[j].UpdatedTime
	})

	return lc
}

// SortByDeletedTime sorts ListComputes by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lc ListComputes) SortByDeletedTime(inverse bool) ListComputes {
	if len(lc.Data) < 2 {
		return lc
	}

	sort.Slice(lc.Data, func(i, j int) bool {
		if inverse {
			return lc.Data[i].DeletedTime > lc.Data[j].DeletedTime
		}

		return lc.Data[i].DeletedTime < lc.Data[j].DeletedTime
	})

	return lc
}
