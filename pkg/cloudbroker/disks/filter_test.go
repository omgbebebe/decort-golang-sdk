package disks

import "testing"

var disks = ListDisks{
	Data: []ItemDisk{
		{
			MachineID:   0,
			MachineName: "",
			RecordDisk: RecordDisk{
				DeviceName: "vda",
				SEPType:    "",
				InfoDisk: InfoDisk{
					AccountID:     132847,
					AccountName:   "std_2",
					ACL:           map[string]interface{}{},
					BootPartition: 0,
					Computes: map[string]string{
						"48500": "test",
					},
					CreatedTime:     1676975177,
					DeletedTime:     0,
					Description:     "",
					DestructionTime: 0,
					DiskPath:        "",
					GID:             212,
					GUID:            65191,
					ID:              65191,
					ImageID:         9884,
					Images:          []uint64{},
					IOTune: IOTune{
						TotalIOPSSec: 2000,
					},
					IQN:        "",
					Login:      "",
					Milestones: 363501,
					Name:       "bootdisk",
					Order:      0,
					Params:     "",
					ParentID:   0,
					Password:   "",
					PCISlot:    6,
					Pool:       "vmstor",
					PresentTo: []uint64{
						27,
					},
					PurgeAttempts:       0,
					PurgeTime:           0,
					RealityDeviceNumber: 0,
					ReferenceID:         "sample",
					ResID:               "sample",
					ResName:             "sample",
					Role:                "",
					SEPID:               2504,
					Shareable:           false,
					SizeMax:             2,
					SizeUsed:            2,
					Snapshots:           []ItemSnapshot{},
					Status:              "ASSIGNED",
					TechStatus:          "ALLOCATED",
					Type:                "B",
					VMID:                48500,
				},
			},
		},
		{
			MachineID:   0,
			MachineName: "",
			RecordDisk: RecordDisk{
				DeviceName: "vda",
				SEPType:    "",
				InfoDisk: InfoDisk{
					AccountID:     132852,
					AccountName:   "std",
					ACL:           map[string]interface{}{},
					BootPartition: 0,
					Computes: map[string]string{
						"48502": "stdvm2",
					},
					CreatedTime:     1676982606,
					DeletedTime:     0,
					Description:     "",
					DestructionTime: 0,
					DiskPath:        "",
					GID:             212,
					GUID:            65193,
					ID:              65193,
					ImageID:         9885,
					Images:          []uint64{},
					IOTune: IOTune{
						TotalIOPSSec: 2000,
					},
					IQN:        "",
					Login:      "",
					Milestones: 363516,
					Name:       "bootdisk",
					Order:      0,
					Params:     "",
					ParentID:   0,
					Password:   "",
					PCISlot:    6,
					Pool:       "vmstor",
					PresentTo: []uint64{
						27,
						27,
					},
					PurgeAttempts:       0,
					PurgeTime:           0,
					RealityDeviceNumber: 0,
					ReferenceID:         "sample",
					ResID:               "sample",
					ResName:             "sample",
					Role:                "",
					SEPID:               2504,
					Shareable:           false,
					SizeMax:             4,
					SizeUsed:            4,
					Snapshots:           []ItemSnapshot{},
					Status:              "ASSIGNED",
					TechStatus:          "ALLOCATED",
					Type:                "B",
					VMID:                48502,
				},
			},
		},
	},
	EntryCount: 2,
}

func TestFilterByID(t *testing.T) {
	actual := disks.FilterByID(65193)

	if len(actual.Data) == 0 {
		t.Fatal("No elements were found")
	}

	actualItem := actual.FindOne()

	if actualItem.ID != 65193 {
		t.Fatal("expected ID 65193, found: ", actualItem.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := disks.FilterByName("bootdisk")

	if len(actual.Data) != 2 {
		t.Fatal("expected 2 elements, found: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.Name != "bootdisk" {
			t.Fatal("expected 'bootdisk' name, found: ", item.Name)
		}
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := disks.FilterByStatus("ASSIGNED")

	if len(actual.Data) == 0 {
		t.Fatal("No elements were found")
	}

	for _, item := range actual.Data {
		if item.Status != "ASSIGNED" {
			t.Fatal("expected 'ASSIGNED' status, found: ", item.Status)
		}
	}
}

func TestFilterByTechStatus(t *testing.T) {
	actual := disks.FilterByTechStatus("ALLOCATED")

	if len(actual.Data) == 0 {
		t.Fatal("No elements were found")
	}

	for _, item := range actual.Data {
		if item.TechStatus != "ALLOCATED" {
			t.Fatal("expected 'ALLOCATED' techStatus, found: ", item.TechStatus)
		}
	}
}

func TestFilterByImageID(t *testing.T) {
	actual := disks.FilterByImageID(9885)

	if len(actual.Data) == 0 {
		t.Fatal("No elements were found")
	}

	if actual.Data[0].ImageID != 9885 {
		t.Fatal("expected 9885 ImageID, found: ", actual.Data[0].ImageID)
	}
}

func TestFilterFunc(t *testing.T) {
	actual := disks.FilterFunc(func(id ItemDisk) bool {
		return len(id.PresentTo) == 2
	})

	if len(actual.Data) == 0 {
		t.Fatal("No elements were found")
	}

	if len(actual.Data[0].PresentTo) != 2 {
		t.Fatal("expected 2 elements in PresentTo, found: ", len(actual.Data[0].PresentTo))
	}
}

func TestSortByCreatedTime(t *testing.T) {
	actual := disks.SortByCreatedTime(false)

	if actual.Data[0].ID != 65191 {
		t.Fatal("expected ID 65191, found: ", actual.Data[0].ID)
	}

	actual = disks.SortByCreatedTime(true)

	if actual.Data[0].ID != 65193 {
		t.Fatal("expected ID 65193, found: ", actual.Data[0].ID)
	}
}
