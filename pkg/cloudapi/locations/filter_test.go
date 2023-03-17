package locations

import "testing"

var locationItems = ListLocations{
	{
		GID:          212,
		ID:           1,
		GUID:         1,
		LocationCode: "alfa",
		Name:         "alfa",
		Flag:         "",
		Meta: []interface{}{
			"cloudbroker",
			"location",
			1,
		},
		CKey: "",
	},
	{
		GID:          222,
		ID:           2,
		GUID:         2,
		LocationCode: "beta",
		Name:         "beta",
		Flag:         "",
		Meta: []interface{}{
			"cloudbroker",
			"location",
			1,
		},
		CKey: "",
	},
	{
		GID:          232,
		ID:           3,
		GUID:         3,
		LocationCode: "gamma",
		Name:         "gamma",
		Flag:         "",
		Meta: []interface{}{
			"cloudbroker",
			"location",
			1,
		},
		CKey: "",
	},
}

func TestFilterByID(t *testing.T) {
	actual := locationItems.FilterByID(1).FindOne()

	if actual.ID != 1 {
		t.Fatal("expected ID 1, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := locationItems.FilterByName("gamma").FindOne()

	if actual.Name != "gamma" {
		t.Fatal("expected Name 'gamma', found: ", actual.Name)
	}
}

func TestFilterFunc(t *testing.T) {
	actual := locationItems.FilterFunc(func(il ItemLocation) bool {
		return il.GID == 212
	}).
		FindOne()

	if actual.GID != 212 {
		t.Fatal("expected GID 212, found: ", actual.GID)
	}
}
