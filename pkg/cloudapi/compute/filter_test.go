package compute

import "testing"

var computes = ListComputes{
	ItemCompute{
		ACL:           []interface{}{},
		AccountID:     132847,
		AccountName:   "std_2",
		AffinityLabel: "",
		AffinityRules: []ItemRule{
			{
				GUID:     "",
				Key:      "aff_key",
				Mode:     "ANY",
				Policy:   "RECOMMENDED",
				Topology: "compute",
				Value:    "aff_val",
			},
		},
		AffinityWeight: 0,
		AntiAffinityRules: []ItemRule{
			{
				GUID:     "",
				Key:      "antiaff_key",
				Mode:     "ANY",
				Policy:   "RECOMMENDED",
				Topology: "compute",
				Value:    "antiaff_val",
			},
		},
		Architecture: "X86_64",
		BootOrder: []string{
			"hd", "cdrom",
		},
		BootDiskSize:   0,
		CloneReference: 0,
		Clones:         []uint64{},
		ComputeCIID:    0,
		CPU:            4,
		CreatedBy:      "timofey_tkachev_1@decs3o",
		CreatedTime:    1676975175,
		CustomFields:   map[string]interface{}{},
		DeletedBy:      "",
		DeletedTime:    0,
		Description:    "",
		Devices:        nil,
		Disks: []InfoDisk{
			{
				ID:      65191,
				PCISlot: 6,
			},
		},
		Driver:           "KVM_X86",
		GID:              212,
		GUID:             48500,
		ID:               48500,
		ImageID:          9884,
		Interfaces:       []ItemVNFInterface{},
		LockStatus:       "UNLOCKED",
		ManagerID:        0,
		ManagerType:      "",
		MigrationJob:     0,
		Milestones:       363500,
		Name:             "test",
		Pinned:           false,
		RAM:              4096,
		ReferenceID:      "c7cb19ac-af4a-4067-852f-c5572949207e",
		Registered:       true,
		ResName:          "compute-48500",
		RGID:             79724,
		RGName:           "std_broker2",
		SnapSets:         []ItemSnapSet{},
		StatelessSepID:   0,
		StatelessSepType: "",
		Status:           "ENABLED",
		Tags:             map[string]string{},
		TechStatus:       "STOPPED",
		TotalDiskSize:    2,
		UpdatedBy:        "",
		UpdatedTime:      1677058904,
		UserManaged:      true,
		VGPUs:            []uint64{},
		VINSConnected:    0,
		VirtualImageID:   0,
	},
	ItemCompute{
		ACL:               []interface{}{},
		AccountID:         132848,
		AccountName:       "std_broker",
		AffinityLabel:     "",
		AffinityRules:     []ItemRule{},
		AffinityWeight:    0,
		AntiAffinityRules: []ItemRule{},
		Architecture:      "X86_64",
		BootOrder: []string{
			"hd", "cdrom",
		},
		BootDiskSize:   0,
		CloneReference: 0,
		Clones:         []uint64{},
		ComputeCIID:    0,
		CPU:            6,
		CreatedBy:      "timofey_tkachev_1@decs3o",
		CreatedTime:    1677579436,
		CustomFields:   map[string]interface{}{},
		DeletedBy:      "",
		DeletedTime:    0,
		Description:    "",
		Devices:        nil,
		Disks: []InfoDisk{
			{
				ID:      65248,
				PCISlot: 6,
			},
		},
		Driver:           "KVM_X86",
		GID:              212,
		GUID:             48556,
		ID:               48556,
		ImageID:          9884,
		Interfaces:       []ItemVNFInterface{},
		LockStatus:       "UNLOCKED",
		ManagerID:        0,
		ManagerType:      "",
		MigrationJob:     0,
		Milestones:       363853,
		Name:             "compute_2",
		Pinned:           false,
		RAM:              4096,
		ReferenceID:      "a542c449-5b1c-4f90-88c5-7bb5f8ae68ff",
		Registered:       true,
		ResName:          "compute-48556",
		RGID:             79727,
		RGName:           "sdk_negative_fields_test",
		SnapSets:         []ItemSnapSet{},
		StatelessSepID:   0,
		StatelessSepType: "",
		Status:           "ENABLED",
		Tags:             map[string]string{},
		TechStatus:       "STARTED",
		TotalDiskSize:    1,
		UpdatedBy:        "",
		UpdatedTime:      1677579436,
		UserManaged:      true,
		VGPUs:            []uint64{},
		VINSConnected:    0,
		VirtualImageID:   0,
	},
}

func TestFilterByID(t *testing.T) {
    actual := computes.FilterByID(48500).FindOne()

    if actual.ID != 48500 {
        t.Fatal("expected ID 48500, found: ", actual.ID)
    }

    actualEmpty := computes.FilterByID(0)

    if len(actualEmpty) != 0 {
        t.Fatal("expected empty, actual: ", len(actualEmpty))
    }
}

func TestFilterByName(t *testing.T) {
    actual := computes.FilterByName("test").FindOne()

    if actual.Name != "test" {
        t.Fatal("expected compute with name 'test', found: ", actual.Name)
    }
}

func TestFilterByStatus(t *testing.T) {
    actual := computes.FilterByStatus("ENABLED")

    for _, item := range actual {
        if item.Status != "ENABLED" {
            t.Fatal("expected ENABLED status, found: ", item.Status)
        }
    }
}

func TestFilterByTechStatus(t *testing.T) {
    actual := computes.FilterByTechStatus("STARTED").FindOne()

    if actual.ID != 48556 {
        t.Fatal("expected 48556 with STARTED techStatus, found: ", actual.ID)
    }
}

func TestFilterByDiskID(t *testing.T) {
    actual := computes.FilterByDiskID(65248).FindOne()

    if actual.ID != 48556 {
        t.Fatal("expected 48556 with DiskID 65248, found: ", actual.ID)
    }
}

func TestFilterFunc(t *testing.T) {
    actual := computes.FilterFunc(func(ic ItemCompute) bool {
        return ic.Registered == true
    })

    if len(actual) != 2 {
        t.Fatal("expected 2 elements found, actual: ", len(actual))
    }

    for _, item := range actual {
        if item.Registered != true {
            t.Fatal("expected Registered to be true, actual: ", item.Registered)
        }
    }
}

func TestSortingByCreatedTime(t *testing.T) {
    actual := computes.SortByCreatedTime(false)

    if actual[0].Name != "test" {
        t.Fatal("expected 'test', found: ", actual[0].Name)
    }

    actual = computes.SortByCreatedTime(true)
    if actual[0].Name != "compute_2" {
        t.Fatal("expected 'compute_2', found: ", actual[0].Name)
    }
}

func TestSortingByCPU(t *testing.T) {
    actual := computes.SortByCPU(false)

    if actual[0].CPU != 4{
        t.Fatal("expected 4 CPU cores, found: ", actual[0].CPU)
    }

    actual = computes.SortByCPU(true)

    if actual[0].CPU != 6 {
        t.Fatal("expected 6 CPU cores, found: ", actual[0].CPU)
    }
}
