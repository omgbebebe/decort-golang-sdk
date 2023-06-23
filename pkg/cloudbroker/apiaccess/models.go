package apiaccess

type ItemAPIAccess struct {
	// APIs
	APIs APIsEndpoints `json:"apis"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Is default
	Default bool `json:"default"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Description
	Description string `json:"decs"`

	// GID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Is protected
	Protected bool `json:"protected"`

	// Status
	Status string `json:"status"`

	//Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`
}

type ListAPIAccess struct {
	Data []ItemAPIAccess `json:"data"`

	EntryCount uint64 `json:"entryCount"`
}

type APIsEndpoints struct {
	// CloudAPI endpoints
	CloudAPI CloudAPIEndpoints `json:"cloudapi,omitempty"`

	// CloudBroker endpoints
	CloudBroker CloudBrokerEndpoints `json:"cloudbroker,omitempty"`

	// LibCloud endpoints
	LibCloud LibCloudEndpoints `json:"libcloud,omitempty"`

	// System endpoints
	System SystemEndpoints `json:"system,omitempty"`
}

type CloudAPIEndpoints struct {
	Account        []string `json:"account,omitempty"`
	BService       []string `json:"bservice,omitempty"`
	CloudSpace     []string `json:"cloudspace,omitempty"`
	Compute        []string `json:"compute,omitempty"`
	ComputeCI      []string `json:"computeci,omitempty"`
	Disks          []string `json:"disks,omitempty"`
	ExtNet         []string `json:"extnet,omitempty"`
	FLIPGroup      []string `json:"flipgroup,omitempty"`
	GPU            []string `json:"gpu,omitempty"`
	Image          []string `json:"image,omitempty"`
	K8CI           []string `json:"k8ci,omitempty"`
	K8S            []string `json:"k8s,omitempty"`
	KVMPPC         []string `json:"kvmppc,omitempty"`
	KVMX86         []string `json:"kvmx86,omitempty"`
	LB             []string `json:"lb,omitempty"`
	Loactions      []string `json:"locations,omitempty"`
	Machine        []string `json:"machine,omitempty"`
	Openshift      []string `json:"openshift,omitempty"`
	OpenshiftCI    []string `json:"openshiftci,omitempty"`
	PCIDevice      []string `json:"pcidevice,omitempty"`
	PortForwarding []string `json:"portforwarding,omitempty"`
	Prometheus     []string `json:"prometheus,omitempty"`
	RG             []string `json:"rg,omitempty"`
	Sizes          []string `json:"sizes,omitempty"`
	Tasks          []string `json:"tasks,omitempty"`
	User           []string `json:"user,omitempty"`
	VGPU           []string `json:"vgpu,omitempty"`
	VINS           []string `json:"vins,omitempty"`
	All            bool     `json:"ALL,omitempty"`
}

type CloudBrokerEndpoints struct {
	Account        []string `json:"account,omitempty"`
	APIAccess      []string `json:"apiaccess,omitempty"`
	Audit          []string `json:"audit,omitempty"`
	AuditBeat      []string `json:"auditbeat,omitempty"`
	AuditCollector []string `json:"auditcollector,omitempty"`
	BackupCreator  []string `json:"backupcreator,omitempty"`
	BService       []string `json:"bservice,omitempty"`
	CloudSpace     []string `json:"cloudspace,omitempty"`
	Compute        []string `json:"compute,omitempty"`
	ComputeCI      []string `json:"computeci,omitempty"`
	Desnode        []string `json:"desnode,omitempty"`
	Diagnostics    []string `json:"diagnostics,omitempty"`
	Disks          []string `json:"disks,omitempty"`
	Eco            []string `json:"eco,omitempty"`
	ExtNet         []string `json:"extnet,omitempty"`
	FlIPgroup      []string `json:"flipgroup,omitempty"`
	Grid           []string `json:"grid,omitempty"`
	Group          []string `json:"group,omitempty"`
	Health         []string `json:"health,omitempty"`
	IaaS           []string `json:"iaas,omitempty"`
	Image          []string `json:"image,omitempty"`
	Job            []string `json:"job,omitempty"`
	K8CI           []string `json:"k8ci,omitempty"`
	K8S            []string `json:"k8s,omitempty"`
	KVMPPC         []string `json:"kvmppc,omitempty"`
	KVMX86         []string `json:"kvmx86,omitempty"`
	LB             []string `json:"lb,omitempty"`
	Machine        []string `json:"machine,omitempty"`
	Metering       []string `json:"metering,omitempty"`
	Milestones     []string `json:"milestones,omitempty"`
	Node           []string `json:"node,omitempty"`
	Openshift      []string `json:"openshift,omitempty"`
	OpenshiftCI    []string `json:"openshiftci,omitempty"`
	Ovsnode        []string `json:"ovsnode,omitempty"`
	PCIDevice      []string `json:"pcidevice,omitempty"`
	PGPU           []string `json:"pgpu,omitempty"`
	Prometheus     []string `json:"prometheus,omitempty"`
	QOS            []string `json:"qos,omitempty"`
	Resmon         []string `json:"resmon,omitempty"`
	RG             []string `json:"rg,omitempty"`
	Sep            []string `json:"sep,omitempty"`
	Stack          []string `json:"stack,omitempty"`
	Tasks          []string `json:"tasks,omitempty"`
	TLock          []string `json:"tlock,omitempty"`
	User           []string `json:"user,omitempty"`
	VGPU           []string `json:"vgpu,omitempty"`
	VINS           []string `json:"vins,omitempty"`
	VNFDev         []string `json:"vnfdev,omitempty"`
	ZeroAccess     []string `json:"zeroaccess,omitempty"`
	All            bool     `json:"ALL,omitempty"`
}

type LibCloudEndpoints struct {
	Libvirt []string `json:"libvirt,omitempty"`
	All     bool     `json:"ALL,omitempty"`
}

type SystemEndpoints struct {
	AgentController       []string `json:"agentcontroller,omitempty"`
	Alerts                []string `json:"alerts,omitempty"`
	Audits                []string `json:"audits,omitempty"`
	ContentManager        []string `json:"contentmanager,omitempty"`
	DocGenerator          []string `json:"docgenerator,omitempty"`
	EmailSender           []string `json:"emailsender,omitempty"`
	ErrorConditionHandler []string `json:"errorconditionhandler,omitempty"`
	GridManager           []string `json:"gridmanager,omitempty"`
	Health                []string `json:"health,omitempty"`
	Info                  []string `json:"info,omitempty"`
	InfoMGR               []string `json:"infomgr,omitempty"`
	Job                   []string `json:"job,omitempty"`
	Log                   []string `json:"log,omitempty"`
	Logo                  []string `json:"logo,omitempty"`
	Oauth                 []string `json:"oauth,omitempty"`
	Task                  []string `json:"task,omitempty"`
	UserManager           []string `json:"usermanager,omitempty"`
	All                   bool     `json:"ALL,omitempty"`
}
