package sep

import "testing"

var seps = ListSEP{
	{
		CKey: "",
		Meta: []interface{}{
			"osismodel",
			"cloudbroker",
			"sep",
			1,
		},
		Config: map[string]interface{}{
			"API_IPs": []string{
				"10.212.3.61",
				"10.212.3.62",
				"10.212.3.63",
			},
		},
		ConsumedBy: []uint64{
			27,
		},
		Description: "",
		GID:         212,
		GUID:        1,
		ID:          1,
		Milestones:  278329,
		Name:        "sep_1",
		ObjStatus:   "CREATED",
		ProvidedBy: []uint64{
			24,
			35,
			29,
		},
		SharedWith: []uint64{},
		TechStatus: "ENABLED",
		Type:       "DES",
	},
	{
		CKey: "",
		Meta: []interface{}{
			"osismodel",
			"cloudbroker",
			"sep",
			1,
		},
		Config: map[string]interface{}{
			"API_IPs": []string{
				"10.212.3.64",
				"10.212.3.65",
				"10.212.3.66",
			},
		},
		ConsumedBy: []uint64{
			32,
			26,
		},
		Description: "",
		GID:         212,
		GUID:        2,
		ID:          2,
		Milestones:  278337,
		Name:        "sep_2",
		ObjStatus:   "CREATED",
		ProvidedBy: []uint64{
			36,
			42,
			35,
		},
		SharedWith: []uint64{},
		TechStatus: "ENABLED",
		Type:       "DES",
	},
	{
		CKey: "",
		Meta: []interface{}{
			"osismodel",
			"cloudbroker",
			"sep",
			1,
		},
		Config: map[string]interface{}{
			"API_IPs": []string{
				"10.212.3.67",
				"10.212.3.68",
				"10.212.3.69",
			},
		},
		ConsumedBy: []uint64{
			38,
			28,
		},
		Description: "",
		GID:         212,
		GUID:        3,
		ID:          3,
		Milestones:  278345,
		Name:        "sep_3",
		ObjStatus:   "DESTROYED",
		ProvidedBy: []uint64{
			49,
			48,
			41,
		},
		SharedWith: []uint64{},
		TechStatus: "DISABLED",
		Type:       "DES",
	},
}

func TestFilterByID(t *testing.T) {
	actual := seps.FilterByID(1).FindOne()

	if actual.ID != 1 {
		t.Fatal("expected ID 1, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := seps.FilterByName("sep_2").FindOne()

	if actual.Name != "sep_2" {
		t.Fatal("expected Name 'sep_2', found: ", actual.Name)
	}
}

func TestFilterByObjStatus(t *testing.T) {
	actual := seps.FilterByObjStatus("CREATED")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.ObjStatus != "CREATED" {
			t.Fatal("expected ObjStatus 'CREATED', found: ", item.ObjStatus)
		}
	}
}

func TestFilterByTechStatus(t *testing.T) {
	actual := seps.FilterByTechStatus("ENABLED")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.TechStatus != "ENABLED" {
			t.Fatal("expected TechStatus 'ENABLED', found: ", item.TechStatus)
		}
	}
}

func TestFilterByType(t *testing.T) {
	actual := seps.FilterByType("DES")

	if len(actual) != 3 {
		t.Fatal("expected 3 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.Type != "DES" {
			t.Fatal("expected Type 'DES', found: ", item.Type)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := seps.FilterFunc(func(rs RecordSEP) bool {
		return len(rs.ConsumedBy) > 1
	})

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if len(item.ConsumedBy) <= 1 {
			t.Fatal("expected ConsumedBy to contain more than 1 element, found: ", len(item.ConsumedBy))
		}
	}
}
