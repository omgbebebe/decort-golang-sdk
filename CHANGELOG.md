## Version 1.6.6

## Bugfix
- Fixed model RecordExtNet in cloudbroker/extnet/models for correct work of get request
- Fixed json tags in ItemResourceConsumption model and delete extra model Consumed in cloudbroker/account/models
- Fixed statelessSepId field type from uint64 to int64 in cloudbroker/compute/models for correct work of list request

## Feature
- Added GetRaw and ListRaw methods that give response as an array of bytes for cloudAPI groups: account, compute, k8s, disks, rg, bservice, disks,extnet, flipgroup, image, k8ci, lb, locations(list), sizes(list), stack, tasks, vins
- Added GetRaw and ListRaw methods that give response as an array of bytes for cloudbroker groups: account, apiaccess, compute, disks, extnet, flipgroup, grid, group, image, k8ci, k8s, lb, rg, sep, stack, tasks, user, vgpu, vins