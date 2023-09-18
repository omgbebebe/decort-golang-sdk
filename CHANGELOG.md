## Version 1.5.8

### Bugfix
- Fix model the RecordK8CI to cloudbroker/k8ci/models to correctly receive information on get request
- Add model the RecordK8CIList to cloudbroker/k8ci/models to correctly receive information on list request
- Refactored clients to fix the problems with concurrency safety
- Add the Routes field in the CreateInRGRequest and CreateInAccountRequest models in cb/vins. The fields for creating the resource are matched

For get request to work correctly:
- Fix the Rules field (fix type) in the NATConfig model in cb/vins/models and ca/vins/models
- Fix the InfoVNF model (remove the excess field, add the Routes field) in cb/vins/models 
- Add the ItemRoutes model in cb/vins/models 
- Fix the RecordVINS model (remove the excess fields) in cb/vins/models 
