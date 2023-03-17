package sizes

import "testing"

var sizeItems = ListSizes{
	{
		Description: "",
		Disks:       []uint64{},
		ID:          1,
		Memory:      512,
		Name:        "size_1",
		VCPUs:       2,
	},
	{
		Description: "",
		Disks:       []uint64{},
		ID:          2,
		Memory:      1024,
		Name:        "size_2",
		VCPUs:       4,
	},
	{
		Description: "",
		Disks:       []uint64{},
		ID:          2,
		Memory:      2048,
		Name:        "size_3",
		VCPUs:       6,
	},
}

func TestFilterByID(t *testing.T) {
	actual := sizeItems.FilterByID(1).FindOne()

	if actual.ID != 1 {
		t.Fatal("expected ID 1, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := sizeItems.FilterByName("size_2").FindOne()

	if actual.Name != "size_2" {
		t.Fatal("expected Name 'size_2', found: ", actual.Name)
	}
}

func TestFilterFunc(t *testing.T) {
	actual := sizeItems.FilterFunc(func(is ItemSize) bool {
		return is.Memory > 512
	})

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.Memory <= 512 {
			t.Fatal("expected Memory greater than 512, found: ", item.Memory)
		}
	}
}
