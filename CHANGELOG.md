## Version 1.5.2

### Bugfix
- Fix tag 'url' VINSID field in cloudbroker/account/list_vins

- Made the field Reason not required in cloudbroker/account/enable and cloudbroker/account/disable 

- Made the field RecursiveDelete required in cloudbroker/account/deleteUser 

- Add a validator function to a workersGroupName field - must be 3 or more symbol 

- Add a token actuality check, add an error handler in client/client-transport and client/http-client

- Add a fields SEPID and Pool in ListRequest struct in cloudbroker/disks/list and cloudapi/disks/list
