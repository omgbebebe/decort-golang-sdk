package validators

var (
	accountAccessTypeValues = []string{"R", "RCX", "ARCXDU"}
	accountCUTypeValues     = []string{"CU_M", "CU_C", "CU_D", "CU_S", "CU_A", "CU_NO", "CU_I", "CU_NP"}

	bserviceDriverValues = []string{"KVM_X86, KVM_PPC"}
	bserviceModeValues   = []string{"ABSOLUTE", "RELATIVE"}

	computeTopologyValues = []string{"compute", "node"}
	computePolicyValues   = []string{"RECOMMENDED", "REQUIRED"}
	computeModeValues     = []string{"EQ", "EN", "ANY"}
	computeDiskTypeValues = []string{"D", "B"}
	computeNetTypeValues  = []string{"EXTNET", "VINS"}
	computeProtoValues    = []string{"tcp", "udp"}

	imageBootTypeValues     = []string{"uefi", "bios"}
	imageTypeValues         = []string{"windows", "linux", "other"}
	imageDriversValues      = []string{"KVM_X86"}
	imageArchitectureValues = []string{"X86_64", "PPC64_LE"}
)
