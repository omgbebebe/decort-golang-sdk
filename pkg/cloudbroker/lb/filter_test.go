package lb

import "testing"

var lbs = ListLB{
	Data: []ItemLBList{
		{
			HAMode:        true,
			ACL:           []interface{}{},
			Backends:      []ItemBackend{},
			CreatedBy:     "test_user_1",
			CreatedTime:   1636667448,
			DeletedBy:     "",
			DeletedTime:   0,
			Description:   "",
			DPAPIPassword: "0000",
			DPAPIUser:     "api_user",
			ExtNetID:      2522,
			Frontends:     []ItemFrontend{},
			GID:           212,
			GUID:          1,
			ID:            1,
			Milestones:    129000,
			Name:          "k8s-lb-test-1",
			PrimaryNode:   Node{},
			RGID:          25090,
			RGName:        "",
			SecondaryNode: Node{},
			Status:        "ENABLED",
			TechStatus:    "STARTED",
			UpdatedBy:     "",
			UpdatedTime:   0,
			VINSID:        101,
		},
		{
			HAMode:        false,
			ACL:           []interface{}{},
			Backends:      []ItemBackend{},
			CreatedBy:     "test_user_2",
			CreatedTime:   1636667506,
			DeletedBy:     "",
			DeletedTime:   0,
			Description:   "",
			DPAPIPassword: "0000",
			DPAPIUser:     "api_user_2",
			ExtNetID:      2524,
			Frontends:     []ItemFrontend{},
			GID:           212,
			GUID:          2,
			ID:            2,
			Milestones:    129013,
			Name:          "k8s-lb-test-2",
			PrimaryNode:   Node{},
			RGID:          25092,
			RGName:        "",
			SecondaryNode: Node{},
			Status:        "ENABLED",
			TechStatus:    "STOPPED",
			UpdatedBy:     "",
			UpdatedTime:   0,
			VINSID:        102,
		},
		{
			HAMode:        true,
			ACL:           []interface{}{},
			Backends:      []ItemBackend{},
			CreatedBy:     "te2t_user_3",
			CreatedTime:   1636667534,
			DeletedBy:     "",
			DeletedTime:   0,
			Description:   "",
			DPAPIPassword: "0000",
			DPAPIUser:     "api_user_3",
			ExtNetID:      2536,
			Frontends:     []ItemFrontend{},
			GID:           212,
			GUID:          3,
			ID:            3,
			Milestones:    129025,
			Name:          "k8s-lb-test-3",
			PrimaryNode:   Node{},
			RGID:          25106,
			RGName:        "",
			SecondaryNode: Node{},
			Status:        "DISABLED",
			TechStatus:    "STOPPED",
			UpdatedBy:     "",
			UpdatedTime:   0,
			VINSID:        118,
		},
	},
	EntryCount: 3,
}

func TestFilterByID(t *testing.T) {
	actual := lbs.FilterByID(2).FindOne()

	if actual.ID != 2 {
		t.Fatal("expected ID 2, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := lbs.FilterByName("k8s-lb-test-3").FindOne()

	if actual.Name != "k8s-lb-test-3" {
		t.Fatal("expected Name 'k8s-lb-test-3', found: ", actual.Name)
	}
}

func TestFilterByExtNetID(t *testing.T) {
	actual := lbs.FilterByExtNetID(2522).FindOne()

	if actual.ExtNetID != 2522 {
		t.Fatal("expected ExtNetID 2522, found: ", actual.ExtNetID)
	}
}

func TestFilterFunc(t *testing.T) {
	actual := lbs.FilterFunc(func(rl ItemLBList) bool {
		return rl.Status == "DISABLED"
	})

	for _, item := range actual.Data {
		if item.Status != "DISABLED" {
			t.Fatal("expected Status 'DISABLED', found: ", item.Status)
		}
	}
}

func TestSortByCreatedTime(t *testing.T) {
	actual := lbs.SortByCreatedTime(true)

	if actual.Data[0].CreatedTime != 1636667534 || actual.Data[2].CreatedTime != 1636667448 {
		t.Fatal("expected descending order, found ascending")
	}
}
