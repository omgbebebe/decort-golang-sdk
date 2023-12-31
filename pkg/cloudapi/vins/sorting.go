package vins

import "sort"

// SortByCreatedTime sorts ListVINS by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lv ListVINS) SortByCreatedTime(inverse bool) ListVINS {
	if len(lv.Data) < 2 {
		return lv
	}

	sort.Slice(lv.Data, func(i, j int) bool {
		if inverse {
			return lv.Data[i].CreatedTime > lv.Data[j].CreatedTime
		}

		return lv.Data[i].CreatedTime < lv.Data[j].CreatedTime
	})

	return lv
}

// SortByUpdatedTime sorts ListVINS by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lv ListVINS) SortByUpdatedTime(inverse bool) ListVINS {
	if len(lv.Data) < 2 {
		return lv
	}

	sort.Slice(lv.Data, func(i, j int) bool {
		if inverse {
			return lv.Data[i].UpdatedTime > lv.Data[j].UpdatedTime
		}

		return lv.Data[i].UpdatedTime < lv.Data[j].UpdatedTime
	})

	return lv
}

// SortByDeletedTime sorts ListVINS by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (lv ListVINS) SortByDeletedTime(inverse bool) ListVINS {
	if len(lv.Data) < 2 {
		return lv
	}

	sort.Slice(lv.Data, func(i, j int) bool {
		if inverse {
			return lv.Data[i].DeletedTime > lv.Data[j].DeletedTime
		}

		return lv.Data[i].DeletedTime < lv.Data[j].DeletedTime
	})

	return lv
}
