package vins

import "sort"

// SortByCreatedTime sorts ListVINS by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lv ListVINS) SortByCreatedTime(inverse bool) ListVINS {
	if len(lv) < 2 {
		return lv
	}

	sort.Slice(lv, func(i, j int) bool {
		if inverse {
			return lv[i].CreatedTime > lv[j].CreatedTime
		}

		return lv[i].CreatedTime < lv[j].CreatedTime
	})

	return lv
}

// SortByUpdatedTime sorts ListVINS by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lv ListVINS) SortByUpdatedTime(inverse bool) ListVINS {
	if len(lv) < 2 {
		return lv
	}

	sort.Slice(lv, func(i, j int) bool {
		if inverse {
			return lv[i].UpdatedTime > lv[j].UpdatedTime
		}

		return lv[i].UpdatedTime < lv[j].UpdatedTime
	})

	return lv
}

// SortByDeletedTime sorts ListVINS by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lv ListVINS) SortByDeletedTime(inverse bool) ListVINS {
	if len(lv) < 2 {
		return lv
	}

	sort.Slice(lv, func(i, j int) bool {
		if inverse {
			return lv[i].DeletedTime > lv[j].DeletedTime
		}

		return lv[i].DeletedTime < lv[j].DeletedTime
	})

	return lv
}
