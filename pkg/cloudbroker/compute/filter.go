package compute

import (
	"context"

	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/k8s"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/lb"
)

// FilterByID returns ListComputes with specified ID.
func (lc ListComputes) FilterByID(id uint64) ListComputes {
	predicate := func(ic ItemCompute) bool {
		return ic.ID == id
	}

	return lc.FilterFunc(predicate)
}

// FilterByName returns ListComputes with specified Name.
func (lc ListComputes) FilterByName(name string) ListComputes {
	predicate := func(ic ItemCompute) bool {
		return ic.Name == name
	}

	return lc.FilterFunc(predicate)
}

// FilterByStatus returns ListComputes with specified Status.
func (lc ListComputes) FilterByStatus(status string) ListComputes {
	predicate := func(ic ItemCompute) bool {
		return ic.Status == status
	}

	return lc.FilterFunc(predicate)
}

// FilterByTechStatus returns ListComputes with specified TechStatus.
func (lc ListComputes) FilterByTechStatus(techStatus string) ListComputes {
	predicate := func(ic ItemCompute) bool {
		return ic.TechStatus == techStatus
	}

	return lc.FilterFunc(predicate)
}

// FilterByDiskID return ListComputes with specified DiskID.
func (lc ListComputes) FilterByDiskID(diskID uint64) ListComputes {
	predicate := func(ic ItemCompute) bool {
		for _, disk := range ic.Disks {
			if disk.ID == diskID {
				return true
			}
		}
		return false
	}

	return lc.FilterFunc(predicate)
}

// FilterByK8SID returns master and worker nodes (ListComputes) inside specified K8S cluster.
func (lc ListComputes) FilterByK8SID(ctx context.Context, k8sID uint64, decortClient interfaces.Caller) (*ListComputes, error) {
	caller := k8s.New(decortClient)

	req := k8s.GetRequest{
		K8SID: k8sID,
	}

	cluster, err := caller.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	predicate := func(ic ItemCompute) bool {
		for _, info := range cluster.K8SGroups.Masters.DetailedInfo {
			if info.ID == ic.ID {
				return true
			}
		}

		for _, worker := range cluster.K8SGroups.Workers {
			for _, info := range worker.DetailedInfo {
				if info.ID == ic.ID {
					return true
				}
			}
		}

		return false
	}

	result := lc.FilterFunc(predicate)

	return &result, nil
}

// K8SMasters is used to filter master nodes. Best used after FilterByK8SID function.
func (lc ListComputes) FilterByK8SMasters() ListComputes {
	predicate := func(ic ItemCompute) bool {
		for _, rule := range ic.AntiAffinityRules {
			if rule.Value == "master" {
				return true
			}
		}
		return false
	}

	return lc.FilterFunc(predicate)
}

// K8SMasters is used to filter worker nodes. Best used after FilterByK8SID function.
func (lc ListComputes) FilterByK8SWorkers() ListComputes {
	predicate := func(ic ItemCompute) bool {
		for _, rule := range ic.AntiAffinityRules {
			if rule.Value == "worker" {
				return true
			}
		}
		return false
	}

	return lc.FilterFunc(predicate)
}

// FilterByLBID is used to filter ListComputes used by specified Load Balancer.
func (lc ListComputes) FilterByLBID(ctx context.Context, lbID uint64, decortClient interfaces.Caller) (*ListComputes, error) {
	caller := lb.New(decortClient)

	req := lb.GetRequest{
		LBID: lbID,
	}

	foundLB, err := caller.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	predicate := func(ic ItemCompute) bool {
		return ic.ID == foundLB.PrimaryNode.ComputeID || ic.ID == foundLB.SecondaryNode.ComputeID
	}

	result := lc.FilterFunc(predicate)

	return &result, nil
}

// FilterFunc allows filtering ListComputes based on a user-specified predicate.
func (lc ListComputes) FilterFunc(predicate func(ItemCompute) bool) ListComputes {
	var result ListComputes

	for _, item := range lc.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(result.Data))

	return result
}

// FindOne returns first found ItemCompute
// If none was found, returns an empty struct.
func (lc ListComputes) FindOne() ItemCompute {
	if len(lc.Data) == 0 {
		return ItemCompute{}
	}

	return lc.Data[0]
}
