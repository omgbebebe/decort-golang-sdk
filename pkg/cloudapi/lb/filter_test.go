package lb

import "testing"

var lbs = ListLB{
	ItemLoadBalancer{
		DPAPIPassword: "0000",
		RecordLB: RecordLB{
			HAMode:      true,
			ACL:         []interface{}{},
			Backends:    []ItemBackend{},
			CreatedBy:   "test_user_1",
			CreatedTime: 1636667448,
			DeletedBy:   "",
			DeletedTime: 0,
			Description: "",
			DPAPIUser:   "api_user",
			ExtNetID:    2522,
			Frontends:   []ItemFrontend{},
			GID:         212,
			GUID:        1,
			ID:          1,
			ImageID:     2121,
			Milestones:  129000,
			Name:        "k8s-lb-test-1",
			RGID:        25090,
			RGName:      "",
			Status:      "ENABLED",
			TechStatus:  "STARTED",
			UpdatedBy:   "",
			UpdatedTime: 0,
			VINSID:      101,
		},
	},
	ItemLoadBalancer{
		DPAPIPassword: "0000",
		RecordLB: RecordLB{
			HAMode:      false,
			ACL:         []interface{}{},
			Backends:    []ItemBackend{},
			CreatedBy:   "test_user_2",
			CreatedTime: 1636667506,
			DeletedBy:   "",
			DeletedTime: 0,
			Description: "",
			DPAPIUser:   "api_user_2",
			ExtNetID:    2524,
			Frontends:   []ItemFrontend{},
			GID:         212,
			GUID:        2,
			ID:          2,
			ImageID:     2129,
			Milestones:  129013,
			Name:        "k8s-lb-test-2",
			RGID:        25092,
			RGName:      "",
			Status:      "ENABLED",
			TechStatus:  "STOPPED",
			UpdatedBy:   "",
			UpdatedTime: 0,
			VINSID:      102,
		},
	},
	ItemLoadBalancer{
		DPAPIPassword: "0000",
		RecordLB: RecordLB{
			HAMode:      true,
			ACL:         []interface{}{},
			Backends:    []ItemBackend{},
			CreatedBy:   "te2t_user_3",
			CreatedTime: 1636667534,
			DeletedBy:   "",
			DeletedTime: 0,
			Description: "",
			DPAPIUser:   "api_user_3",
			ExtNetID:    2536,
			Frontends:   []ItemFrontend{},
			GID:         212,
			GUID:        3,
			ID:          3,
			ImageID:     2139,
			Milestones:  129025,
			Name:        "k8s-lb-test-3",
			RGID:        25106,
			RGName:      "",
			Status:      "DISABLED",
			TechStatus:  "STOPPED",
			UpdatedBy:   "",
			UpdatedTime: 0,
			VINSID:      118,
		},
	},
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

func TestFilterByImageID(t *testing.T) {
	actual := lbs.FilterByImageID(2139).FindOne()

	if actual.ImageID != 2139 {
		t.Fatal("expected ImageID 2139, found: ", actual.ImageID)
	}
}

func TestFilterFunc(t *testing.T) {
	actual := lbs.FilterFunc(func(rl ItemLoadBalancer) bool {
		return rl.Status == "DISABLED"
	})

	for _, item := range actual {
		if item.Status != "DISABLED" {
			t.Fatal("expected Status 'DISABLED', found: ", item.Status)
		}
	}
}

func TestSortByCreatedTime(t *testing.T) {
	actual := lbs.SortByCreatedTime(true)

	if actual[0].CreatedTime != 1636667534 || actual[2].CreatedTime != 1636667448 {
		t.Fatal("expected descending order, found ascending")
	}
}
