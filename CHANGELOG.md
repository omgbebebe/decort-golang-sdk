## Version 1.6.0

### Bugfix
- Fix cloudaApi/client and cloudapi/legacy-client, the cyclicity of sending requests has been removed
- Edit description field SSLSkipVerify in cloudapi/config/config and cloudapi/config/legacy-config
- Remove tags required fields ExtNetID VINSID Start in model CreateRequest in cloudapi/lb/create
- Add tags required fields BindingName BindingAddress BindingPort in model FrontendBindRequest in cloudapi/lb/frontend_bind 
- Add tags required fields BindingAddress BindingPort in model FrontendBindUpdateRequest in cloudapi/lb/frontend_bind_update

### Feature
- Add field UserData in cloudapi/bservice/group_add
- Add fields VinsId, LbSysctlParams, HighlyAvailable, AdditionalSANs, InitConfiguration, ClusterConfiguration, KubeletConfiguration, KubeProxyConfiguration, JoinConfiguration, UserData, ExtNetOnly, OidcCertificate in model request cloudapi/k8s/create    
- Add field Externalip in model ItemDetailedInfo in cloudapi/k8s/models 
- Add fields SysctlParams, HighlyAvailable in model CreateRequest in cloudapi/lb/create
- Add fields BackendHAIP, FrontendHAIP, PartK8s, SysctlParams in model RecordLB in cloudapi/lb/models
- Add models InfoStack, ItemStack, ListStacks cloudapi/stack/models
- Add field Routes and type Route in CreateInAccountRequest and CreateInRGRequest models in cloudapi/vins/create_in_rg and /cloudapi/vins/create_in_account
- Add field Enabled in model ItemVNFInterface 
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