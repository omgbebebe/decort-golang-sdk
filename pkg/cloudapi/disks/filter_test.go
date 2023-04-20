package disks

import (
	"testing"
)

var disks = ListDisks{
	ItemDisk{
		MachineID:   0,
		MachineName: "",
		DeviceName:  "vda",
		AccountID:   132847,
		AccountName: "std_2",
		ACL:         map[string]interface{}{},
		Computes: map[string]string{
			"48500": "test",
		},
		CreatedTime:     1676975177,
		DeletedTime:     0,
		Description:     "",
		DestructionTime: 0,
		GID:             212,
		ID:              65191,
		ImageID:         9884,
		Images:          []uint64{},
		IOTune: IOTune{
			TotalIOPSSec: 2000,
		},
		Name:     "bootdisk",
		Order:    0,
		Params:   "",
		ParentID: 0,
		PCISlot:  6,
		Pool:     "vmstor",
		PresentTo: []uint64{
			27,
		},
		PurgeTime:  0,
		ResID:      "sample",
		ResName:    "sample",
		Role:       "",
		Shareable:  false,
		SizeMax:    2,
		SizeUsed:   2,
		Snapshots:  []ItemSnapshot{},
		Status:     "ASSIGNED",
		TechStatus: "ALLOCATED",
		Type:       "B",
		VMID:       48500,
	},
	ItemDisk{
		MachineID:   0,
		MachineName: "",
		DeviceName:  "vda",
		AccountID:   132852,
		AccountName: "std",
		ACL:         map[string]interface{}{},
		Computes: map[string]string{
			"48502": "stdvm2",
		},
		CreatedTime:     1676982606,
		DeletedTime:     0,
		Description:     "",
		DestructionTime: 0,
		GID:             212,
		ID:              65193,
		ImageID:         9885,
		Images:          []uint64{},
		IOTune: IOTune{
			TotalIOPSSec: 2000,
		},
		Name:     "bootdisk",
		Order:    0,
		Params:   "",
		ParentID: 0,
		PCISlot:  6,
		Pool:     "vmstor",
		PresentTo: []uint64{
			27,
			27,
		},
		PurgeTime:  0,
		ResID:      "sample",
		ResName:    "sample",
		Role:       "",
		Shareable:  false,
		SizeMax:    4,
		SizeUsed:   4,
		Snapshots:  []ItemSnapshot{},
		Status:     "ASSIGNED",
		TechStatus: "ALLOCATED",
		Type:       "B",
		VMID:       48502,
	},
}

func TestFilterByID(t *testing.T) {
	actual := disks.FilterByID(65193)

	if len(actual) == 0 {
		t.Fatal("No elements were found")
	}

	actualItem := actual.FindOne()

	if actualItem.ID != 65193 {
		t.Fatal("expected ID 65193, found: ", actualItem.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := disks.FilterByName("bootdisk")

	if len(actual) != 2 {
		t.Fatal("expected 2 elements, found: ", len(actual))
	}

	for _, item := range actual {
		if item.Name != "bootdisk" {
			t.Fatal("expected 'bootdisk' name, found: ", item.Name)
		}
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := disks.FilterByStatus("ASSIGNED")

	if len(actual) == 0 {
		t.Fatal("No elements were found")
	}

	for _, item := range actual {
		if item.Status != "ASSIGNED" {
			t.Fatal("expected 'ASSIGNED' status, found: ", item.Status)
		}
	}
}

func TestFilterByTechStatus(t *testing.T) {
	actual := disks.FilterByTechStatus("ALLOCATED")

	if len(actual) == 0 {
		t.Fatal("No elements were found")
	}

	for _, item := range actual {
		if item.TechStatus != "ALLOCATED" {
			t.Fatal("expected 'ALLOCATED' techStatus, found: ", item.TechStatus)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := disks.FilterFunc(func(id ItemDisk) bool {
		return len(id.PresentTo) == 2
	})

	if len(actual) == 0 {
		t.Fatal("No elements were found")
	}

	if len(actual[0].PresentTo) != 2 {
		t.Fatal("expected 2 elements in PresentTo, found: ", len(actual[0].PresentTo))
	}
}

func TestSortByCreatedTime(t *testing.T) {
	actual := disks.SortByCreatedTime(false)

	if actual[0].ID != 65191 {
		t.Fatal("expected ID 65191, found: ", actual[0].ID)
	}

	actual = disks.SortByCreatedTime(true)

	if actual[0].ID != 65193 {
		t.Fatal("expected ID 65193, found: ", actual[0].ID)
	}
}

var unattachedDisks = ListDisksUnattached{
	{
		CKey: "",
		Meta: []interface{}{
			"cloudbroker",
			"disk",
			1,
		},
		AccountID:       149,
		AccountName:     "test_account1",
		ACL:             map[string]interface{}{},
		BootPartition:   0,
		CreatedTime:     1681477547,
		DeletedTime:     0,
		Description:     "",
		DestructionTime: 0,
		DiskPath:        "",
		GID:             2002,
		GUID:            22636,
		ID:              22636,
		ImageID:         0,
		Images:          []uint64{},
		IOTune: IOTune{
			TotalIOPSSec: 2000,
		},
		IQN:                 "",
		Login:               "",
		Milestones:          43834,
		Name:                "test_disk",
		Order:               0,
		Params:              "",
		ParentID:            0,
		Password:            "",
		PCISlot:             -1,
		Pool:                "data05",
		PresentTo:           []uint64{},
		PurgeAttempts:       0,
		PurgeTime:           0,
		RealityDeviceNumber: 0,
		ReferenceID:         "",
		ResID:               "79bd3bd8-3424-48d3-963f-1870d506f169",
		ResName:             "volumes/volume_22636",
		Role:                "",
		SEPID:               1,
		Shareable:           false,
		SizeMax:             0,
		SizeUsed:            0,
		Snapshots:           nil,
		Status:              "CREATED",
		TechStatus:          "ALLOCATED",
		Type:                "D",
		VMID:                0,
	},
	{
		CKey: "",
		Meta: []interface{}{
			"cloudbroker",
			"disk",
			1,
		},
		AccountID:       150,
		AccountName:     "test_account",
		ACL:             map[string]interface{}{},
		BootPartition:   0,
		CreatedTime:     1681477558,
		DeletedTime:     0,
		Description:     "",
		DestructionTime: 0,
		DiskPath:        "",
		GID:             2002,
		GUID:            22637,
		ID:              22637,
		ImageID:         0,
		Images:          []uint64{},
		IOTune: IOTune{
			TotalIOPSSec: 2000,
		},
		IQN:        "",
		Login:      "",
		Milestones: 43834,
		Name:       "test_disk",
		Order:      0,
		Params:     "",
		ParentID:   0,
		Password:   "",
		PCISlot:    -1,
		Pool:       "data05",
		PresentTo: []uint64{
			27,
			27,
		},
		PurgeAttempts:       0,
		PurgeTime:           0,
		RealityDeviceNumber: 0,
		ReferenceID:         "",
		ResID:               "79bd3bd8-3424-48d3-963f-1870d506f169",
		ResName:             "volumes/volume_22637",
		Role:                "",
		SEPID:               1,
		Shareable:           false,
		SizeMax:             0,
		SizeUsed:            0,
		Snapshots:           nil,
		Status:              "CREATED",
		TechStatus:          "ALLOCATED",
		Type:                "B",
		VMID:                0,
	},
}

func TestListDisksUnattached_FilterByID(t *testing.T) {
	actual := unattachedDisks.FilterByID(22636)

	if len(actual) == 0 {
		t.Fatal("No elements were found")
	}

	actualItem := actual.FindOne()

	if actualItem.ID != 22636 {
		t.Fatal("expected ID 22636, found: ", actualItem.ID)
	}
}

func TestListDisksUnattached_FilterByName(t *testing.T) {
	actual := unattachedDisks.FilterByName("test_disk")

	if len(actual) != 2 {
		t.Fatal("expected 2 elements, found: ", len(actual))
	}

	for _, item := range actual {
		if item.Name != "test_disk" {
			t.Fatal("expected 'test_disk' name, found: ", item.Name)
		}
	}
}

func TestListDisksUnattached_FilterByStatus(t *testing.T) {
	actual := unattachedDisks.FilterByStatus("CREATED")

	if len(actual) == 0 {
		t.Fatal("No elements were found")
	}

	for _, item := range actual {
		if item.Status != "CREATED" {
			t.Fatal("expected 'CREATED' status, found: ", item.Status)
		}
	}
}

func TestListDisksUnattached_FilterByTechStatus(t *testing.T) {
	actual := unattachedDisks.FilterByTechStatus("ALLOCATED")

	if len(actual) == 0 {
		t.Fatal("No elements were found")
	}

	for _, item := range actual {
		if item.TechStatus != "ALLOCATED" {
			t.Fatal("expected 'ALLOCATED' techStatus, found: ", item.TechStatus)
		}
	}
}

func TestListDisksUnattached_FilterFunc(t *testing.T) {
	actual := unattachedDisks.FilterFunc(func(id ItemDiskUnattached) bool {
		return len(id.PresentTo) == 2
	})

	if len(actual) == 0 {
		t.Fatal("No elements were found")
	}

	if len(actual[0].PresentTo) != 2 {
		t.Fatal("expected 2 elements in PresentTo, found: ", len(actual[0].PresentTo))
	}
}

func TestListDisksUnattached_SortByCreatedTime(t *testing.T) {
	actual := unattachedDisks.SortByCreatedTime(false)

	if actual[0].ID != 22636 {
		t.Fatal("expected ID 22636, found: ", actual[0].ID)
	}

	actual = unattachedDisks.SortByCreatedTime(true)

	if actual[0].ID != 22637 {
		t.Fatal("expected ID 22637, found: ", actual[0].ID)
	}

}
