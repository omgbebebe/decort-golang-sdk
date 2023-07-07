package image

import "testing"

var images = ListImages{
	Data: []RecordImage{
		{
			UNCPath: "",
			CKey:    "",
			Meta: []interface{}{
				"osismodel",
				"cloudbroker",
				"image",
				1,
			},
			AccountID:    0,
			ACL:          []ACL{},
			Architecture: "X86_64",
			BootType:     "bios",
			Bootable:     true,
			ComputeCIID:  0,
			DeletedTime:  0,
			Description:  "",
			Drivers: []string{
				"KVM_X86",
			},
			Enabled:       true,
			GID:           212,
			GUID:          9882,
			History:       []History{},
			HotResize:     true,
			ID:            9882,
			LastModified:  0,
			LinkTo:        0,
			Milestones:    363491,
			Name:          "u16",
			Password:      "",
			Pool:          "vmstor",
			PresentTo:     []uint64{},
			ProviderName:  "",
			PurgeAttempts: 0,
			ReferenceID:   "sample_reference_id_u16",
			ResID:         "b321318-3214as-324-213-fdas",
			ResName:       "templates/image_9882",
			RescueCD:      false,
			SEPID:         2504,
			SharedWith:    []uint64{},
			Size:          5,
			Status:        "CREATED",
			TechStatus:    "ALLOCATED",
			Type:          "linux",
			URL:           "http://sample_url:8000/u16",
			Username:      "",
			Version:       "",
			Virtual:       false,
		},
		{
			UNCPath: "",
			CKey:    "",
			Meta: []interface{}{
				"osismodel",
				"cloudbroker",
				"image",
				1,
			},
			AccountID:    0,
			ACL:          []ACL{},
			Architecture: "X86_64",
			BootType:     "bois",
			Bootable:     true,
			ComputeCIID:  0,
			DeletedTime:  0,
			Description:  "",
			Drivers: []string{
				"KVM_X86",
			},
			Enabled:       false,
			GID:           212,
			GUID:          9884,
			History:       []History{},
			HotResize:     false,
			ID:            9884,
			LastModified:  0,
			LinkTo:        0,
			Milestones:    363499,
			Name:          "alpine-virt-3.17",
			Password:      "",
			Pool:          "vmstor",
			PresentTo:     []uint64{},
			ProviderName:  "",
			PurgeAttempts: 0,
			ReferenceID:   "sample_reference_id_alpine",
			ResID:         "31d1d410-74f1-4e09-866b-046a5a8433c3",
			ResName:       "templates/image_9884",
			RescueCD:      false,
			SEPID:         2504,
			SharedWith:    []uint64{},
			Size:          1,
			Status:        "CREATED",
			TechStatus:    "ALLOCATED",
			Type:          "linux",
			URL:           "http://sample_url:8000/alpine-virt-3",
			Username:      "",
			Version:       "",
			Virtual:       true,
		},
		{
			UNCPath: "",
			CKey:    "",
			Meta: []interface{}{
				"osismodel",
				"cloudbroker",
				"image",
				1,
			},
			AccountID:    1,
			ACL:          []ACL{},
			Architecture: "X86_64",
			BootType:     "bios",
			Bootable:     true,
			ComputeCIID:  0,
			DeletedTime:  0,
			Description:  "",
			Drivers: []string{
				"KVM_X86",
			},
			Enabled:       true,
			GID:           212,
			GUID:          9885,
			History:       []History{},
			HotResize:     true,
			ID:            9885,
			LastModified:  0,
			LinkTo:        0,
			Milestones:    363513,
			Name:          "test",
			Password:      "",
			Pool:          "vmstor",
			PresentTo:     []uint64{},
			ProviderName:  "",
			PurgeAttempts: 0,
			ReferenceID:   "sample_reference_id_test",
			ResID:         "1f53b815-1ac9-4a4b-af98-a0a3b69a34bb",
			ResName:       "templates/image_9885",
			RescueCD:      false,
			SEPID:         2505,
			SharedWith:    []uint64{},
			Size:          4,
			Status:        "DESTROYED",
			TechStatus:    "ALLOCATED",
			Type:          "linux",
			URL:           "http://sample_url:8000/test",
			Username:      "",
			Version:       "",
			Virtual:       false,
		},
	},
}

func TestFilterByID(t *testing.T) {
	actual := images.FilterByID(9885).FindOne()

	if actual.ID != 9885 {
		t.Fatal("expected ID 9885, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := images.FilterByName("u16").FindOne()

	if actual.Name != "u16" {
		t.Fatal("expected Name 'u16', found: ", actual.Name)
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := images.FilterByStatus("CREATED")

	if len(actual.Data) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.Status != "CREATED" {
			t.Fatal("expected Status 'CREATED', found: ", item.Status)
		}
	}
}

func TestFilterByBootType(t *testing.T) {
	actual := images.FilterByBootType("bios")

	if len(actual.Data) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.BootType != "bios" {
			t.Fatal("expected BootType 'bios', found: ", item.BootType)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := images.FilterFunc(func(ri RecordImage) bool {
		return ri.Virtual == true
	})

	if len(actual.Data) != 1 {
		t.Fatal("expected 1 found, actual: ", len(actual.Data))
	}

	if actual.Data[0].Virtual != true {
		t.Fatal("expected Virtual true, found false")
	}
}
