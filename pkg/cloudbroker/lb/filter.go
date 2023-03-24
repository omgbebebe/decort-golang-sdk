package lb

import (
	"context"

	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/k8s"
)

// FilterByID returns ListLB with specified ID.
func (ll ListLB) FilterByID(id uint64) ListLB {
	predicate := func(rlb RecordLB) bool {
		return rlb.ID == id
	}

	return ll.FilterFunc(predicate)
}

// FilterByName returns ListLB with specified Name.
func (ll ListLB) FilterByName(name string) ListLB {
	predicate := func(rlb RecordLB) bool {
		return rlb.Name == name
	}

	return ll.FilterFunc(predicate)
}

// FilterByExtNetID returns ListLB with specified ExtNetID.
func (ll ListLB) FilterByExtNetID(extNetID uint64) ListLB {
	predicate := func(rlb RecordLB) bool {
		return rlb.ExtNetID == extNetID
	}

	return ll.FilterFunc(predicate)
}

// FilterByImageID returns ListLB with specified ImageID.
func (ll ListLB) FilterByImageID(imageID uint64) ListLB {
	predicate := func(rlb RecordLB) bool {
		return rlb.ImageID == imageID
	}

	return ll.FilterFunc(predicate)
}

// FilterByK8SID returns ListLB used by specified K8S cluster.
func (ll ListLB) FilterByK8SID(ctx context.Context, k8sID uint64, decortClient interfaces.Caller) (ListLB, error) {
	caller := k8s.New(decortClient)

	req := k8s.GetRequest{
		K8SID: k8sID,
	}

	cluster, err := caller.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	predicate := func(rlb RecordLB) bool {
		return cluster.LBID == rlb.ID
	}

	return ll.FilterFunc(predicate), nil
}

// FilterFunc allows filtering ListLB based on a user-specified predicate.
func (ll ListLB) FilterFunc(predicate func(RecordLB) bool) ListLB {
	var result ListLB

	for _, item := range ll {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FindOne returns first found RecordLB
// If none was found, returns an empty struct.
func (ll ListLB) FindOne() RecordLB {
	if len(ll) == 0 {
		return RecordLB{}
	}

	return ll[0]
}
