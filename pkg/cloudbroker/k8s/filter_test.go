package k8s

import "testing"

var k8sItems = ListK8S{
	ItemK8S{
		AccountID:      1,
		AccountName:    "test_1",
		ACL:            []interface{}{},
		BServiceID:     1,
		CIID:           1,
		Config:         nil,
		CreatedBy:      "test_user",
		CreatedTime:    132454563,
		DeletedBy:      "",
		DeletedTime:    0,
		Description:    "",
		ExtNetID:       1,
		GID:            0,
		GUID:           1,
		ID:             1,
		LBID:           1,
		Milestones:     999999,
		Name:           "k8s_1",
		RGID:           1,
		RGName:         "rg_1",
		ServiceAccount: ServiceAccount{},
		SSHKey:         "sample_key",
		Status:         "ENABLED",
		TechStatus:     "STARTED",
		UpdatedBy:      "",
		UpdatedTime:    0,
		VINSID:         0,
		WorkersGroup:   []RecordK8SGroup{},
	},
	ItemK8S{
		AccountID:      2,
		AccountName:    "test_2",
		ACL:            []interface{}{},
		BServiceID:     2,
		CIID:           2,
		Config:         nil,
		CreatedBy:      "test_user",
		CreatedTime:    132454638,
		DeletedBy:      "",
		DeletedTime:    0,
		Description:    "",
		ExtNetID:       2,
		GID:            0,
		GUID:           2,
		ID:             2,
		LBID:           2,
		Milestones:     999999,
		Name:           "k8s_2",
		RGID:           2,
		RGName:         "rg_2",
		ServiceAccount: ServiceAccount{},
		SSHKey:         "sample_key",
		Status:         "ENABLED",
		TechStatus:     "STARTED",
		UpdatedBy:      "",
		UpdatedTime:    0,
		VINSID:         0,
		WorkersGroup:   []RecordK8SGroup{},
	},
	ItemK8S{
		AccountID:      3,
		AccountName:    "test_3",
		ACL:            []interface{}{},
		BServiceID:     3,
		CIID:           3,
		Config:         nil,
		CreatedBy:      "test_user",
		CreatedTime:    132454682,
		DeletedBy:      "",
		DeletedTime:    0,
		Description:    "",
		ExtNetID:       3,
		GID:            0,
		GUID:           3,
		ID:             3,
		LBID:           3,
		Milestones:     999999,
		Name:           "k8s_3",
		RGID:           3,
		RGName:         "rg_3",
		ServiceAccount: ServiceAccount{},
		SSHKey:         "sample_key",
		Status:         "DISABLED",
		TechStatus:     "STOPPED",
		UpdatedBy:      "",
		UpdatedTime:    0,
		VINSID:         0,
		WorkersGroup:   []RecordK8SGroup{},
	},
}

func TestFilterByID(t *testing.T) {
	actual := k8sItems.FilterByID(1).FindOne()

	if actual.ID != 1 {
		t.Fatal("expected 1 ID, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := k8sItems.FilterByName("k8s_3").FindOne()

	if actual.Name != "k8s_3" {
		t.Fatal("expected Name 'k8s_3', found: ", actual.Name)
	}
}

func TestFilterByAccountID(t *testing.T) {
	actual := k8sItems.FilterByAccountID(2).FindOne()

	if actual.AccountID != 2 {
		t.Fatal("expected AccountID 2, found: ", actual.AccountID)
	}
}

func TestFilterByRGID(t *testing.T) {
	actual := k8sItems.FilterByRGID(3).FindOne()

	if actual.RGID != 3 {
		t.Fatal("expected RGID 3, found: ", actual.RGID)
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := k8sItems.FilterByStatus("ENABLED")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.Status != "ENABLED" {
			t.Fatal("expected Status 'ENABLED', found: ", item.Status)
		}
	}
}

func TestFilterByTechStatus(t *testing.T) {
	actual := k8sItems.FilterByTechStatus("STARTED")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.TechStatus != "STARTED" {
			t.Fatal("expected TechStatus 'STARTED', found: ", item.TechStatus)
		}
	}
}

func TestFilterByCreatedBy(t *testing.T) {
	actual := k8sItems.FilterByCreatedBy("test_user")

	if len(actual) != 3 {
		t.Fatal("expected 3 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.CreatedBy != "test_user" {
			t.Fatal("expected CreatedBy 'test_user', found: ", item.CreatedBy)
		}
	}
}

func TestFilterByDeletedBy(t *testing.T) {
	actual := k8sItems.FilterByDeletedBy("test_user")

	if len(actual) != 0 {
		t.Fatal("expected 0 found, actual: ", len(actual))
	}
}

func TestFilterFunc(t *testing.T) {
	actual := k8sItems.FilterFunc(func(iks ItemK8S) bool {
		return iks.AccountName == "test_2"
	}).
		FindOne()

	if actual.AccountName != "test_2" {
		t.Fatal("expected AccountName 'test_2', found: ", actual.AccountName)
	}
}

func TestSortByCreatedTime(t *testing.T) {
	actual := k8sItems.SortByCreatedTime(false)

	if actual[0].CreatedTime != 132454563 || actual[2].CreatedTime != 132454682 {
		t.Fatal("expected ascending sort, seems to be inversed")
	}
}
