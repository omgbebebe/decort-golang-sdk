## Version 1.6.0

### Bugfix
- Fix client and legacy-client, the cyclicity of sending requests has been removed
- Edit description field SSLSkipVerify in config/config and config/legacy-config
- Remove tags required fields ExtNetID VINSID Start in model CreateRequest in cloudapi/lb/create and cloudbroker/lb/create
- Add tags required fields BindingName BindingAddress BindingPort in model FrontendBindRequest in cloudapi/lb/frontend_bind 
- Add tags required fields BindingAddress BindingPort in model FrontendBindUpdateRequest in cloudapi/lb/frontend_bind_update
- Remove tags omitempty in field Start in model kvmx86 and kvmppc in cloudapi and cloudbroker /create /mass_create
- Add field Driver in models cloudbroker/kvmx86/create and cloudbroker/kvmx86/create_blank
- Add field ExternalIp in model ItemDetailedInfo cloudbroker/k8s/models and cloudapi/k8s/models 
- Add fields StackID and Reason in model CreateRequest cloudbroker/kvmppc

### Feature
- Add field UserData in cloudapi/bservice/group_add
- Add fields VinsId, LbSysctlParams, HighlyAvailable, AdditionalSANs, InitConfiguration, ClusterConfiguration, KubeletConfiguration, KubeProxyConfiguration, JoinConfiguration, UserData, ExtNetOnly, OidcCertificate in model request cloudapi/k8s/create and cloudbroker/k8s/create  
- Add fields SysctlParams, HighlyAvailable in model CreateRequest in cloudapi/lb/create and cloudbroker/lb/create
- Add fields BackendHAIP, FrontendHAIP, PartK8s, SysctlParams in model RecordLB in cloudapi/lb/models and cloudbroker/lb/models
- Add models InfoStack, ItemStack, ListStacks cloudapi/stack/models
- Add field Routes and type Route in CreateInAccountRequest and CreateInRGRequest models in cloudapi/vins/create_in_rg and /cloudapi/vins/create_in_account
- Add field Enabled in model ItemVNFInterface in cloudapi/vins/models 
- Add fields Routes in models RecordNAT, RecordDHCP, RecordGW and add type ListStaticRoutes, ListRoutes, ItemRoutes in cloudapi/vins/models 

- Added new endpoints:
- cloudapi/k8s/get_worker_nodes_meta_data
- cloudapi/k8s/update_worker_nodes_meta_data
- cloudapi/lb/make_highly_available
- cloudapi/lb/updateSysctParams
- cloudapi/stack/get
- cloudapi/stack/list
- cloudapi/vins/static_route_list
- cloudapi/vins/static_route_access_grant
- cloudapi/vins/static_route_access_revoke
- cloudapi/vins/static_route_add
- cloudapi/vins/static_route_del
- cloudbroker/compute/set_custom_fields
- cloudbroker/k8s/get_worker_nodes_meta_data
- cloudbroker/k8s/update_worker_nodes_meta_data
- cloudbroker/lb/make_highly_available
- cloudbroker/lb/updateSysctParams
- cloudbroker/vins/static_route_list
- cloudbroker/vins/static_route_access_grant
- cloudbroker/vins/static_route_access_revoke
- cloudbroker/vins/static_route_add
- cloudbroker/vins/static_route_del
- cloudbroker/extnet/static_route_list
- cloudbroker/extnet/static_route_access_grant
- cloudbroker/extnet/static_route_access_revoke
- cloudbroker/extnet/static_route_add
- cloudbroker/extnet/static_route_del


