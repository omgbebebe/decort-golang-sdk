package flipgroup

import "testing"

var flipgroups = ListFLIPGroups{
	{
		AccountID:   1,
		AccountName: "std_1",
		ClientIDs: []uint64{
			1,
		},
		ClientNames: []string{
			"compute_1",
		},
		ClientType:  "compute",
		ConnID:      1,
		ConnType:    "",
		CreatedBy:   "sample_user_1@decs3o",
		CreatedTime: 1234456789,
		DefaultGW:   "",
		DeletedBy:   "",
		DeletedTime: 0,
		Description: "",
		GID:         212,
		GUID:        1,
		ID:          1,
		IP:          "",
		Milestones:  999001,
		Name:        "flipgroup_1",
		NetID:       111,
		NetType:     "EXTNET",
		Network:     "",
		RGID:        7971,
		RGName:      "rg_1",
		Status:      "CREATED",
		UpdatedBy:   "",
		UpdatedTime: 0,
	},
	{
		AccountID:   2,
		AccountName: "std_2",
		ClientIDs: []uint64{
			2,
		},
		ClientNames: []string{
			"compute_2",
		},
		ClientType:  "compute",
		ConnID:      2,
		ConnType:    "",
		CreatedBy:   "sample_user_1@decs3o",
		CreatedTime: 1234456830,
		DefaultGW:   "",
		DeletedBy:   "sample_user_1@decs3o",
		DeletedTime: 1234456879,
		Description: "",
		GID:         212,
		GUID:        2,
		ID:          2,
		IP:          "",
		Milestones:  999002,
		Name:        "flipgroup_2",
		NetID:       222,
		NetType:     "EXTNET",
		Network:     "",
		RGID:        7972,
		RGName:      "rg_2",
		Status:      "DESTROYED",
		UpdatedBy:   "",
		UpdatedTime: 0,
	},
	{
		AccountID:   3,
		AccountName: "std_3",
		ClientIDs: []uint64{
			3,
		},
		ClientNames: []string{
			"compute_3",
		},
		ClientType:  "compute",
		ConnID:      3,
		ConnType:    "",
		CreatedBy:   "sample_user_2@decs3o",
		CreatedTime: 1234456866,
		DefaultGW:   "",
		DeletedBy:   "",
		DeletedTime: 0,
		Description: "",
		GID:         212,
		GUID:        3,
		ID:          3,
		IP:          "",
		Milestones:  999003,
		Name:        "flipgroup_3",
		NetID:       223,
		NetType:     "EXTNET",
		Network:     "",
		RGID:        7973,
		RGName:      "rg_3",
		Status:      "CREATED",
		UpdatedBy:   "",
		UpdatedTime: 0,
	},
}

func TestFilterByID(t *testing.T) {
	actual := flipgroups.FilterByID(3).FindOne()

	if actual.ID != 3 {
		t.Fatal("expected ID 3, found: ", actual.ID)
	}
}

func TestFilterByAccountID(t *testing.T) {
	actual := flipgroups.FilterByAccountID(1).FindOne()

	if actual.AccountID != 1 {
		t.Fatal("expected AccountID 1, found: ", actual.AccountID)
	}
}

func TestFilterByRGID(t *testing.T) {
	actual := flipgroups.FilterByRGID(7972).FindOne()

	if actual.RGID != 7972 {
		t.Fatal("expected RGID 7972, found: ", actual.RGID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := flipgroups.FilterByName("flipgroup_1").FindOne()

	if actual.Name != "flipgroup_1" {
		t.Fatal("expected Name 'flipgroup_1', found: ", actual.Name)
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := flipgroups.FilterByStatus("CREATED")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.Status != "CREATED" {
			t.Fatal("expected Status 'CREATED', found: ", item.Status)
		}
	}
}

func TestFilterByCreatedBy(t *testing.T) {
	actual := flipgroups.FilterByCreatedBy("sample_user_1@decs3o")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.CreatedBy != "sample_user_1@decs3o" {
			t.Fatal("expected CreatedBy 'sample_user_1@decs3o', found: ", item.CreatedBy)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := flipgroups.FilterFunc(func(ifg ItemFLIPGroup) bool {
		return ifg.NetType == "EXTNET"
	})

	if len(actual) != 3 {
		t.Fatal("expected 3 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.NetType != "EXTNET" {
			t.Fatal("expected NetType 'EXTNET', found: ", item.NetType)
		}
	}
}

func TestSortByCreatedTime(t *testing.T) {
	actual := flipgroups.SortByCreatedTime(false)

	if actual[0].CreatedTime != 1234456789 || actual[2].CreatedTime != 1234456866 {
		t.Fatal("expected ascending order, found descending")
	}
}
