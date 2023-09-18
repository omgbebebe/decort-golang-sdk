package k8ci

import "testing"

var k8ciItems = ListK8CI{
	Data: []ItemK8CI{
		{
			CreatedTime: 123902139,
			RecordK8CIList: RecordK8CIList{
				Description:    "",
				GID:            0,
				GUID:           1,
				ID:             1,
				LBImageID:      5,
				MasterDriver:   "KVM_X86",
				MasterImageID:  120,
				MaxMasterCount: 2,
				MaxWorkerCount: 3,
				Name:           "purple_snake",
				SharedWith:     []uint64{},
				Status:         "ENABLED",
				Version:        "1",
				WorkerDriver:   "KVM_X86",
				WorkerImageID:  120,
			},
		},
		{
			CreatedTime: 123902232,
			RecordK8CIList: RecordK8CIList{
				Description:    "",
				GID:            0,
				GUID:           2,
				ID:             2,
				LBImageID:      10,
				MasterDriver:   "KVM_X86",
				MasterImageID:  121,
				MaxMasterCount: 3,
				MaxWorkerCount: 5,
				Name:           "green_giant",
				SharedWith:     []uint64{},
				Status:         "DISABLED",
				Version:        "2",
				WorkerDriver:   "KVM_X86",
				WorkerImageID:  121,
			},
		},
		{
			CreatedTime: 123902335,
			RecordK8CIList: RecordK8CIList{
				Description:    "",
				GID:            0,
				GUID:           3,
				ID:             3,
				LBImageID:      12,
				MasterDriver:   "KVM_X86",
				MasterImageID:  98,
				MaxMasterCount: 5,
				MaxWorkerCount: 9,
				Name:           "magenta_cloud",
				SharedWith:     []uint64{},
				Status:         "ENABLED",
				Version:        "3",
				WorkerDriver:   "KVM_X86",
				WorkerImageID:  98,
			},
		},
	},
	EntryCount: 3,
}

func TestFilterByID(t *testing.T) {
	actual := k8ciItems.FilterByID(2).FindOne()

	if actual.ID != 2 {
		t.Fatal("expected ID 2, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := k8ciItems.FilterByName("magenta_cloud").FindOne()

	if actual.Name != "magenta_cloud" {
		t.Fatal("expected Name 'magenta_cloud', found: ", actual.Name)
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := k8ciItems.FilterByStatus("ENABLED")

	if len(actual.Data) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.Status != "ENABLED" {
			t.Fatal("expected Status 'ENABLED', found: ", item.Status)
		}
	}
}

func TestFilterByWorkerImageID(t *testing.T) {
	actual := k8ciItems.FilterByWorkerImageID(98).FindOne()

	if actual.WorkerImageID != 98 {
		t.Fatal("expected WorkerImageID 98, found: ", actual.WorkerImageID)
	}
}

func TestFilterByLBImageID(t *testing.T) {
	actual := k8ciItems.FilterByLBImageID(10).FindOne()

	if actual.LBImageID != 10 {
		t.Fatal("expected LBImageID 10, found: ", actual.LBImageID)
	}
}

func TestFilterByMasterImageID(t *testing.T) {
	actual := k8ciItems.FilterByMasterImageID(120).FindOne()

	if actual.MasterImageID != 120 {
		t.Fatal("expected MasterImageID 120, found: ", actual.MasterImageID)
	}
}

func TestFilterFunc(t *testing.T) {
	actual := k8ciItems.FilterFunc(func(ikc ItemK8CI) bool {
		return ikc.CreatedTime > 123902139
	})

	if len(actual.Data) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.CreatedTime < 123902139 {
			t.Fatal("expected CreatedTime greater than 123902139, found: ", item.CreatedTime)
		}
	}
}

func TestSortingByCreatedTime(t *testing.T) {
	actual := k8ciItems.SortByCreatedTime(true)

	if actual.Data[0].CreatedTime != 123902335 && actual.Data[2].CreatedTime != 123902139 {
		t.Fatal("expected inverse sort, found normal")
	}
}
