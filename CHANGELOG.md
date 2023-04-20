## Version 1.3.1

### Features

- Added FilterByGID for cloudapi/locations/list handler response, used to filter locations by specified GID.
- Added /cloudbroker/pcidevices endpoints support
  - /cloudbroker/pcidevices/create
  - /cloudbroker/pcidevices/delete
  - /cloudbroker/pcidevices/disable
  - /cloudbroker/pcidevices/enable
  - /cloudbroker/pcidevices/list
- Added /cloudbroker/vgpu endpoints support
  - /cloudbroker/vgpu/allocate
  - /cloudbroker/vgpu/create
  - /cloudbroker/vgpu/deallocate
  - /cloudbroker/vgpu/destroy
  - /cloudbroker/vgpu/list

### Bug Fixes

- Fixed cloudbroker/cloudapi/account/update request model types.
- Fixed cloudbroker/cloudapi/rg/update request model types.
- Fixed cloudapi/account DeactivationTime field type.
- Fixed cloudapi/k8s/workersGroupAdd return value type.
- Fixed cloudapi/disks/listUnattached return value type.
- Added ListDisksUnattached model as a cloudapi/disks/listUnattached handler response with filters.
- Fixed cloudapi/extnet Excluded field type.
- Fixed cloudapi/rg RecordResourceUsage model.
- Fixed cloudapi/compute ItemACL model.

### Tests

- Covered cloudapi/disks ListDisksUnattached filters with unit tests.
