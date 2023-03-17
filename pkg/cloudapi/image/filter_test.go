package image

import "testing"

var images = ListImages{
	ItemImage{
		AccountID:    0,
		Architecture: "X86_64",
		BootType:     "bios",
		Bootable:     true,
		Description:  "",
		Drivers: []string{
			"KVM_X86",
		},
		HotResize: true,
		ID:        9882,
		LinkTo:    0,
		Name:      "u16",
		Pool:      "vmstor",
		Size:      5,
		Status:    "CREATED",
		Type:      "linux",
		Username:  "",
		Virtual:   false,
	},
	ItemImage{
		AccountID:    0,
		Architecture: "X86_64",
		BootType:     "bois",
		Bootable:     true,
		Description:  "",
		Drivers: []string{
			"KVM_X86",
		},
		HotResize: false,
		ID:        9884,
		LinkTo:    0,
		Name:      "alpine-virt-3.17",
		Pool:      "vmstor",
		Size:      1,
		Status:    "CREATED",
		Type:      "linux",
		Username:  "",
		Virtual:   true,
	},
	ItemImage{
		AccountID:    1,
		Architecture: "X86_64",
		BootType:     "bios",
		Bootable:     true,
		Description:  "",
		Drivers: []string{
			"KVM_X86",
		},
		HotResize: true,
		ID:        9885,
		LinkTo:    0,
		Name:      "test",
		Pool:      "vmstor",
		Size:      4,
		Status:    "DESTROYED",
		Type:      "linux",
		Username:  "",
		Virtual:   false,
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

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.Status != "CREATED" {
			t.Fatal("expected Status 'CREATED', found: ", item.Status)
		}
	}
}

func TestFilterByBootType(t *testing.T) {
	actual := images.FilterByBootType("bios")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.BootType != "bios" {
			t.Fatal("expected BootType 'bios', found: ", item.BootType)
		}
	}
}

func TestFilterFunc(t *testing.T) {
    actual := images.FilterFunc(func(ii ItemImage) bool {
        return ii.Virtual == true
    })

	if len(actual) != 1 {
		t.Fatal("expected 1 found, actual: ", len(actual))
	}

	if actual[0].Virtual != true {
		t.Fatal("expected Virtual true, found false")
	}
}
