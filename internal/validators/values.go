package validators

var (
	driverValues     = []string{"KVM_X86", "KVM_PPC"}
	accessTypeValues = []string{"R", "RCX", "ARCXDU"}
	resTypesValues   = []string{"compute", "vins", "k8s", "openshift", "lb", "flipgroup"}
	protoValues      = []string{"tcp", "udp"}

	accountCUTypeValues = []string{"CU_M", "CU_C", "CU_D", "CU_DM", "CU_S", "CU_A", "CU_NO", "CU_I", "CU_NP"}

	bserviceModeValues = []string{"ABSOLUTE", "RELATIVE"}

	computeTopologyValues  = []string{"compute", "node"}
	computePolicyValues    = []string{"RECOMMENDED", "REQUIRED"}
	computeModeValues      = []string{"EQ", "EN", "ANY"}
	computeDiskTypeValues  = []string{"D", "B"}
	computeNetTypeValues   = []string{"EXTNET", "VINS"}
	computeOrderValues     = []string{"cdrom", "network", "hd"}
	computeDataDisksValues = []string{"KEEP", "DETACH", "DESTROY"}

	diskTypeValues = []string{"B", "T", "D"}

	flipgroupClientTypeValues = []string{"compute", "vins"}

	kvmNetTypeValues = []string{"EXTNET", "VINS", "NONE"}

	lbAlgorithmValues = []string{"roundrobin", "static-rr", "leastconn"}

	rgDefNetValues  = []string{"PRIVATE", "PUBLIC", "NONE"}
	rgNetTypeValues = []string{"PUBLIC", "PRIVATE"}

	vinsTypeValues = []string{"DHCP", "VIP", "EXCLUDE"}

	imageBootTypeValues     = []string{"uefi", "bios"}
	imageTypeValues         = []string{"windows", "linux", "other"}
	imageDriversValues      = []string{"KVM_X86"}
	imageArchitectureValues = []string{"X86_64", "PPC64_LE"}

	sepFieldTypeValues = []string{"int", "str", "bool", "list", "dict"}

	networkPluginValues = []string{"flannel", "weawenet", "calico"}

	strictLooseValues = []string{"strict", "loose"}

	interfaceStateValues = []string{"on", "off"}
)
