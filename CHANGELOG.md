## Version 1.3.0

### Features

- Created CloudAPI/CloudBroker filtering, sorting and serialization functions for List requests.
	- Every handler with present List request has available FilterBy functions. Filtering by ID, Name is common for each handler.
	- In case user needs to filter response by uncommon field FilterFunc with user-specified predicate is also available.
	- CloudAPI/CloudBroker computes, disks and lb also have specific Filter methods predefined, to name a few:
		- computes: 
			- FilterByK8SID, used to filter computes used by specified k8s cluster;
			- FilterByK8SMasters, FilterByK8SWorkers, used to filter master/workers nodes. Best used after FilterByK8SID call;
			- FilterByLBID, used to filter computes used by specified load balancer;

		- disks:
			- FilterByK8SID, used to filter disks attached to computes inside specified k8s cluster;
			- FilterByLBID, used to filter disks attached to computes inside specified load balancer;

		- lb:
			- FilterByK8SID, used to filter load balancers used by specified k8s cluster;

- Reinvented request validation using go-validator. Made easier to manipulate and add on to.
	- Request/Config validation now uses tags instead of hard-coded validation functions;

- Added ability to parse client configuration from JSON or YAML formatted files.

### Bug Fixes

- Fixed SSO_URL trailing slash possibly breaking authentication process.
- Fixed cloudbroker/vins/nat_rule_add request model types.
- Fixed cloudbroker/grid DiskSize field type 
- Fixed TasksResult, InfoResult in cloudbroker/cloudapi/tasks/models JSON unmarshalling.

### Tests

- Covered CloudAPI/CloudBroker filters with unit tests.

### Other

- Updated module to new repository
