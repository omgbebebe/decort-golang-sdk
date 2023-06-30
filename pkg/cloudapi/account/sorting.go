package account

import "sort"

// SortByCreatedTime sorts ListAccounts by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (la ListAccounts) SortByCreatedTime(inverse bool) ListAccounts {
	if la.EntryCount < 2 {
		return la
	}

	sort.Slice(la.Data, func(i, j int) bool {
		if inverse {
			return la.Data[i].CreatedTime > la.Data[j].CreatedTime
		}

		return la.Data[i].CreatedTime < la.Data[j].CreatedTime
	})

	return la
}

// SortByUpdatedTime sorts ListAccounts by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (la ListAccounts) SortByUpdatedTime(inverse bool) ListAccounts {
	if la.EntryCount < 2 {
		return la
	}

	sort.Slice(la.Data, func(i, j int) bool {
		if inverse {
			return la.Data[i].UpdatedTime > la.Data[j].UpdatedTime
		}

		return la.Data[i].UpdatedTime < la.Data[j].UpdatedTime
	})

	return la
}

// SortByDeletedTime sorts ListAccounts by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (la ListAccounts) SortByDeletedTime(inverse bool) ListAccounts {
	if la.EntryCount < 2 {
		return la
	}

	sort.Slice(la.Data, func(i, j int) bool {
		if inverse {
			return la.Data[i].DeletedTime > la.Data[j].DeletedTime
		}

		return la.Data[i].DeletedTime < la.Data[j].DeletedTime
	})

	return la
}
