package bservice

import "sort"

// SortByCreatedTime sorts ListBasicServices by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lbs ListBasicServices) SortByCreatedTime(inverse bool) ListBasicServices {
	if lbs.EntryCount < 2 {
		return lbs
	}

	sort.Slice(lbs.Data, func(i, j int) bool {
		if inverse {
			return lbs.Data[i].CreatedTime > lbs.Data[j].CreatedTime
		}

		return lbs.Data[i].CreatedTime < lbs.Data[j].CreatedTime
	})

	return lbs
}

// SortByUpdatedTime sorts ListBasicServices by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lbs ListBasicServices) SortByUpdatedTime(inverse bool) ListBasicServices {
	if lbs.EntryCount < 2 {
		return lbs
	}

	sort.Slice(lbs.Data, func(i, j int) bool {
		if inverse {
			return lbs.Data[i].UpdatedTime > lbs.Data[j].UpdatedTime
		}

		return lbs.Data[i].UpdatedTime < lbs.Data[j].UpdatedTime
	})

	return lbs
}

// SortByDeletedTime sorts ListBasicServices by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lbs ListBasicServices) SortByDeletedTime(inverse bool) ListBasicServices {
	if lbs.EntryCount < 2 {
		return lbs
	}

	sort.Slice(lbs.Data, func(i, j int) bool {
		if inverse {
			return lbs.Data[i].DeletedTime > lbs.Data[j].DeletedTime
		}

		return lbs.Data[i].DeletedTime < lbs.Data[j].DeletedTime
	})

	return lbs
}
