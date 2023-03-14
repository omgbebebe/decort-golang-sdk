package disks

import "testing"

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
