package rg

import "testing"

var rgs = ListResourceGroups{
	{
		AccountID:   1,
		AccountName: "std",
		ACL: ListACL{
			{
				Explicit:    true,
				GUID:        "",
				Right:       "ARCXDU",
				Status:      "CONFIRMED",
				Type:        "U",
				UserGroupID: "sample_user_1@decs3o",
			},
		},
		CreatedBy:        "sample_user_1@decs3o",
		CreatedTime:      1676645305,
		DefNetID:         1,
		DefNetType:       "NONE",
		DeletedBy:        "",
		DeletedTime:      0,
		Description:      "",
		GID:              212,
		GUID:             7971,
		ID:               7971,
		LockStatus:       "UNLOCKED",
		Milestones:       363459,
		Name:             "rg_1",
		RegisterComputes: false,
		ResourceLimits: ResourceLimits{
			CUC:      -1,
			CUI:      -1,
			CUM:      -1,
			CUNP:     -1,
			GPUUnits: -1,
		},
		Secret:      "",
		Status:      "CREATED",
		UpdatedBy:   "",
		UpdatedTime: 0,
		VINS:        []uint64{},
		Computes:    []uint64{},
		ResTypes:    []string{},
		UniqPools:   []string{},
	},
	{
		AccountID:   2,
		AccountName: "std_2",
		ACL: ListACL{
			{
				Explicit:    true,
				GUID:        "",
				Right:       "ARCXDU",
				Status:      "CONFIRMED",
				Type:        "U",
				UserGroupID: "sample_user_1@decs3o",
			},
		},
		CreatedBy:        "sample_user_1@decs3o",
		CreatedTime:      1676645461,
		DefNetID:         2,
		DefNetType:       "NONE",
		DeletedBy:        "",
		DeletedTime:      0,
		Description:      "",
		GID:              212,
		GUID:             7972,
		ID:               7972,
		LockStatus:       "UNLOCKED",
		Milestones:       363468,
		Name:             "rg_2",
		RegisterComputes: false,
		ResourceLimits: ResourceLimits{
			CUC:      -1,
			CUI:      -1,
			CUM:      -1,
			CUNP:     -1,
			GPUUnits: -1,
		},
		Secret:      "",
		Status:      "CREATED",
		UpdatedBy:   "",
		UpdatedTime: 0,
		VINS:        []uint64{},
		Computes:    []uint64{},
		ResTypes:    []string{},
		UniqPools:   []string{},
	},
	{
		AccountID:   3,
		AccountName: "std_3",
		ACL: ListACL{
			{
				Explicit:    true,
				GUID:        "",
				Right:       "ARCXDU",
				Status:      "CONFIRMED",
				Type:        "U",
				UserGroupID: "sample_user_2@decs3o",
			},
		},
		CreatedBy:        "sample_user_2@decs3o",
		CreatedTime:      1676645548,
		DefNetID:         3,
		DefNetType:       "NONE",
		DeletedBy:        "",
		DeletedTime:      0,
		Description:      "",
		GID:              212,
		GUID:             7973,
		ID:               7973,
		LockStatus:       "kjLOCKED",
		Milestones:       363471,
		Name:             "rg_3",
		RegisterComputes: false,
		ResourceLimits: ResourceLimits{
			CUC:      -1,
			CUI:      -1,
			CUM:      -1,
			CUNP:     -1,
			GPUUnits: -1,
		},
		Secret:      "",
		Status:      "DISABLED",
		UpdatedBy:   "",
		UpdatedTime: 0,
		VINS:        []uint64{},
		Computes: []uint64{
			48500,
		},
		ResTypes:  []string{},
		UniqPools: []string{},
	},
}

func TestFilterByID(t *testing.T) {
	actual := rgs.FilterByID(7972).FindOne()

	if actual.ID != 7972 {
		t.Fatal("expected ID 2, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := rgs.FilterByName("rg_1").FindOne()

	if actual.Name != "rg_1" {
		t.Fatal("expected Name 'rg_1', found: ", actual.Name)
	}
}

func TestFilterByCreatedBy(t *testing.T) {
	actual := rgs.FilterByCreatedBy("sample_user_1@decs3o")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.CreatedBy != "sample_user_1@decs3o" {
			t.Fatal("expected CreatedBy 'sample_user_1@decs3o', found: ", item.CreatedBy)
		}
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := rgs.FilterByStatus("CREATED")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.Status != "CREATED" {
			t.Fatal("expected Status 'ENABLED', found: ", item.Status)
		}
	}
}

func TestFilterByLockStatus(t *testing.T) {
	actual := rgs.FilterByLockStatus("UNLOCKED")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.LockStatus != "UNLOCKED" {
			t.Fatal("expected LockStatus 'UNLOCKED', found: ", item.LockStatus)
		}
	}
}

func TestFilterByDefNetType(t *testing.T) {
	actual := rgs.FilterByDefNetType("NONE")

	if len(actual) != 3 {
		t.Fatal("expected 3 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.DefNetType != "NONE" {
			t.Fatal("expected DefNetType 'NONE', found: ", item.DefNetType)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := rgs.FilterFunc(func(ir ItemResourceGroup) bool {
		return len(ir.Computes) > 0
	})

	if len(actual) < 1 {
		t.Fatal("expected 1 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if len(item.Computes) < 1 {
			t.Fatal("expected VMs to contain at least 1 element, found empty")
		}
	}
}

func TestSortByCreatedTime(t *testing.T) {
	actual := rgs.SortByCreatedTime(true)

	if actual[0].CreatedTime != 1676645548 || actual[2].CreatedTime != 1676645305 {
		t.Fatal("expected descending order, found ascending")
	}
}
