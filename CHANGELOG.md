## Version 1.5.0

### Feature

- Updated lists responses models in:
- - cloudapi/account/list
- - cloudbroker/account/list
- - cloudapi/bservice/list
- - cloudapi/compute/list
- - cloudbroker/compute/list
- - cloudapi/computeci/list
- - cloudapi/disks/list
- - cloudbroker/disks/list
- - cloudapi/extnet/list
- - cloudbroker/extnet/list
- - cloudapi/flipgroup/list
- - cloudapi/image/list
- - cloudbroker/image/list
- - cloudapi/k8ci/list
- - cloudbroker/k8ci/list
- - cloudapi/k8s/list
- - cloudapi/k8s/listDeleted
- - cloudbroker/k8s/list
- - cloudbroker/k8s/listDeleted
- - cloudapi/tasks/list
- - cloudbroker/tasks/list
- - cloudapi/rg/list
- - cloudbroker/rg/list
- - cloudapi/vins/list
- - cloudbroker/vins/list
- - cloudapi/account/listDeleted
- - cloudapi/account/listCompute
- - cloudapi/account/listDisks
- - cloudapi/account/listFlipGroups
- - cloudapi/account/listRG
- - cloudapi/account/listVINS
- - cloudapi/compute/listDeleted
- - cloudapi/disks/listDeleted
- - cloudapi/disks/listTypes
- - cloudapi/k8ci/listDeleted
- - cloudapi/lb/list
- - cloudapi/lb/listDeleted
- - cloudbroker/lb/list
- - cloudbroker/lb/listDeleted
- - cloudapi/rg/listComputes
- - cloudapi/rg/listDeleted
- - cloudapi/rg/listLb
- - cloudapi/rg/listPFW
- - cloudapi/rg/listVins
- - cloudapi/vins/listDeleted
- - cloudbroker/account/listComputes
- - cloudbroker/account/listDeleted
- - cloudbroker/account/listDisks 
- - cloudbroker/account/listFlipGroups
- - cloudbroker/account/listRG
- - cloudbroker/account/listVINS
- - cloudbroker/compute/listDeleted
- - cloudapi/compute/listPciDevice
- - cloudbroker/compute/listPciDevice
- - cloudapi/compute/listVGpu
- - cloudbroker/compute/listVGpu
- - cloudbroker/disks/listTypes
- - cloudbroker/grid/list
- - cloudbroker/grid/listEmails
- - cloudbroker/k8ci/listDeleted
- - cloudbroker/pcidevice/list
- - cloudbroker/rg/affinityGroupsList
- - cloudbroker/rg/listDeleted
- - cloudbroker/rg/listComputes
- - cloudbroker/rg/listLB
- - cloudbroker/rg/listPfw
- - cloudbroker/rg/listResourceConsumption
- - cloudbroker/rg/listVins
- - cloudbroker/sep/list
- - cloudbroker/vgpu/list
- - cloudbroker/vins/extnetList
- - cloudbroker/vins/IpList
- - cloudbroker/vins/natRuleList

- Added new endpoints:
- - cloudapi/rg/getResourceConsumption
- - cloudapi/rg/listResourceConsumption
- - cloudbroker/rg/getResourceConsumption
- - cloudbroker/rg/listResourceConsumption
- - cloudapi/account/getResourceConsumption
- - cloudapi/account/listResourceConsumption
- - cloudbroker/account/getResourceConsumption
- - cloudbroker/account/listResourceConsumption
- - cloudbroker/grid/getResourceConsumption
- - cloudbroker/grid/listResourceConsumption

- Added field CU_DM to ResourceLimits model (account, rg)

- Added field ReferenceID to SnapshotExtended model in cloudapi/compute/get

- Added field Interfaces in requests:
- - cloudapi/kvmppc/create 
- - cloudapi/kvmppc/createBlank
- - cloudapi/kvmx86/create 
- - cloudapi/kvmx86/createBlank
- - cloudbroker/kvmppc/create 
- - cloudbroker/kvmppc/createBlank
- - cloudbroker/kvmx86/create
- - cloudbroker/kvmx86/createBlank

- Added UpdatedBy field to task model in cloudbroker/task/get

- Made optional fields in requests:
- - Reason (cloudbroker/account/delete)
- - Reason (cloudbroker/account/restore)
- - Gateway (cloudbroker/extnet/create)
- - Reason (cloudbroker/image/delete)
- - Num (cloudapi/k8s/workerAdd)
- - NetID (cloudbroker/vins/extnetConnect)

- Updated cloudapi/rg/get model

- Deleted field Username from cloudbroker/account/update request

- Deleted field EmailAddress from cloudbroker/account/update request

- Added field DiskType to cloudbroker/compute/diskAttach request

- Added field Reason to cloudbroker/compute/diskQos request

- Added field Enabled to cloudbroker/compute/netAttach response model

- Added field CPUAllocationRatio to cloudbroker/image/listStacks response model

- Added field Descr to cloudbroker/image/listStacks response model

- Added field MemAllocationRatio to cloudbroker/image/listStacks response model

- Updated cloudapi/k8s/workersGroupByName response model

- Deleted field LBImageID from cloudbroker/k8ci/create request

- Deleted field ImageID from cloudbroker/lb/create request

- Deleted field Reason from cloudbroker/vins/extnetList request

### Bugfix
- Changed the Excluded field type in cloudbroker/extnet/get response model
