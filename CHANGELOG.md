## Version 1.5.7

### Bugfix
- Remove the required tag of the start field in the CreateRequest model in cb/lb/create, since it is impossible to create an lb without starting it
- Fix model the RecordGrid, add the ItemGridList model to cloudbroker/grid/models to correctly receive information on get and list requests
- Fix tag json field GID in model RecordResourcesConsumption cb/grid/models