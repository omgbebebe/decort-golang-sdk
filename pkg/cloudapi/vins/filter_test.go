package vins

import "testing"

var vinsItems = ListVINS{
	{
		AccountID:   1,
		AccountName: "std",
		CreatedBy:   "sample_user_1@decs3o",
		CreatedTime: 1676898844,
		DeletedBy:   "",
		DeletedTime: 0,
		ExternalIP:  "",
		ID:          1,
		Name:        "vins01",
		Network:     "192.168.1.0/24",
		RGID:        7971,
		RGName:      "rg_01",
		Status:      "ENABLED",
		UpdatedBy:   "",
		UpdatedTime: 0,
		VXLANID:     3544,
	},
	{
		AccountID:   2,
		AccountName: "std2",
		CreatedBy:   "sample_user_1@decs3o",
		CreatedTime: 1676898948,
		DeletedBy:   "",
		DeletedTime: 0,
		ExternalIP:  "",
		ID:          2,
		Name:        "vins02",
		Network:     "192.168.2.0/24",
		RGID:        7972,
		RGName:      "rg_02",
		Status:      "ENABLED",
		UpdatedBy:   "",
		UpdatedTime: 0,
		VXLANID:     3545,
	},
	{
		AccountID:   3,
		AccountName: "std3",
		CreatedBy:   "sample_user_2@decs3o",
		CreatedTime: 1676899026,
		DeletedBy:   "",
		DeletedTime: 0,
		ExternalIP:  "",
		ID:          3,
		Name:        "vins03",
		Network:     "192.168.3.0/24",
		RGID:        7973,
		RGName:      "rg_03",
		Status:      "DISABLED",
		UpdatedBy:   "",
		UpdatedTime: 0,
		VXLANID:     3546,
	},
}

func TestFilterByID(t *testing.T) {
	actual := vinsItems.FilterByID(2).FindOne()

	if actual.ID != 2 {
		t.Fatal("expected ID 2, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := vinsItems.FilterByName("vins01").FindOne()

	if actual.Name != "vins01" {
		t.Fatal("expected Name 'vins01', found: ", actual.Name)
	}
}

func TestFilterByAccountID(t *testing.T) {
	actual := vinsItems.FilterByAccountID(3).FindOne()

	if actual.AccountID != 3 {
		t.Fatal("expected AccountID 3, found: ", actual.AccountID)
	}
}

func TestFilterByCreatedBy(t *testing.T) {
	actual := vinsItems.FilterByCreatedBy("sample_user_1@decs3o")

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
	actual := vinsItems.FilterFunc(func(iv ItemVINS) bool {
		return iv.RGID == 7971
	}).
		FindOne()

	if actual.RGID != 7971 {
		t.Fatal("expected RGID 7971, found: ", actual.RGID)
	}
}

func TestSortByCreatedTime(t *testing.T) {
	actual := vinsItems.SortByCreatedTime(false)

	if actual[0].CreatedTime != 1676898844 || actual[2].CreatedTime != 1676899026 {
		t.Fatal("expected ascending order, found descending")
	}
}
