package disks

import (
	"context"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/interfaces"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/k8s"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/lb"
)

// FilterByID returns ListDisks with specified ID.
func (ld ListDisks) FilterByID(id uint64) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		return idisk.ID == id
	}

	return ld.FilterFunc(predicate)
}

// FilterByName returns ListDisks with specified Name.
func (ld ListDisks) FilterByName(name string) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		return idisk.Name == name
	}

	return ld.FilterFunc(predicate)
}

// FilterByStatus returns ListDisks with specified Status.
func (ld ListDisks) FilterByStatus(status string) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		return idisk.Status == status
	}

	return ld.FilterFunc(predicate)
}

// FilterByTechStatus returns ListDisks with specified TechStatus.
func (ld ListDisks) FilterByTechStatus(techStatus string) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		return idisk.TechStatus == techStatus
	}

	return ld.FilterFunc(predicate)
}

// FilterByImageID returns ListDisks with specified ImageID.
func (ld ListDisks) FilterByImageID(imageID uint64) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		return idisk.ImageID == imageID
	}

	return ld.FilterFunc(predicate)
}

// FilterByComputeID is used to filter ListDisks attached to specified compute.
func (ld ListDisks) FilterByComputeID(computeID uint64) ListDisks {
	predicate := func(idisk ItemDisk) bool {
		for k := range idisk.Computes {
			if k == strconv.FormatUint(computeID, 10) {
				return true
			}
		}

		return false
	}

	return ld.FilterFunc(predicate)
}

// FilterByK8SID is used to filter ListDisks by specified K8S cluster.
func (ld ListDisks) FilterByK8SID(ctx context.Context, k8sID uint64, decortClient interfaces.Caller) (*ListDisks, error) {
	caller := k8s.New(decortClient)

	req := k8s.GetRequest{
		K8SID: k8sID,
	}

	cluster, err := caller.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	var result ListDisks

	for _, masterCompute := range cluster.K8SGroups.Masters.DetailedInfo {
		result.Data = append(result.Data, ld.FilterByComputeID(masterCompute.ID).Data...)
	}

	for _, workerGroup := range cluster.K8SGroups.Workers {
		for _, workerCompute := range workerGroup.DetailedInfo {
			result.Data = append(result.Data, ld.FilterByComputeID(workerCompute.ID).Data...)
		}
	}

	return &result, nil
}

// FilterByLBID is used to filter ListDisks used by computes inside specified Load Balancer.
func (ld ListDisks) FilterByLBID(ctx context.Context, lbID uint64, decortClient interfaces.Caller) (*ListDisks, error) {
	caller := lb.New(decortClient)

	req := lb.GetRequest{
		LBID: lbID,
	}

	foundLB, err := caller.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	var result ListDisks
	result.Data = append(result.Data, ld.FilterByComputeID(foundLB.PrimaryNode.ComputeID).Data...)
	result.Data = append(result.Data, ld.FilterByComputeID(foundLB.SecondaryNode.ComputeID).Data...)

	return &result, nil
}

// FilterFunc allows filtering ListDisks based on a user-specified predicate.
func (ld ListDisks) FilterFunc(predicate func(ItemDisk) bool) ListDisks {
	var result ListDisks

	for _, item := range ld.Data {
		if predicate(item) {
			result.Data = append(result.Data, item)
		}
	}

	result.EntryCount = uint64(len(ld.Data))

	return result
}

// FindOne returns first found ItemDisk
// If none was found, returns an empty struct.
func (ld ListDisks) FindOne() ItemDisk {
	if len(ld.Data) == 0 {
		return ItemDisk{}
	}

	return ld.Data[0]
}
