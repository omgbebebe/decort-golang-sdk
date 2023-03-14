package extnet

import "testing"

var extnets = ListExtNet{
	ItemExtNet{
		CKey:               "",
		Meta:               []interface{}{},
		CheckIPs:           []string{},
		Default:            false,
		DefaultQOS:         QOS{},
		Description:        "",
		FreeIPs:            0,
		GID:                212,
		GUID:               3,
		ID:                 3,
		IPCIDR:             "176.118.164.0/24",
		Milestones:         1355466,
		Name:               "176.118.164.0/24",
		NetworkID:          0,
		OVSBridge:          "",
		PreReservationsNum: 0,
		PriVNFDevID:        0,
		SharedWith:         []interface{}{},
		Status:             "ENABLED",
		VLANID:             0,
		VNFs:               VNFs{},
	},
	ItemExtNet{
		CKey:               "",
		Meta:               []interface{}{},
		CheckIPs:           []string{},
		Default:            false,
		DefaultQOS:         QOS{},
		Description:        "",
		FreeIPs:            0,
		GID:                212,
		GUID:               10,
		ID:                 10,
		IPCIDR:             "45.134.255.0/24",
		Milestones:         2135543,
		Name:               "45.134.255.0/24",
		NetworkID:          0,
		OVSBridge:          "",
		PreReservationsNum: 0,
		PriVNFDevID:        0,
		SharedWith:         []interface{}{},
		Status:             "ENABLED",
		VLANID:             0,
		VNFs:               VNFs{},
	},
	ItemExtNet{
		CKey:               "",
		Meta:               []interface{}{},
		CheckIPs:           []string{},
		Default:            false,
		DefaultQOS:         QOS{},
		Description:        "",
		FreeIPs:            0,
		GID:                212,
		GUID:               13,
		ID:                 13,
		IPCIDR:             "88.218.249.0/24",
		Milestones:         1232134,
		Name:               "88.218.249.0/24",
		NetworkID:          0,
		OVSBridge:          "",
		PreReservationsNum: 0,
		PriVNFDevID:        0,
		SharedWith:         []interface{}{},
		Status:             "DISABLED",
		VLANID:             0,
		VNFs:               VNFs{},
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

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.Status != "ENABLED" {
			t.Fatal("expected Status 'ENABLED', found: ", item.Status)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := extnets.FilterFunc(func(ien ItemExtNet) bool {
		return ien.IPCIDR == ien.Name
	})

	if len(actual) != 3 {
		t.Fatal("expected 3 elements, found: ", len(actual))
	}
}
