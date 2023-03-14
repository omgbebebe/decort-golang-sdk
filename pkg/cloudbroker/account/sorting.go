package account

import "sort"

// SortByCreatedTime sorts ListAccounts by the CreatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (la ListAccounts) SortByCreatedTime(inverse bool) ListAccounts {
	if len(la) < 2 {
		return la
	}

	sort.Slice(la, func(i, j int) bool {
		if inverse {
			return la[i].CreatedTime > la[j].CreatedTime
		}

		return la[i].CreatedTime < la[j].CreatedTime
	})

	return la
}

// SortByUpdatedTime sorts ListAccounts by the UpdatedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (la ListAccounts) SortByUpdatedTime(inverse bool) ListAccounts {
	if len(la) < 2 {
		return la
	}

	sort.Slice(la, func(i, j int) bool {
		if inverse {
			return la[i].UpdatedTime > la[j].UpdatedTime
		}

		return la[i].UpdatedTime < la[j].UpdatedTime
	})

	return la
}

// SortByDeletedTime sorts ListAccounts by the DeletedTime field in ascending order.
//
// If inverse param is set to true, the order is reversed.
func (la ListAccounts) SortByDeletedTime(inverse bool) ListAccounts {
	if len(la) < 2 {
		return la
	}

	sort.Slice(la, func(i, j int) bool {
		if inverse {
			return la[i].DeletedTime > la[j].DeletedTime
		}

		return la[i].DeletedTime < la[j].DeletedTime
	})

	return la
}
