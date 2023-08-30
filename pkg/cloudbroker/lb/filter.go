package lb

import (
	"context"

	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/k8s"
)

// FilterByID returns ListLB with specified ID.
func (ll ListLB) FilterByID(id uint64) ListLB {
	predicate := func(rlb ItemLBList) bool {
		return rlb.ID == id
	}

	return ll.FilterFunc(predicate)
}

// FilterByName returns ListLB with specified Name.
func (ll ListLB) FilterByName(name string) ListLB {
	predicate := func(rlb ItemLBList) bool {
		return rlb.Name == name
	}

	return ll.FilterFunc(predicate)
}

// FilterByExtNetID returns ListLB with specified ExtNetID.
func (ll ListLB) FilterByExtNetID(extNetID uint64) ListLB {
	predicate := func(rlb ItemLBList) bool {
		return rlb.ExtNetID == extNetID
	}

	return ll.FilterFunc(predicate)
}

// FilterByK8SID returns ListLB used by specified K8S cluster.
func (ll ListLB) FilterByK8SID(ctx context.Context, k8sID uint64, decortClient interfaces.Caller) (*ListLB, error) {
	caller := k8s.New(decortClient)

	req := k8s.GetRequest{
		K8SID: k8sID,
	}

	cluster, err := caller.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	predicate := func(rlb ItemLBList) bool {
		return cluster.LBID == rlb.ID
	}

	res := ll.FilterFunc(predicate)

	return &res, nil
}

// FilterFunc allows filtering ListLB based on a user-specified predicate.
func (ll ListLB) FilterFunc(predicate func(ItemLBList) bool) ListLB {
	var result ListLB

	for _, item := range ll.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(result.Data))

	return result
}

// FindOne returns first found RecordLB
// If none was found, returns an empty struct.
func (ll ListLB) FindOne() ItemLBList {
	if len(ll.Data) == 0 {
		return ItemLBList{}
	}

	return ll.Data[0]
}
