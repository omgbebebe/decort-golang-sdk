package bservice

import "sort"

// SortByCreatedTime sorts ListBasicServices by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lbs ListBasicServices) SortByCreatedTime(inverse bool) ListBasicServices {
	if len(lbs) < 2 {
		return lbs
	}

	sort.Slice(lbs, func(i, j int) bool {
		if inverse {
			return lbs[i].CreatedTime > lbs[j].CreatedTime
		}

		return lbs[i].CreatedTime < lbs[j].CreatedTime
	})

	return lbs
}

// SortByUpdatedTime sorts ListBasicServices by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lbs ListBasicServices) SortByUpdatedTime(inverse bool) ListBasicServices {
	if len(lbs) < 2 {
		return lbs
	}

	sort.Slice(lbs, func(i, j int) bool {
		if inverse {
			return lbs[i].UpdatedTime > lbs[j].UpdatedTime
		}

		return lbs[i].UpdatedTime < lbs[j].UpdatedTime
	})

	return lbs
}

// SortByDeletedTime sorts ListBasicServices by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lbs ListBasicServices) SortByDeletedTime(inverse bool) ListBasicServices {
	if len(lbs) < 2 {
		return lbs
	}

	sort.Slice(lbs, func(i, j int) bool {
		if inverse {
			return lbs[i].DeletedTime > lbs[j].DeletedTime
		}

		return lbs[i].DeletedTime < lbs[j].DeletedTime
	})

	return lbs
}
