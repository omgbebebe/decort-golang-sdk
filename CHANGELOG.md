## Version 1.5.4

### Feature
- Added cloudbroker/stack group of endpoints support. Added the possibility to get information about stacks (list handler) and about a specific stack by ID (get handler)
- Added cloudbroker/flipgroup group of endpoints support. Added the possibility to get information about flipgroup (list handler) and about a specific flipgroup by ID (get handler), add and remote compute, create, edit and delete flipgroup.


### Bugfix
- Fixed the field image/models/history.guid, added a custom unmarshall. Fixed an error with get information about images (handlers - list, get) in cloudbroker