package extnet

import "testing"

var extnets = ListExtNets{
	Data: []ItemExtNet{
		{
			ID:     3,
			IPCIDR: "176.118.164.0/24",
			Name:   "176.118.164.0/24",
			Status: "ENABLED",
		},
		{
			ID:     10,
			IPCIDR: "45.134.255.0/24",
			Name:   "45.134.255.0/24",
			Status: "ENABLED",
		},
		{
			ID:     13,
			IPCIDR: "88.218.249.0/24",
			Name:   "88.218.249.0/24",
			Status: "DISABLED",
		},
	},
}

func TestFilterByID(t *testing.T) {
	actual := extnets.FilterByID(10).FindOne()

	if actual.ID != 10 {
		t.Fatal("expected ID 10, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	name := "88.218.249.0/24"
	actual := extnets.FilterByName(name).FindOne()

	if actual.Name != name {
		t.Fatal("expected ", name, " found: ", actual.Name)
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := extnets.FilterByStatus("ENABLED")

	if len(actual.Data) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.Status != "ENABLED" {
			t.Fatal("expected Status 'ENABLED', found: ", item.Status)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := extnets.FilterFunc(func(ien ItemExtNet) bool {
		return ien.IPCIDR == ien.Name
	})

	if len(actual.Data) != 3 {
		t.Fatal("expected 3 elements, found: ", len(actual.Data))
	}
}
