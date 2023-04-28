## Version 1.4.0

### Features

- Actualized SDK to platform version 3.8.6:
  - Added required field networkPlugin to requests in:
    - /cloudapi/cloudbroker/k8s/create;
    - /cloudbroker/k8ci/create.
  - Added networkPlugin field in models:
    - /cloudapi/cloudbroker/k8s;
    - /cloudbroker/k8ci.
  - Updated list of compute objects fields and added list of group objects in bservice model.
  - Added cpuAllocationRatio and cpuAllocationParameter fields in models:
    - /cloudapi/cloudbroker/rg;
    - /cloudapi/cloudbroker/account.
  - Added setCpuAllocationRatio endpoint support in:
    - /cloudbroker/account;
    - /cloudbroker/rg.
  - Added /cloudbrocker/grid/setCpuAllocationRatioForVM endpoint support.
  - Added setCpuAllocationParameter endpoint support in:
    - /cloudbroker/account;
    - /cloudbroker/rg;
    - /cloudbroker/grid.
  - Added cloudapi/cloudbroker/compute/changeLinkState endpoint support.

- Added enabled field in cloudapi/compute models:
  - interfaces in compute/list response;
  - RecordNetAttach (compute/netAttach response).

- Added cloudapi/compute/bootOrderSet endpoint support.

- Added cloudapi/compute/bootOrderGet endpoint support.

### Bug Fixes

- Fixed pciSlot field type in models:
  - cloudapi/cloudbroker/computes;
  - cloudapi/cloudbroker/vins.

- Fixed handling cloudapi/account/restore endpoint response (panicked when marhalling).

- Added missing field diskType in cloudapi/compute/diskAttach request.

- Added missing eBurst field in cloudapi/extnet QOS model.
