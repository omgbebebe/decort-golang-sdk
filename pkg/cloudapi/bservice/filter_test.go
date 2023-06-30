package bservice

import "testing"

var bservices = ListBasicServices{
	Data: []ItemBasicService{
		{
			AccountID:     1,
			AccountName:   "std_1",
			BaseDomain:    "",
			CreatedBy:     "sample_user_1@decs3o",
			CreatedTime:   1677743675,
			DeletedBy:     "",
			DeletedTime:   0,
			GID:           212,
			Groups:        []uint64{},
			GUID:          1,
			ID:            1,
			Name:          "bservice_1",
			ParentSrvID:   0,
			ParentSrvType: "",
			RGID:          7971,
			RGName:        "rg_1",
			SSHUser:       "",
			Status:        "CREATED",
			TechStatus:    "STOPPED",
			UpdatedBy:     "",
			UpdatedTime:   0,
			UserManaged:   true,
		},
		{
			AccountID:     2,
			AccountName:   "std_2",
			BaseDomain:    "",
			CreatedBy:     "sample_user_1@decs3o",
			CreatedTime:   1677743736,
			DeletedBy:     "",
			DeletedTime:   0,
			GID:           212,
			Groups:        []uint64{},
			GUID:          2,
			ID:            2,
			Name:          "bservice_2",
			ParentSrvID:   0,
			ParentSrvType: "",
			RGID:          7972,
			RGName:        "rg_2",
			SSHUser:       "",
			Status:        "CREATED",
			TechStatus:    "STOPPED",
			UpdatedBy:     "",
			UpdatedTime:   0,
			UserManaged:   true,
		},
		{
			AccountID:     3,
			AccountName:   "std_3",
			BaseDomain:    "",
			CreatedBy:     "sample_user_2@decs3o",
			CreatedTime:   1677743830,
			DeletedBy:     "",
			DeletedTime:   0,
			GID:           212,
			Groups:        []uint64{},
			GUID:          3,
			ID:            3,
			Name:          "bservice_3",
			ParentSrvID:   0,
			ParentSrvType: "",
			RGID:          7973,
			RGName:        "rg_3",
			SSHUser:       "",
			Status:        "ENABLED",
			TechStatus:    "STARTED",
			UpdatedBy:     "",
			UpdatedTime:   0,
			UserManaged:   true,
		},
	},
	EntryCount: 3,
}

func TestFilterByID(t *testing.T) {
	actual := bservices.FilterByID(1).FindOne()

	if actual.ID != 1 {
		t.Fatal("expected ID 1, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := bservices.FilterByName("bservice_3").FindOne()

	if actual.Name != "bservice_3" {
		t.Fatal("expected Name 'bservice_3', found: ", actual.Name)
	}
}

func TestFilterByRGID(t *testing.T) {
	actual := bservices.FilterByRGID(7971).FindOne()

	if actual.RGID != 7971 {
		t.Fatal("expected RGID 7971, found: ", actual.RGID)
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := bservices.FilterByStatus("CREATED")

	if len(actual.Data) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.Status != "CREATED" {
			t.Fatal("expected Status 'CREATED', found: ", item.Status)
		}
	}
}

func TestFilterByTechStatus(t *testing.T) {
	actual := bservices.FilterByTechStatus("STOPPED")

	if len(actual.Data) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.TechStatus != "STOPPED" {
			t.Fatal("expected TechStatus 'STOPPED', found: ", item.TechStatus)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := bservices.FilterFunc(func(ibs ItemBasicService) bool {
		return ibs.CreatedBy == "sample_user_2@decs3o"
	})

	if len(actual.Data) > 1 {
		t.Fatal("expected 1 found, actual: ", len(actual.Data))
	}

	if actual.FindOne().CreatedBy != "sample_user_2@decs3o" {
		t.Fatal("expected 'sample_user_2@decs3o', found: ", actual.FindOne().CreatedBy)
	}
}

func TestSortByCreatedTime(t *testing.T) {
	actual := bservices.SortByCreatedTime(true)

	if actual.Data[0].CreatedTime != 1677743830 || actual.Data[2].CreatedTime != 1677743675 {
		t.Fatal("expected descending order, found ascending")
	}
}
